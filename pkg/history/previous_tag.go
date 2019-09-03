package history

import (
	"errors"
	"log"
	"sort"
	"time"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

var (
	// ErrPrevTagNotAvailable is returned when no previous tag is found.
	ErrPrevTagNotAvailable = errors.New("previous tag is not available")
)

type tag struct {
	hash plumbing.Hash
	time time.Time
}

// PreviousTag sorts tags based on when their commit happened and returns the one previous
// to the current.
func (g *Git) PreviousTag(currentHash plumbing.Hash) (plumbing.Hash, error) {
	tagrefs, err := g.repo.Tags()

	if err != nil {
		return currentHash, err
	}

	defer tagrefs.Close()

	var tagHashes []tag

	err = tagrefs.ForEach(func(t *plumbing.Reference) error {
		commitDate, err := g.commitDate(t.Hash())

		if err != nil {
			return err
		}

		tagHashes = append(tagHashes, tag{time: commitDate, hash: t.Hash()})
		return nil
	})

	if err != nil {
		if g.Debug {
			log.Printf("[ERR] getting previous tag failed: %v", err)
		}
		return currentHash, err
	}

	// Tags are alphabetically ordered. We need to sort them by date.
	sortedTags := sortTags(g.repo, tagHashes)

	if g.Debug {
		log.Println("Sorted tag output: ")
		for _, taginstance := range sortedTags {
			log.Printf("hash: %v time: %v", taginstance.hash, taginstance.time.UTC())
		}
	}

	// If there are fewer than two tags assume that the currentCommit is the newest tag
	if len(sortedTags) < 2 {
		if g.Debug {
			log.Println("[ERR] previous tag not available")
		}
		return currentHash, ErrPrevTagNotAvailable
	}

	if sortedTags[0].hash != currentHash {
		if g.Debug {
			log.Println("[ERR] current commit does not have a tag attached, building from this commit")
		}
		return sortedTags[0].hash, nil
	}

	if g.Debug {
		log.Printf("success: previous tag found at %v", sortedTags[1].hash)
	}

	return sortedTags[1].hash, nil
}

// sortTags sorts the tags according to when their parent commit happened.
func sortTags(repo *git.Repository, tags []tag) []tag {
	sort.Slice(tags, func(i, j int) bool {
		return tags[i].time.After(tags[j].time)
	})

	return tags
}
