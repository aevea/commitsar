package text

import (
	"regexp"
)

var mergeCommitRegex = regexp.MustCompile(`^Merge commit '(?P<hash>\S+)'`)
var mergeBranchRegex = regexp.MustCompile(`^Merge branch '(?P<incoming>\w+)' into (?P<current>\S+)`)
var kodiakMergeBranchRegex = regexp.MustCompile(`^Merge (?P<incoming>\w+) into (?P<current>\S+)`)

// IsMergeCommit tests message string against expected format of a merge commit and returns true/false based on it
func IsMergeCommit(message string) bool {
	mergeCommitMatch := mergeCommitRegex.FindStringSubmatch(message)

	mergeBranchMatch := mergeBranchRegex.FindStringSubmatch(message)

	kodiakMergeBranchMatch := kodiakMergeBranchRegex.FindStringSubmatch(message)

	if mergeCommitMatch != nil || mergeBranchMatch != nil || kodiakMergeBranchMatch != nil {
		return true
	}

	return false
}
