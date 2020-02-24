package text

import "regexp"

var initCommitRegex = regexp.MustCompile(`^Initial .+`)

// IsInitialCommit checks if a commit needs to be filtered, because it is the init commit of any given repo
func IsInitialCommit(commitMessage string) bool{
	initCommitMatch := initCommitRegex.FindStringSubmatch(commitMessage)

	return initCommitMatch != nil
}
