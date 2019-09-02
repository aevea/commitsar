package history

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLatestCommitOnBranch(t *testing.T) {
	repo := setupRepo()
	createTestHistory(repo)

	head, _ := repo.Head()

	testGit := &Git{repo: repo}

	commit, err := testGit.LatestCommitOnBranch(head.Name().String())

	assert.NoError(t, err)
	assert.Equal(t, "third commit on new branch", commit.Message)
	assert.Equal(t, err, nil)
}
