package jira

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindReferences(t *testing.T) {
	tests := []struct {
		expected []string
		keys     []string
		message  string
	}{
		{
			expected: []string{"TES-1"},
			keys:     nil,
			message:  "TES-1: added a tes feature",
		},
		{
			expected: []string{"TES-1"},
			keys:     []string{"TES"},
			message:  "TES-1: added a tes feature",
		},
		{
			expected: nil,
			keys:     []string{"TES"},
			message:  "REST-1: added a tes feature",
		},
		{
			expected: []string{"REST-1", "TEST-2"},
			keys:     []string{"TEST", "REST"},
			message:  "REST-1: added a tes feature TEST-2",
		},
		{
			expected: []string{"TEST-2"},
			keys:     []string{"TEST"},
			message:  "REST-1: added a tes feature TEST-2",
		},
		{
			expected: []string{"REST-1", "TEST-2"},
			keys:     nil,
			message:  "REST-1: added a tes feature TEST-2",
		},
		{
			expected: []string{"QA-336"},
			keys:     nil,
			message:  "QA-336 test: add workaround for test until issue fixed",
		},
	}

	for _, test := range tests {
		references, err := FindReferences(test.keys, test.message)

		assert.NoError(t, err)
		assert.Equal(t, test.expected, references)
	}
}
