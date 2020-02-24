package root_runner

import (
	history "github.com/outillage/git/v2"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

func commitsBetweenBranches(gitRepo *history.Git, options RunnerOptions) ([]plumbing.Hash, error) {
	var commits []plumbing.Hash

	currentBranch, currentBranchErr := gitRepo.CurrentBranch()
	if currentBranchErr != nil {
		return nil, currentBranchErr
	}

	sameBranch, err := IdentifySameBranch(currentBranch.Name().String(), options.UpstreamBranch, gitRepo)

	if err != nil {
		return nil, err
	}

	if sameBranch {
		commitsOnSameBranch, err := gitRepo.CommitsOnBranch(currentBranch.Hash())

		if err != nil {
			return nil, err
		}

		if options.AllCommits {
			return commitsOnSameBranch, nil
		}

		// If no limit is set then check just the last commits. This is to prevent breaking repositories that did not check commits before.
		if options.Limit == 0 {
			commits = append(commits, commitsOnSameBranch[0])
			return commits, nil
		}

		limit := options.Limit

		// The limit cannot be longer than the amount of commits found
		if limit > len(commitsOnSameBranch) {
			limit = len(commitsOnSameBranch)
		}

		for index := 0; index < limit; index++ {
			commits = append(commits, commitsOnSameBranch[index])
		}

		return commits, nil
	}

	commitsOnBranch, err := gitRepo.BranchDiffCommits(currentBranch.Name().String(), options.UpstreamBranch)

	if err != nil {
		return nil, err
	}

	commits = commitsOnBranch

	return commits, nil
}
