package root_runner

import (
	"errors"
	"os"

	"github.com/aevea/commitsar/pkg/text"
	history "github.com/aevea/git/v2"
	"github.com/aevea/quoad"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/logrusorgru/aurora"
)

// Run executes the base command for Commitsar
func (runner *Runner) Run(options RunnerOptions, args ...string) error {
	gitRepo, err := history.OpenGit(options.Path, runner.DebugLogger)

	if err != nil {
		return err
	}

	err = runner.logBranch(gitRepo)

	if err != nil {
		return err
	}

	var commits []plumbing.Hash

	if len(args) == 0 {
		commitsBetweenBranches, err := commitsBetweenBranches(gitRepo, options)

		if err != nil {
			return err
		}

		commits = commitsBetweenBranches
	} else {
		commitsBetweenHashes, err := commitsBetweenHashes(gitRepo, args)

		if err != nil {
			return err
		}

		commits = commitsBetweenHashes
	}

	var filteredCommits []plumbing.Hash

	for _, commitHash := range commits {
		commitObject, commitErr := gitRepo.Commit(commitHash)

		if commitErr != nil {
			return commitErr
		}

		parsedCommit := quoad.ParseCommitMessage(commitObject.Message)

		runner.DebugLogger.Printf("Commit found: [hash] %v [message] %v \n", parsedCommit.Hash.String(), parsedCommit.Heading)

		if !text.IsMergeCommit(commitObject.Message) && !text.IsInitialCommit(commitObject.Message) {
			filteredCommits = append(filteredCommits, commitHash)
		}
	}

	runner.Logger.Printf("\n%v commits filtered out\n", len(commits)-len(filteredCommits))
	runner.Logger.Printf("\nFound %v commit to check\n", len(filteredCommits))

	if len(filteredCommits) == 0 {
		return errors.New(aurora.Red("No commits found, please check you are on a branch outside of main").String())
	}

	var faultyCommits []text.FailingCommit

	for _, commitHash := range filteredCommits {
		commitObject, commitErr := gitRepo.Commit(commitHash)

		if commitErr != nil {
			return commitErr
		}

		parsedCommit := quoad.ParseCommitMessage(commitObject.Message)

		textErr := text.CheckMessageTitle(parsedCommit, options.Strict)

		if textErr != nil {
			faultyCommits = append(faultyCommits, text.FailingCommit{Hash: commitHash.String(), Message: parsedCommit.Heading, Error: textErr})
		}
	}

	if len(faultyCommits) != 0 {
		failingCommitTable := text.FormatFailingCommits(faultyCommits)
		failingCommitTable.SetOutputMirror(os.Stdout)
		failingCommitTable.Render()

		runner.Logger.Printf("%v of %v commits are not conventional commit compliant\n", aurora.Red(len(faultyCommits)), aurora.Red(len(commits)))

		runner.Logger.Print("\nExpected format is for example:      chore(ci): this is a test\n")
		runner.Logger.Print("Please see https://www.conventionalcommits.org for help on how to structure commits\n\n")

		return errors.New(aurora.Red("Not all commits are conventional commits, please check the commits listed above").String())
	}

	runner.Logger.Print(aurora.Sprintf(aurora.Green("All %v commits are conventional commit compliant\n"), len(filteredCommits)))

	return nil
}
