package text

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsMergeCommit(t *testing.T) {
	tests := map[string]bool{
		"Merge commit '900a395d573f2b046d0b901be22808bf55319fc7'\n": true,
		"Merge branch 'master' into three\n":                        true,
		"Merge branch 'master' into feature/something-word":         true,
		"chore: something\n":                                        false,
		"fix: test":                                                 false,
	}

	for test, expected := range tests {
		err := IsMergeCommit(test)
		assert.Equal(t, expected, err)
	}
}
