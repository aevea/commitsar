package commitpipeline

import history "github.com/aevea/git/v2"

// logBranch outputs the branch which is being checked into the console
func (pipeline *Pipeline) logBranch(gitRepo *history.Git) error {
	branch, err := gitRepo.CurrentBranch()

	if err != nil {
		return err
	}

	pipeline.Logger.Printf("Starting analysis of commits on branch %s", branch.Name())

	return nil
}
