package history

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepo(t *testing.T) {
	_, err := Repo("../../")

	// Should not error if this git repository is valid
	assert.Equal(t, err, nil)

	_, unhappyErr := Repo(".")

	// Should error opening a folder with missing .git
	assert.NotEqual(t, unhappyErr, nil)

}
