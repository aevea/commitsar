package text

import "strings"

func IsRevertCommit(commitMessage string) bool {
	if strings.HasPrefix(commitMessage, "Revert") {
		return true
	}

	return strings.HasPrefix(commitMessage, "revert")
}
