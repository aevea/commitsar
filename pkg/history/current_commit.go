package history

import (
	"fmt"

	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

// CurrentCommit returns the commit that HEAD is at
func (g *Git) CurrentCommit() (*object.Commit, error) {
	head, err := g.repo.Head()

	if err != nil {
		return nil, err
	}

	currentCommitHash := head.Hash()

	if g.Debug {
		fmt.Printf("\ncurrent commitHash: %v \n", currentCommitHash)
	}

	commitObject, err := g.repo.CommitObject(currentCommitHash)

	if err != nil {
		return nil, err
	}

	return commitObject, nil
}
