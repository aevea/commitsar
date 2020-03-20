package text

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatFailingCommits(t *testing.T) {
	var testString bytes.Buffer

	var commits []FailingCommit

	commits = append(commits, FailingCommit{Hash: "testhash", Message: "chore:add seomthing", Error: errCategoryMissing}, FailingCommit{Hash: "testhash2", Message: "broken", Error: errCategoryMissing})

	table := FormatFailingCommits(commits)
	table.SetOutputMirror(&testString)
	table.Render()

	assert.Equal(t, "+-----------+------------------+---------------------+\n| HASH      | FAILURE          | TEXT                |\n+-----------+------------------+---------------------+\n| testhash  | category missing | chore:add seomthing |\n| testhash2 | category missing | broken              |\n+-----------+------------------+---------------------+\n", testString.String())
}
