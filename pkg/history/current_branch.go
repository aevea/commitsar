package history

import (
	"gopkg.in/src-d/go-git.v4/plumbing"
)


// CurrentBranch returns the reference HEAD is at right now
func (g *Git) CurrentBranch() (*plumbing.Reference, error) {
	head, err := g.repo.Head()

	if err != nil {
		return nil, err
	}

	return head, nil
}