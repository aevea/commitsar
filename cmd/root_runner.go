package cmd

import (
	"log"

	"github.com/fallion/commitsar/internal/history"
	"github.com/fallion/commitsar/internal/text"
	"github.com/spf13/cobra"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

func runRoot(cmd *cobra.Command, args []string) {
	log.Print("Starting analysis of commits on branch")

	repo, repoErr := history.Repo(".")

	if repoErr != nil {
		panic(repoErr)
	}

	currentBranch, currentBranchErr := repo.Head()

	if currentBranchErr != nil {
		panic(currentBranchErr)
	}

	masterRef := plumbing.NewBranchReferenceName("master")

	commits, commitsErr := history.CommitsOnBranch(repo, currentBranch.Name(), masterRef)

	if commitsErr != nil {
		panic(commitsErr)
	}

	log.Printf("Found %v commit to check", len(commits))

	for _, commitHash := range commits {
		commitObject, commitErr := repo.CommitObject(commitHash)

		if commitErr != nil {
			panic(commitErr)
		}

		textErr := text.CheckMessageTitle(commitObject.Message)

		if textErr != nil {
			panic(textErr)
		}
	}

	log.Printf("All %v commits are conventional commit compliant", len(commits))
}
