package history

import (
	"gopkg.in/src-d/go-git.v4"
)

// Git is the struct used to house all methods in use in Commitsar.
type Git struct {
	repo *git.Repository
	// Debug flag is passed to make debugging easier during development/problematic deploys
	Debug bool
}

// OpenGit loads Repo on path and returns a new Git struct to work with.
func OpenGit(path string, debug bool) (*Git, error) {
	repo, repoErr := Repo(path)

	if repoErr != nil {
		return nil, repoErr
	}

	return &Git{repo: repo, Debug: debug}, nil
}
