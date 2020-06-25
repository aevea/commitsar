package commitpipeline

import (
	"github.com/aevea/commitsar/internal/dispatcher"
	"github.com/aevea/commitsar/pkg/text"
	history "github.com/aevea/git/v2"
	"github.com/aevea/quoad"
)

// Run runs the CommitPipeline pushing errors to the errChannel or returning the successMessage
func (pipeline *CommitPipeline) Run(errChannel chan dispatcher.PipelineError) *dispatcher.PipelineSuccess {

	gitRepo, err := history.OpenGit(".", pipeline.debugLogger)

	if err != nil {
		errChannel <- dispatcher.PipelineError{
			PipelineName: "commit",
			Error:        err,
		}
		return nil
	}

	commits, err := pipeline.getCommits(gitRepo)

	if err != nil {
		errChannel <- dispatcher.PipelineError{
			PipelineName: "commit",
			Error:        err,
		}
		return nil
	}

	for _, commit := range commits {
		parsedCommit := quoad.ParseCommitMessage(commit.Message)

		textErr := text.CheckMessageTitle(parsedCommit, true)

		if textErr != nil {
			errChannel <- dispatcher.PipelineError{
				PipelineName: pipeline.Name(),
				Data: []dispatcher.FailureData{
					{Name: "Commit", Value: parsedCommit.Heading},
				},
				Error: textErr,
			}
		}
	}

	return nil
}
