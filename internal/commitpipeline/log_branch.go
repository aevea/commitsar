package commitpipeline

import (
	history "github.com/aevea/git/v3"
	"github.com/apex/log"
)

// logBranch outputs the branch which is being checked into the console
func (pipeline *Pipeline) logBranch(gitRepo *history.Git) error {
	branch, err := gitRepo.CurrentBranch()

	if err != nil {
		return err
	}

	log.Infof("Starting analysis of commits on branch %s", branch.Name())

	return nil
}
