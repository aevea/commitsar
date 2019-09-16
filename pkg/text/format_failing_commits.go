package text

import (
	"strings"
)

// FailingCommit is just a formatted commit struct
type FailingCommit struct {
	Hash    string
	Message string
	Error   error
}

// FormatFailingCommits takes in slice of commit hashes and messages and formats it for nice output
func FormatFailingCommits(commits []FailingCommit) string {
	builder := strings.Builder{}
	// Extra spacing to make it nicer
	builder.WriteString("\nFollowing commits failed the check: \n")

	for _, commit := range commits {
		builder.WriteString("\n")
		builder.WriteString("FAIL   ")
		builder.WriteString(commit.Hash)
		builder.WriteString("   ")
		builder.WriteString(commit.Message)
		builder.WriteString("   ")
		builder.WriteString(commit.Error.Error())
	}

	builder.WriteString("\n\n")
	return builder.String()
}
