package text

import (
	"strings"

	"github.com/jedib0t/go-pretty/table"
)

// FailingCommit is just a formatted commit struct
type FailingCommit struct {
	Hash    string
	Message string
	Error   error
}

// FormatFailingCommits takes in slice of commit hashes and messages and formats it for nice output
func FormatFailingCommits(commits []FailingCommit) table.Writer {
	t := table.NewWriter()
	t.AppendHeader(table.Row{"hash", "failure", "text"})

	builder := strings.Builder{}
	// Extra spacing to make it nicer
	builder.WriteString("\nFollowing commits failed the check: \n")

	for _, commit := range commits {
		t.AppendRow(table.Row{commit.Hash, commit.Error.Error(), commit.Message})
	}

	return t
}
