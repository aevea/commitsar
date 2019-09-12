package text

import (
	"regexp"
)

var mergeCommitRegex = regexp.MustCompile(`^Merge .+`)

// IsMergeCommit tests message string against expected format of a merge commit and returns true/false based on it
func IsMergeCommit(message string) bool {
	mergeCommitMatch := mergeCommitRegex.FindStringSubmatch(message)

	return mergeCommitMatch != nil
}
