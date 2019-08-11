package cmd

import (
	"log"

	"github.com/fallion/commitsar/internal/history"
	"github.com/fallion/commitsar/internal/text"
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

	masterRef := plumbing.ReferenceName("master")

	commits, commitsErr := history.CommitsOnBranch(repo, currentBranch.Name(), masterRef)

	if commitsErr != nil {
		return commitsErr
	}

	log.Printf("Found %v commit to check", len(commits))

	for _, commitHash := range commits {
		commitObject, commitErr := repo.CommitObject(commitHash)

		if commitErr != nil {
			return commitErr
		}

		textErr := text.CheckMessageTitle(commitObject.Message)

		if textErr != nil {
			return textErr
		}
	}

	log.Printf("All %v commits are conventional commit compliant", len(commits))

	return nil
}
