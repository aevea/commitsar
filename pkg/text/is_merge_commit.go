package text

import (
	"regexp"
)

var mergeCommitRegex = regexp.MustCompile(`^Merge commit '(?P<hash>\w+)'`)
var mergeBranchRegex = regexp.MustCompile(`^Merge branch '(?P<incoming>\w+)' into (?P<current>\w+)`)

// IsMergeCommit tests message string against expected format of a merge commit and returns true/false based on it
func IsMergeCommit(message string) bool {
	mergeCommitMatch := mergeCommitRegex.FindStringSubmatch(message)

	mergeBranchMatch := mergeBranchRegex.FindStringSubmatch(message)

	if mergeCommitMatch != nil || mergeBranchMatch != nil {
		return true
	}

	return false
}
