package history

import (
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

// Commit find a commit based on commit hash and returns the Commit object
func (g *Git) Commit(hash plumbing.Hash) (*object.Commit, error) {
	commitObject, err := g.repo.CommitObject(hash)

	if err != nil {
		return nil, err
	}

	return commitObject, nil
}
