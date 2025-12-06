package commitpipeline

import (
	"fmt"

	history "github.com/aevea/git/v4"
	"github.com/apex/log"
)

func (pipeline *Pipeline) commitsBetweenBranches(gitRepo *history.Git) ([]history.Hash, error) {
	var commits []history.Hash

	log.Debug("Getting current branch for commit analysis...")
	currentBranch, currentBranchErr := gitRepo.CurrentBranch()
	if currentBranchErr != nil {
		log.Errorf("Failed to get current branch in commitsBetweenBranches: %v", currentBranchErr)
		return nil, fmt.Errorf("failed to get current branch: %w", currentBranchErr)
	}

	log.Debugf("Current branch: %s, upstream branch: %s", currentBranch.Name(), pipeline.options.UpstreamBranch)

	log.Debug("Checking if current branch is the same as upstream...")
	sameBranch, err := IdentifySameBranch(currentBranch.Name(), pipeline.options.UpstreamBranch, gitRepo)

	if err != nil {
		log.Errorf("Failed to identify if same branch: %v", err)
		return nil, fmt.Errorf("failed to identify if on same branch as upstream '%s': %w", pipeline.options.UpstreamBranch, err)
	}

	log.Debugf("Same branch as upstream: %v", sameBranch)

	if sameBranch {
		log.Debugf("On same branch, getting commits on branch with hash: %s", currentBranch.Hash())
		commitsOnSameBranch, err := gitRepo.CommitsOnBranch(currentBranch.Hash())

		if err != nil {
			log.Errorf("Failed to get commits on branch: %v", err)
			return nil, fmt.Errorf("failed to get commits on branch '%s': %w", currentBranch.Name(), err)
		}

		log.Debugf("Found %d commits on branch", len(commitsOnSameBranch))

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

	log.Debugf("On different branch, getting diff commits between '%s' and '%s'", currentBranch.Name(), pipeline.options.UpstreamBranch)
	commitsOnBranch, err := gitRepo.BranchDiffCommits(currentBranch.Name(), pipeline.options.UpstreamBranch)

	if err != nil {
		log.Errorf("Failed to get branch diff commits between '%s' and '%s': %v", currentBranch.Name(), pipeline.options.UpstreamBranch, err)
		return nil, fmt.Errorf("failed to get commits between branch '%s' and upstream '%s' (ensure full git history is available, not a shallow clone): %w", currentBranch.Name(), pipeline.options.UpstreamBranch, err)
	}

	log.Debugf("Found %d commits in branch diff", len(commitsOnBranch))
	commits = commitsOnBranch

	return commits, nil
}
