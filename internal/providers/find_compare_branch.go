package providers

import (
	"fmt"
	"os"
)

// FindCompareBranch tries to find the merge request compare branch based on environment variables used by different CIs.
func FindCompareBranch() string {
	githubRef := os.Getenv("GITHUB_BASE_REF")

	if githubRef != "" {
		return fmt.Sprintf("origin/%v", githubRef)
	}

	gitlabRef := os.Getenv("CI_MERGE_REQUEST_TARGET_BRANCH_NAME")
	if gitlabRef != "" {
		return fmt.Sprintf("origin/%v", gitlabRef)
	}

	return "origin/master"
}
