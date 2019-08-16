package text

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMessageTitle(t *testing.T) {
	tests := map[string]string{
		"chore: add something\n":             "chore: add something",
		"chore: add something\n description": "chore: add something",
	}

	for test, expected := range tests {
		err := MessageTitle(test)
		assert.Equal(t, expected, err)
	}
}
