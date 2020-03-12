package text

import (
	"testing"

	"github.com/outillage/quoad"
	"github.com/stretchr/testify/assert"
)

func TestCheckMessageTitleNonStrict(t *testing.T) {
	tests := map[error][]quoad.Commit{
		nil: []quoad.Commit{
			quoad.Commit{Category: "chore", Heading: "add something"},
			quoad.Commit{Category: "chore", Scope: "ci", Heading: "added new CI stuff"},
			quoad.Commit{Category: "feat", Heading: "added a new feature"},
			quoad.Commit{Category: "test", Scope: "full", Heading: "a heading", Body: "body is here\nit can have multiple lines"},
			quoad.Commit{Category: "test", Heading: "a heading", Body: "BREAKING CHANGE: this happened", Breaking: true},
			quoad.Commit{Category: "chore", Scope: "new integration", Heading: "added new integration"},
		},
		errBCMissingText: []quoad.Commit{
			quoad.Commit{Category: "test", Heading: "a heading", Body: "body is here", Breaking: true},
		},
		errMissingBCBody: []quoad.Commit{
			quoad.Commit{Category: "fix", Breaking: true, Heading: "breaking change"},
			quoad.Commit{Category: "fix", Scope: "security", Breaking: true, Heading: "breaking"},
		},
		errCategoryWrongFormat: []quoad.Commit{
			quoad.Commit{Category: "fix!", Breaking: true, Heading: "breaking"},
		},
		errScopeNonConform: []quoad.Commit{
			quoad.Commit{Category: "fix", Scope: "security(stuff)", Heading: "should break"},
		},
		errCategoryMissing: []quoad.Commit{
			quoad.Commit{},
			quoad.Commit{Heading: "nope"},
		},
		errCategoryWrongFormat: []quoad.Commit{
			quoad.Commit{Category: "chore(", Heading: "bad"},
			quoad.Commit{Category: "perf()", Heading: "nope"},
		},
	}

	for expected, testCases := range tests {
		for _, testCase := range testCases {
			err := CheckMessageTitle(testCase, false)
			assert.Equal(t, expected, err)
		}
	}
}

func TestCheckMessageTitleStrict(t *testing.T) {
	tests := make(map[error]quoad.Commit)

	for _, cat := range allowedCategories {
		tests[nil] = quoad.Commit{Category: cat, Heading: "add something"}
	}

	tests[errNonStandardCategory] =  quoad.Commit{Category: "thisshouldneverbeacategory", Heading: "add something"}

	for expected, test := range tests {
		err := CheckMessageTitle(test, true)
		assert.Equal(t, expected, err)
	}
}
