package root_runner

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommitsOnMaster(t *testing.T) {
	err := RunCommitsar("../../testdata/commits-on-master", "master", false, true)

	assert.NoError(t, err)
}

func TestCommitsOnBranch(t *testing.T) {
	err := RunCommitsar("../../testdata/commits-on-different-branches", "master", false, true)

	assert.Error(t, err)
}

func TestFromToCommits(t *testing.T) {
	err := RunCommitsar("../../testdata/commits-on-different-branches", "master", false, true, "7dbf3e7db93ae2e02902cae9d2f1de1b1e5c8c92...d0240d3ed34685d0a5329b185e120d3e8c205be4")

	assert.NoError(t, err)
}
