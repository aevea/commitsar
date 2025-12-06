package commitpipeline

import (
	"fmt"

	history "github.com/aevea/git/v4"
	"github.com/apex/log"
)

// logBranch outputs the branch which is being checked into the console
func (pipeline *Pipeline) logBranch(gitRepo *history.Git) error {
	log.Debug("Attempting to get current branch...")

	branch, err := gitRepo.CurrentBranch()

	if err != nil {
		log.Errorf("Failed to get current branch: %v", err)
		return fmt.Errorf("failed to get current branch (this often indicates a detached HEAD state or shallow clone): %w", err)
	}

	log.Debugf("Current branch name: %s, hash: %s", branch.Name(), branch.Hash())

	// Detect and log detached HEAD state
	if branch.Name() == "HEAD" {
		log.Warn("Running in detached HEAD state. This is common in CI environments.")
		log.Debugf("Detached HEAD at commit: %s", branch.Hash())
	}

	log.Infof("Starting analysis of commits on branch %s", branch.Name())

	return nil
}
