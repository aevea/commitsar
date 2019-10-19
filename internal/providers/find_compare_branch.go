package providers

import "os"

// FindCompareBranch tries to find the merge request compare branch based on environment variables used by different CIs.
func FindCompareBranch() string {
	if os.Getenv("GITHUB_BASE_REF") != "" {
		return os.Getenv("GITHUB_BASE_REF")
	}

	if os.Getenv("CI_MERGE_REQUEST_TARGET_BRANCH_NAME") != "" {
		return os.Getenv("CI_MERGE_REQUEST_TARGET_BRANCH_NAME")
	}

	return "origin/master"
}
