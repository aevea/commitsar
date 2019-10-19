package providers

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindCompareBranch(t *testing.T) {
	// Github action specific environment variable
	os.Setenv("GITHUB_BASE_REF", "github-develop")

	actionCompareBranch := FindCompareBranch()

	assert.Equal(t, "github-develop", actionCompareBranch)

	os.Setenv("GITHUB_BASE_REF", "")

	// Gitlab specific environment variable
	os.Setenv("CI_MERGE_REQUEST_TARGET_BRANCH_NAME", "gitlab-develop")

	gitlabCompareBranch := FindCompareBranch()

	assert.Equal(t, "gitlab-develop", gitlabCompareBranch)

	os.Setenv("CI_MERGE_REQUEST_TARGET_BRANCH_NAME", "")

	// Should default to master if no conditions are satisfied
	defaultMaster := FindCompareBranch()

	assert.Equal(t, "origin/master", defaultMaster)
}
