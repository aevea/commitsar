package git

import (
	"gopkg.in/src-d/go-git.v4"
)

// Repo accepts path to the repo and returns a repository struct
func Repo(path string) (*git.Repository, error) {
	repo, repoErr := git.PlainOpen(path)

	if repoErr != nil {
		return nil, repoErr
	}

	return repo, repoErr
}
