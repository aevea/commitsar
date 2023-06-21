package commitpipeline

import (
	"errors"
	"os"

	"github.com/aevea/commitsar/internal/dispatcher"
	"github.com/aevea/commitsar/pkg/text"
	history "github.com/aevea/git/v3"
	"github.com/aevea/quoad"
	"github.com/apex/log"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/logrusorgru/aurora"
)

func (pipeline *Pipeline) Run() (*dispatcher.PipelineSuccess, error) {
	gitRepo, err := history.OpenGit(pipeline.options.Path)

	if err != nil {
		return nil, err
	}

	err = pipeline.logBranch(gitRepo)

	if err != nil {
		return nil, err
	}

	var commits []plumbing.Hash

	if len(pipeline.args) == 0 {
		commitsBetweenBranches, err := pipeline.commitsBetweenBranches(gitRepo)

		if err != nil {
			return nil, err
		}

		commits = commitsBetweenBranches
	} else {
		commitsBetweenHashes, err := commitsBetweenHashes(gitRepo, pipeline.args)

		if err != nil {
			return nil, err
		}

		commits = commitsBetweenHashes
	}

	var filteredCommits []plumbing.Hash

	for _, commitHash := range commits {
		commitObject, commitErr := gitRepo.Commit(commitHash)

		if commitErr != nil {
			return nil, commitErr
		}

		parsedCommit := quoad.ParseCommitMessage(commitObject.Message)

		log.Debugf("Commit found: [hash] %v [message] %v", commitObject.Hash, parsedCommit.Heading)

		if !text.IsMergeCommit(commitObject.Message) && !text.IsInitialCommit(commitObject.Message) && !text.IsRevertCommit(commitObject.Message) {
			filteredCommits = append(filteredCommits, commitHash)
		}
	}

	log.Infof("%v commits filtered out", len(commits)-len(filteredCommits))
	log.Infof("Found %v commit to check", len(filteredCommits))

	if len(filteredCommits) == 0 && len(commits) > 0 {
		return &dispatcher.PipelineSuccess{
			Message:      aurora.Sprintf(aurora.Green("All commits have been filtered out. Nothing to check."), len(commits)),
			PipelineName: pipeline.Name(),
		}, nil
	}
	if len(filteredCommits) == 0 {
		return nil, errors.New(aurora.Red("No commits found, please check you are on a branch outside of main").String())
	}

	var faultyCommits []text.FailingCommit
	requiredScopeChecker := text.RequiredScopeChecker(pipeline.options.RequiredScopes)

	for _, commitHash := range filteredCommits {
		commitObject, commitErr := gitRepo.Commit(commitHash)

		if commitErr != nil {
			return nil, commitErr
		}

		parsedCommit := quoad.ParseCommitMessage(commitObject.Message)

		textErr := text.CheckMessageTitle(parsedCommit, pipeline.options.Strict)

		if textErr != nil {
			faultyCommits = append(faultyCommits, text.FailingCommit{Hash: commitHash.String(), Message: parsedCommit.Heading, Error: textErr})
			continue
		}

		scopeErr := requiredScopeChecker(parsedCommit.Scope)

		if scopeErr != nil {
			faultyCommits = append(faultyCommits, text.FailingCommit{Hash: commitHash.String(), Message: parsedCommit.Heading, Error: scopeErr})
		}

	}

	if len(faultyCommits) != 0 {
		failingCommitTable := text.FormatFailingCommits(faultyCommits)
		failingCommitTable.SetOutputMirror(os.Stdout)
		failingCommitTable.Render()

		log.Infof("%v of %v commits are not conventional commit compliant", aurora.Red(len(faultyCommits)), aurora.Red(len(commits)))

		log.Info("Expected format is for example:      chore(ci): this is a test")
		log.Info("Please see https://www.conventionalcommits.org for help on how to structure commits")

		return nil, errors.New(aurora.Red("Not all commits are conventional commits, please check the commits listed above").String())
	}

	return &dispatcher.PipelineSuccess{
		Message:      aurora.Sprintf(aurora.Green("All %v commits are conventional commit compliant"), len(filteredCommits)),
		PipelineName: pipeline.Name(),
	}, nil
}
