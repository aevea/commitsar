package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommitsOnMaster(t *testing.T) {
	err := runCommitsar("../testdata/commits-on-master", "master", false, true)

	assert.NoError(t, err)
}

func TestCommitsOnBranch(t *testing.T) {
	err := runCommitsar("../testdata/commits-on-different-branches", "master", false, true)

	assert.Error(t, err)
}
