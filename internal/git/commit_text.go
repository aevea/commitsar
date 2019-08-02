package git

import (
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

// CommitMessage returns message of given commit
func CommitMessage(repo *git.Repository, commit plumbing.Hash) (string, error) {
	commitObject, err := repo.CommitObject(commit)

	if err != nil {
		return "", err
	}

	return commitObject.Message, nil
}
