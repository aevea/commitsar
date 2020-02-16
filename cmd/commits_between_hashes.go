package cmd

import (
	"strings"

	history "github.com/outillage/git/pkg"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

func commitsBetweenHashes(gitRepo *history.Git, args []string) ([]plumbing.Hash, error) {
	var commits []plumbing.Hash
	var fromCommit plumbing.Hash
	var toCommit plumbing.Hash

	arg := args[0]

	splitArgs := strings.Split(arg, "...")

	if len(splitArgs) == 1 {
		currentCommit, err := gitRepo.CurrentCommit()

		if err != nil {
			return nil, err
		}

		fromCommit = currentCommit.Hash

		toCommit = plumbing.NewHash(splitArgs[0])
	}

	if len(splitArgs) == 2 {
		fromCommit = plumbing.NewHash(splitArgs[0])
		toCommit = plumbing.NewHash(splitArgs[1])
	}

	logCommits, err := gitRepo.CommitsBetween(fromCommit, toCommit)

	if err != nil {
		return nil, err
	}

	commits = logCommits

	return commits, nil
}
