package commitpipeline

import (
	history "github.com/aevea/git/v3"
	"github.com/go-git/go-git/v5/plumbing"
)

func (pipeline *Pipeline) commitsBetweenBranches(gitRepo *history.Git) ([]plumbing.Hash, error) {
	var commits []plumbing.Hash

	currentBranch, currentBranchErr := gitRepo.CurrentBranch()
	if currentBranchErr != nil {
		return nil, currentBranchErr
	}

	sameBranch, err := IdentifySameBranch(currentBranch.Name().String(), pipeline.options.UpstreamBranch, gitRepo)

	if err != nil {
		return nil, err
	}

	if sameBranch {
		commitsOnSameBranch, err := gitRepo.CommitsOnBranch(currentBranch.Hash())

		if err != nil {
			return nil, err
		}

		if pipeline.options.AllCommits {
			return commitsOnSameBranch, nil
		}

		// If no limit is set then check just the last commits. This is to prevent breaking repositories that did not check commits before.
		if pipeline.options.Limit == 0 {
			commits = append(commits, commitsOnSameBranch[0])
			return commits, nil
		}

		limit := pipeline.options.Limit

		// The limit cannot be longer than the amount of commits found
		if limit > len(commitsOnSameBranch) {
			limit = len(commitsOnSameBranch)
		}

		for index := 0; index < limit; index++ {
			commits = append(commits, commitsOnSameBranch[index])
		}

		return commits, nil
	}

	commitsOnBranch, err := gitRepo.BranchDiffCommits(currentBranch.Name().String(), pipeline.options.UpstreamBranch)

	if err != nil {
		return nil, err
	}

	commits = commitsOnBranch

	return commits, nil
}
