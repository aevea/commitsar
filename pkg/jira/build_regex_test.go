package jira

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildRegex(t *testing.T) {
	tests := map[string][]string{
		`([A-Z]+-[\d]+)`:          {},
		`(TEST|PROJ|FEOP)-[0-9]+`: {"TEST", "PROJ", "FEOP"},
	}

	for expected, test := range tests {
		assert.Equal(t, expected, buildRegex(test))
	}
}
