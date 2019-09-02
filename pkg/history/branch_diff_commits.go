package history

import (
	"fmt"

	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

// BranchDiffCommits finds the common ancestors between two branches
// Flow:
// 1. Find latest commit on branchA
// 2. Find latest commit on branchB
// 3. Run MergeBase against them
func (g *Git) BranchDiffCommits(branchA string, branchB string) ([]*object.Commit, error) {
	branchACommit, err := g.latestCommitOnBranch(branchA)

	if err != nil {
		return nil, err
	}

	branchBCommit, err := g.latestCommitOnBranch(branchB)

	if err != nil {
		return nil, err
	}

	diffCommits, mergeBaseErr := branchACommit.MergeBase(branchBCommit)

	if mergeBaseErr != nil {
		return nil, mergeBaseErr
	}

	if g.Debug {
		fmt.Printf("\n Following mergeCommits found %v", diffCommits)
	}

	return diffCommits, nil
}

func (g *Git) latestCommitOnBranch(desiredBranch string) (*object.Commit, error) {
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
