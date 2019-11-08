package text

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckMessageTitleNonStrict(t *testing.T) {
	tests := map[Commit]error{
		Commit{Category: "chore", Heading: "add something"}:                             nil,
		Commit{Category: "chore", Scope: "ci", Heading: "added new CI stuff"}:           nil,
		Commit{Category: "feat", Heading: "added a new feature"}:                        nil,
		Commit{Category: "fix", Breaking: true, Heading: "breaking change"}:             errMissingBCBody,
		Commit{Category: "fix", Scope: "security", Breaking: true, Heading: "breaking"}: errMissingBCBody,
		Commit{Category: "fix!", Breaking: true, Heading: "breaking"}:                   errCategoryWrongFormat,
		Commit{Category: "fix", Scope: "security(stuff)", Heading: "should break"}:      errScopeNonConform,
		Commit{}: errCategoryMissing,
		Commit{Category: "perf()", Heading: "nope"}: errCategoryWrongFormat,
		Commit{Category: "chore(", Heading: "bad"}:  errCategoryWrongFormat,
		Commit{Heading: "nope"}:                     errCategoryMissing,
		Commit{Category: "test", Scope: "full", Heading: "a heading", Body: "body is here\nit can have multiple lines"}: nil,
		Commit{Category: "test", Heading: "a heading", Body: "body is here", Breaking: true}:                            errBCMissingText,
		Commit{Category: "test", Heading: "a heading", Body: "BREAKING CHANGE: this happened", Breaking: true}:          nil,
	}

	for test, expected := range tests {
		err := CheckMessageTitle(test, false)
		assert.Equal(t, expected, err)
	}
}

func TestCheckMessageTitleStrict(t *testing.T) {
	tests := make(map[Commit]error)

	for _, cat := range allowedCategories {
		tests[Commit{Category: cat, Heading: "add something"}] = nil
	}

	tests[Commit{Category: "thisshouldneverbeacategory", Heading: "add something"}] = errNonStandardCategory

	for test, expected := range tests {
		err := CheckMessageTitle(test, true)
		assert.Equal(t, expected, err)
	}
}
