package history

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/src-d/go-git.v4"
)

func TestCommitsBetween(t *testing.T) {
	repo, _ := git.PlainOpen("../../testdata/git_tags")
	testGit := &Git{repo: repo, Debug: false}

	head, err := repo.Head()

	assert.NoError(t, err)

	tagHash, err := testGit.PreviousTag(head.Hash())

	assert.NoError(t, err)

	log.Print(tagHash)

	commit, err := repo.CommitObject(tagHash)
	assert.NoError(t, err)
	assert.Equal(t, "chore: first commit on master\n", commit.Message)

	commits, err := testGit.CommitsBetween(head.Hash(), tagHash)

	assert.NoError(t, err)
	assert.Len(t, commits, 3)

	middleCommit, _ := repo.CommitObject(commits[1])

	assert.Equal(t, "chore: third commit on master\n", middleCommit.Message)
}
