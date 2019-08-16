package text

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatFailingCommits(t *testing.T) {
	var commits []FailingCommit

	commits = append(commits, FailingCommit{Hash: "testhash", Message: "chore:add seomthing"}, FailingCommit{Hash: "testhash2", Message: "broken"})

	formattedText := FormatFailingCommits(commits)
	assert.Equal(t, formattedText, "\nFollowing commits failed the check: \n\nFAIL   testhash   chore:add seomthing\nFAIL   testhash2   broken\n\n")
}
