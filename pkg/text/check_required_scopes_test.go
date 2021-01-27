package text

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequiredScopesChecker(t *testing.T) {
	tests := map[error][]struct {
		scope          string
		requiredScopes []string
	}{
		nil: {
			{"ci", []string{}},
			{"ui", nil},
			{"ci", []string{"ci", "project-1", "repo", "ui"}},
			{"ci", []string{"ui", "project-1", "repo", "ci"}},
			{"repo", []string{"repo"}},
		},
		errMissingRequiredScope: {
			{"SECURITY", []string{"security", "ci", "ui"}},
			{"SeCuRiTy", []string{"security", "ci", "ui"}},
			{"sth", []string{"security", "ci", "ui"}},
			{"se curity", []string{"security", "ci", "ui"}},
			{"", []string{"repo"}},
		},
	}

	for expected, testCases := range tests {
		for _, testCase := range testCases {
			err := RequiredScopeChecker(testCase.requiredScopes)(testCase.scope)
			assert.Equal(t, expected, err)
		}
	}
}
