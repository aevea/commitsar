package git

import (
	"testing"

	"time"

	"github.com/stretchr/testify/assert"
	"gopkg.in/src-d/go-billy.v4/memfs"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/storage/memory"
)

func setupRepo() *git.Repository {
	repo, _ := git.Init(memory.NewStorage(), memfs.New())

	return repo
}

func createCommit(repo *git.Repository, message string) *object.Commit {
	w, _ := repo.Worktree()

	commit, _ := w.Commit(message, &git.CommitOptions{
		Author: &object.Signature{
			Name:  "John Doe",
			Email: "john@doe.org",
			When:  time.Now(),
		},
	})

	obj, _ := repo.CommitObject(commit)

	return obj
}

func TestCommitMessage(t *testing.T) {
	repo := setupRepo()
	commit := createCommit(repo, "example commit")

	message, err := CommitMessage(repo, commit.Hash)

	assert.Equal(t, message, "example commit")
	assert.Equal(t, err, nil)
}
