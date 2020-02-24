package root_runner

import (
	"errors"

	"github.com/logrusorgru/aurora"
	"github.com/outillage/commitsar/pkg/text"
	history "github.com/outillage/git/v2"
	"github.com/outillage/quoad"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

// Run executes the base command for Commitsar
func (runner *Runner) Run(options RunnerOptions, args ...string) error {
	runner.Logger.Print("Starting analysis of commits on branch\n")

	gitRepo, err := history.OpenGit(options.Path, runner.DebugLogger)

	if err != nil {
		return err
	}

	var commits []plumbing.Hash

	if len(args) == 0 {
		commitsBetweenBranches, err := commitsBetweenBranches(gitRepo, options.UpstreamBranch)

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

		runner.DebugLogger.Printf("Commit found: [hash] %v [message] %v \n", commitObject.Hash, text.MessageTitle(commitObject.Message))

		if commitErr != nil {
			return commitErr
		}

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

		messageTitle := text.MessageTitle(commitObject.Message)

		parsedCommit := quoad.ParseCommitMessage(commitObject.Message)

		textErr := text.CheckMessageTitle(parsedCommit, runner.Strict)

		if textErr != nil {
			faultyCommits = append(faultyCommits, text.FailingCommit{Hash: commitHash.String(), Message: messageTitle, Error: textErr})
		}
	}

	if len(faultyCommits) != 0 {
		failingCommitMessage := text.FormatFailingCommits(faultyCommits)

		runner.Logger.Print(failingCommitMessage)

		runner.Logger.Printf("%v of %v commits are not conventional commit compliant\n", aurora.Red(len(faultyCommits)), aurora.Red(len(commits)))

		runner.Logger.Print("\nExpected format is for example:      chore(ci): this is a test\n")
		runner.Logger.Print("Please see https://www.conventionalcommits.org for help on how to structure commits\n\n")

		return errors.New(aurora.Red("Not all commits are conventional commits, please check the commits listed above").String())
	}

	runner.Logger.Print(aurora.Sprintf(aurora.Green("All %v commits are conventional commit compliant\n"), len(filteredCommits)))

	return nil
}
