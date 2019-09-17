package cmd

import (
	"errors"
	"fmt"

	"github.com/commitsar-app/commitsar/pkg/history"
	"github.com/commitsar-app/commitsar/pkg/text"
	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

func runRoot(cmd *cobra.Command, args []string) error {
	debug := false
	if cmd.Flag("verbose").Value.String() == "true" {
		debug = true
	}

	fmt.Print("Starting analysis of commits on branch\n")

	gitRepo, err := history.OpenGit(".", debug)

	if err != nil {
		return err
	}

	currentBranch, currentBranchErr := gitRepo.CurrentBranch()
	if currentBranchErr != nil {
		return currentBranchErr
	}

	commits, commitsErr := gitRepo.BranchDiffCommits(currentBranch.Name().String(), "origin/master")

	if commitsErr != nil {
		return commitsErr
	}

	var filteredCommits []plumbing.Hash

	for _, commitHash := range commits {
		commitObject, commitErr := gitRepo.Commit(commitHash)

		if debug {
			fmt.Printf("\n[DEBUG] Commit found: [hash] %v [message] %v \n", commitObject.Hash, text.MessageTitle(commitObject.Message))
		}

		if commitErr != nil {
			return commitErr
		}

		if !text.IsMergeCommit(commitObject.Message) {
			filteredCommits = append(filteredCommits, commitHash)
		}
	}

	fmt.Printf("\n%v commits filtered out\n", len(commits)-len(filteredCommits))
	fmt.Printf("\nFound %v commit to check\n", len(filteredCommits))

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

		parsedCommit := text.ParseCommit(commitObject.Message, commitHash)

		textErr := text.CheckMessageTitle(parsedCommit)

		if textErr != nil {
			faultyCommits = append(faultyCommits, text.FailingCommit{Hash: commitHash.String(), Message: messageTitle, Error: textErr})
		}
	}

	if len(faultyCommits) != 0 {
		failingCommitMessage := text.FormatFailingCommits(faultyCommits)

		fmt.Print(failingCommitMessage)

		fmt.Printf("%v of %v commits are not conventional commit compliant\n", aurora.Red(len(faultyCommits)), aurora.Red(len(commits)))

		fmt.Print("\nExpected format is for example:      chore(ci): this is a test\n")
		fmt.Print("Please see https://www.conventionalcommits.org for help on how to structure commits\n\n")

		return errors.New(aurora.Red("Not all commits are conventiontal commits, please check the commits listed above").String())
	}

	fmt.Print(aurora.Sprintf(aurora.Green("All %v commits are conventional commit compliant\n"), len(filteredCommits)))

	return nil
}
