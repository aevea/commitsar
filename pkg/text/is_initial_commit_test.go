package text

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsInitialCommit(t *testing.T) {
	tests := map[string]bool {
		"Initial commit": true,
	}

	for test, expected := range tests {
		assert.Equal(t, expected, IsInitialCommit(test))
	}
}
