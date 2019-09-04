package history

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCurrentBranch(t *testing.T) {
	repo := setupRepo()
	createTestHistory(repo)

	testGit := &Git{repo: repo}

	currentBranch, err := testGit.CurrentBranch()

	assert.NoError(t, err)
	assert.Equal(t, "refs/heads/my-branch", currentBranch.Name().String())
}
