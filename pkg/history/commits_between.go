package history

import (
	"errors"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

var (
	errReachedToCommit = errors.New("reached to commit")
)

// CommitsBetween returns a slice of commit hashes between two commits
func (g *Git) CommitsBetween(from plumbing.Hash, to plumbing.Hash) ([]plumbing.Hash, error) {
	cIter, _ := g.repo.Log(&git.LogOptions{From: from})

	var commits []plumbing.Hash

	err := cIter.ForEach(func(c *object.Commit) error {
		// If no previous tag is found then from and to are equal
		if from == to {
			return nil
		}
		if c.Hash == to {
			return errReachedToCommit
		}
		commits = append(commits, c.Hash)
		return nil
	})

	if err == errReachedToCommit {
		return commits, nil
	}
	if err != nil {
		return commits, err
	}
	return commits, nil
}
