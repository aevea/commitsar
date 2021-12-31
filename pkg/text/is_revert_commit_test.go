package text

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsRevertCommit(t *testing.T) {
	tests := map[string]bool{
		"Revert some commit":      true,
		"revert \"some commit\"":  true,
		"chore: something\n":      false,
		"fix: test":               false,
		"fix: Kodiak style regex": false,
	}

	for test, expected := range tests {
		err := IsRevertCommit(test)
		assert.Equal(t, expected, err)
	}
}
