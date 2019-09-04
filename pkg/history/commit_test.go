package history

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommit(t *testing.T) {
	repo := setupRepo()
	createTestHistory(repo)

	testGit := &Git{repo: repo}

	head, _ := testGit.CurrentCommit()

	commit, err := testGit.Commit(head.Hash)
	assert.NoError(t, err)
	assert.Equal(t, "third commit on new branch", commit.Message)
	assert.NotEmpty(t, commit.Hash)
}
