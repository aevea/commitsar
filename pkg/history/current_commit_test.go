package history

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCurrentCommit(t *testing.T) {
	repo := setupRepo()
	createTestHistory(repo)

	testGit := &Git{repo: repo}

	currentCommit, err := testGit.CurrentCommit()

	assert.NoError(t, err)
	assert.Equal(t, "third commit on new branch", currentCommit.Message)
}
