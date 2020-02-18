package root_runner

import (
	history "github.com/outillage/git/pkg"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

func commitsBetweenBranches(gitRepo *history.Git, upstreamBranch string) ([]plumbing.Hash, error) {
	var commits []plumbing.Hash

	currentBranch, currentBranchErr := gitRepo.CurrentBranch()
	if currentBranchErr != nil {
		return nil, currentBranchErr
	}

	sameBranch, err := IdentifySameBranch(currentBranch.Name().String(), upstreamBranch, gitRepo)

	if err != nil {
		return nil, err
	}

	if sameBranch {
		commitsOnSameBranch, err := gitRepo.CommitsOnBranch(currentBranch.Hash())

		if err != nil {
			return nil, err
		}

		commits = append(commits, commitsOnSameBranch[0])

	} else {
		commitsOnBranch, err := gitRepo.BranchDiffCommits(currentBranch.Name().String(), upstreamBranch)

		if err != nil {
			return nil, err
		}

		commits = commitsOnBranch
	}

	return commits, nil
}
