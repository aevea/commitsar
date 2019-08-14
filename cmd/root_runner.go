package cmd

import (
	"errors"
	"log"

	"github.com/commitsar-app/commitsar/internal/history"
	"github.com/commitsar-app/commitsar/internal/text"
	"github.com/spf13/cobra"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

func runRoot(cmd *cobra.Command, args []string) error {
	debug := false
	if cmd.Flag("verbose").Value.String() == "true" {
		debug = true
	}

	log.Print("Starting analysis of commits on branch")

	repo, repoErr := history.Repo(".")

	if repoErr != nil {
		return repoErr
	}

	currentBranch, currentBranchErr := repo.Head()

	if debug {
		log.Printf("Current branch %v", currentBranch.Name().String())
		refIter, _ := repo.References()

		refIterErr := refIter.ForEach(func(ref *plumbing.Reference) error {
			log.Printf("[REF] %v", ref.Name().String())
			return nil
		})

		if refIterErr != nil {
			return refIterErr
		}
	}

	if currentBranchErr != nil {
		return currentBranchErr
	}

	commits, commitsErr := history.CommitsOnBranch(repo, currentBranch.Hash(), "origin/master")

	if len(commits) == 0 {
		return errors.New("No commits found, please check you are on a branch outside of main")
	}

	if commitsErr != nil {
		return commitsErr
	}

	log.Printf("Found %v commit to check", len(commits))

	faultyCommits := []plumbing.Hash{}

	for _, commitHash := range commits {
		commitObject, commitErr := repo.CommitObject(commitHash)

		if commitErr != nil {
			return commitErr
		}

		textErr := text.CheckMessageTitle(commitObject.Message)

		if textErr != nil {
			faultyCommits = append(faultyCommits, commitHash)
		}
	}

	if len(faultyCommits) != 0 {
		for _, commitHash := range faultyCommits {
			log.Printf("Commit %v is not conventional commit compliant", commitHash)
		}

		log.Printf("%v of %v commits are not conventional commit compliant", len(faultyCommits), len(commits))

		return errors.New("Not all commits are conventiontal commits, please check the commits listed above")
	}

	log.Printf("All %v commits are conventional commit compliant", len(commits))

	return nil
}
