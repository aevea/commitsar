package commitpipeline

import (
	"strings"

	history "github.com/aevea/git/v4"
)

func commitsBetweenHashes(gitRepo *history.Git, args []string) ([]history.Hash, error) {
	var commits []history.Hash
	var fromCommit history.Hash
	var toCommit history.Hash

	arg := args[0]

	splitArgs := strings.Split(arg, "...")

	if len(splitArgs) == 1 {
		currentCommit, err := gitRepo.CurrentCommit()

		if err != nil {
			return nil, err
		}

		fromCommit = currentCommit.Hash

		toCommit, err = history.NewHash(splitArgs[0])
		if err != nil {
			return nil, err
		}
	}

	if len(splitArgs) == 2 {
		var err error
		fromCommit, err = history.NewHash(splitArgs[1])
		if err != nil {
			return nil, err
		}
		toCommit, err = history.NewHash(splitArgs[0])
		if err != nil {
			return nil, err
		}
	}

	logCommits, err := gitRepo.CommitsBetween(fromCommit, toCommit)

	if err != nil {
		return nil, err
	}

	commits = logCommits

	return commits, nil
}
