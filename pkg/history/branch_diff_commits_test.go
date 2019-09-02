package history

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBranchDiffCommits(t *testing.T) {
	repo := setupRepo()
	createTestHistory(repo)

	git := &Git{repo: repo}

	commits, err := git.BranchDiffCommits("my-branch", "master")

	assert.NoError(t, err)
	assert.Equal(t, "test commit on master", commits[0].Message)
}
