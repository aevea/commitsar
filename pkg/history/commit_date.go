package history

import (
	"time"

	"gopkg.in/src-d/go-git.v4/plumbing"
)

// commitDate gets the commit at hash and returns the time of the commit
func (g *Git) commitDate(commit plumbing.Hash) (time.Time, error) {
	commitObject, err := g.repo.CommitObject(commit)

	if err != nil {
		return time.Now(), err
	}

	when := commitObject.Author.When

	return when, nil
}
