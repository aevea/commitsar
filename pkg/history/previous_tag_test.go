package history

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/src-d/go-git.v4"
)

func TestPreviousTag(t *testing.T) {
	repo, _ := git.PlainOpen("../../testdata/git_tags")
	testGit := &Git{repo: repo, Debug: true}

	head, err := repo.Head()

	assert.NoError(t, err)

	tagHash, err := testGit.PreviousTag(head.Hash())

	assert.NoError(t, err)

	log.Print(tagHash)

	commit, err := repo.CommitObject(tagHash)
	assert.NoError(t, err)
	assert.Equal(t, "chore: first commit on master\n", commit.Message)

}
