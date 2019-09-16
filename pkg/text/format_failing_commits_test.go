package text

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatFailingCommits(t *testing.T) {
	var commits []FailingCommit

	commits = append(commits, FailingCommit{Hash: "testhash", Message: "chore:add seomthing", Error: errCategoryMissing}, FailingCommit{Hash: "testhash2", Message: "broken", Error: errCategoryMissing})

	formattedText := FormatFailingCommits(commits)
	assert.Equal(t, "\nFollowing commits failed the check: \n\nFAIL   testhash   chore:add seomthing   category missing\nFAIL   testhash2   broken   category missing\n\n", formattedText)
}
