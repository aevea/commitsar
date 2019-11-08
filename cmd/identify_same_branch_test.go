package cmd

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIdentifySameBranchTrue(t *testing.T) {
	tests := map[string]string{
		"refs/heads/master": "master",
		"origin/master":     "refs/heads/master",
	}

	for branchA, branchB := range tests {
		assert.Equal(t, true, IdentifySameBranch(branchA, branchB), fmt.Sprintf("Branch %v should be equal to branch %v", branchA, branchB))
	}
}

func TestIdentifySameBranchFalse(t *testing.T) {
	tests := map[string]string{
		"refs/heads/master":       "some-other-branch",
		"refs/heads/other-branch": "refs/heads/weird-branch",
	}

	for branchA, branchB := range tests {
		assert.Equal(t, false, IdentifySameBranch(branchA, branchB), fmt.Sprintf("Branch %v should NOT be equal to branch %v", branchA, branchB))
	}
}
