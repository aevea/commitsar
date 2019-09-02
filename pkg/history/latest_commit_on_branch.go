package history

import (
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

// LatestCommitOnBranch resolves a revision and then returns the latest commit on it.
func (g *Git) LatestCommitOnBranch(desiredBranch string) (*object.Commit, error) {
	desiredHash, err := g.repo.ResolveRevision(plumbing.Revision(desiredBranch))

	if err != nil {
		return nil, err
	}

	desiredCommit, err := g.repo.CommitObject(*desiredHash)

	if err != nil {
		return nil, err
	}

	return desiredCommit, nil
}
