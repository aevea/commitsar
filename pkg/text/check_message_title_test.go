package text

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckMessageTitle(t *testing.T) {
	tests := map[Commit]error{
		Commit{Category: "chore", Heading: "add something"}:                             nil,
		Commit{Category: "chore", Scope: "ci", Heading: "added new CI stuff"}:           nil,
		Commit{Category: "feat", Heading: "added a new feature"}:                        nil,
		Commit{Category: "fix", Breaking: true, Heading: "breaking change"}:             nil,
		Commit{Category: "fix", Scope: "security", Breaking: true, Heading: "breaking"}: nil,
		Commit{Category: "fix!", Breaking: true, Heading: "breaking"}:                   errCategoryWrongFormat,
		Commit{Category: "fix", Scope: "security(stuff)", Heading: "should break"}:      errScopeNonConform,
		Commit{}: errCategoryMissing,
		Commit{Category: "perf()", Heading: "nope"}: errCategoryWrongFormat,
		Commit{Category: "chore(", Heading: "bad"}:  errCategoryWrongFormat,
		Commit{Heading: "nope"}:                     errCategoryMissing,
		Commit{Category: "test", Scope: "full", Heading: "a heading", Body: "body is here\nit can have multiple lines"}: nil,
	}

	for test, expected := range tests {
		err := CheckMessageTitle(test)
		assert.Equal(t, expected, err)
	}
}
