package history

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpenGit(t *testing.T) {
	git, err := OpenGit("../../", true)

	// Should not error if this git repository is valid
	assert.NoError(t, err)
	// Check that debug flag is passed correctly
	assert.Equal(t, true, git.Debug)

	_, unhappyErr := OpenGit(".", false)

	// Should error opening a folder with missing .git
	assert.Error(t, unhappyErr)

}
