package runners

import (
	"testing"

	history "github.com/outillage/git/pkg"
	"github.com/stretchr/testify/assert"
)

func TestCommitsBetweenHashes(t *testing.T) {
	gitRepo, err := history.OpenGit("../../testdata/commits-on-different-branches", false)

	assert.NoError(t, err)

	commits, err := commitsBetweenHashes(gitRepo, []string{"7dbf3e7db93ae2e02902cae9d2f1de1b1e5c8c92...d0240d3ed34685d0a5329b185e120d3e8c205be4"})

	// TODO: Allow including to commit
	assert.Len(t, commits, 1)
	assert.NoError(t, err)
}

func TestCommitsBetweenHashesOnlyTo(t *testing.T) {
	gitRepo, err := history.OpenGit("../../testdata/commits-on-different-branches", false)

	assert.NoError(t, err)

	commits, err := commitsBetweenHashes(gitRepo, []string{"d0240d3ed34685d0a5329b185e120d3e8c205be4"})

	assert.Len(t, commits, 2)
	assert.NoError(t, err)
}
