package history

import (
	"fmt"

	"gopkg.in/src-d/go-git.v4/plumbing"
)

// BranchDiffCommits compares commits from 2 branches and returns of a diff of them.
func (g *Git) BranchDiffCommits(branchA string, branchB string) ([]plumbing.Hash, error) {
	branchACommit, err := g.LatestCommitOnBranch(branchA)

	if err != nil {
		return nil, fmt.Errorf("Failed getting latest commit for branch %v: %v", branchA, err)
	}

	branchBCommit, err := g.LatestCommitOnBranch(branchB)

	if err != nil {
		return nil, fmt.Errorf("Failed getting latest commit for branch %v: %v", branchB, err)
	}

	branchACommits, err := g.CommitsOnBranch(branchACommit.Hash)

	if err != nil {
		return nil, fmt.Errorf("Failed getting commits for branch %v: %v", branchA, err)
	}

	branchBCommits, err := g.CommitsOnBranch(branchBCommit.Hash)

	if err != nil {
		return nil, fmt.Errorf("Failed getting commits for branch %v: %v", branchB, err)
	}

	var diffCommits []plumbing.Hash

	for _, commit := range branchACommits {
		if !contains(branchBCommits, commit) {
			diffCommits = append(diffCommits, commit)
		}
	}

	return diffCommits, nil
}

func contains(s []plumbing.Hash, e plumbing.Hash) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
