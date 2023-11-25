package text

import (
	"errors"
	"regexp"
	"strings"

	"github.com/aevea/quoad"
)

var (
	errCategoryMissing     = errors.New("category missing")
	errCategoryWrongFormat = errors.New("category wrong format")
	errNonStandardCategory = errors.New("category not one of " + strings.Join(allowedCategories[:], ","))
	errScopeNonConform     = errors.New("malformed scope")
	errMissingBCBody       = errors.New("breaking change must contain commit body")
	errBCMissingText       = errors.New("breaking change commit body must start with BREAKING CHANGE: ")

	// Fields such as category and chore should contain only word characters
	categoryRegex = regexp.MustCompile(`^\w+$`)
	scopeRegex    = regexp.MustCompile(`^\w+(-\w+|_\w+|/\w+| \w+)*$`)

	// Commits with breaking changes should contain text with BREAKING CHANGE: at start
	bcRegex = regexp.MustCompile(`^BREAKING CHANGE: `)

	// Types allowed by the commitlint's conventional preset
	allowedCategories = []string{"build", "ci", "chore", "docs", "feat", "fix", "perf", "refactor", "revert", "style", "test"}
)

func isAllowedCategory(category string) bool {
	for _, val := range allowedCategories {
		if val == category {
			return true
		}
	}

	return false
}

// CheckMessageTitle verifies that the message title conforms to
// conventional commit standard https://www.conventionalcommits.org/en/v1.0.0-beta.4/#summary
func CheckMessageTitle(commit quoad.Commit, strict bool) error {
	if commit.Category == "" {
		return errCategoryMissing
	}
	categoryMatch := categoryRegex.FindStringSubmatch(commit.Category)

	if categoryMatch == nil {
		return errCategoryWrongFormat
	}

	if strict && !isAllowedCategory(categoryMatch[0]) {
		return errNonStandardCategory
	}

	scopeMatch := scopeRegex.FindStringSubmatch(commit.Scope)
	if commit.Scope != "" && scopeMatch == nil {
		return errScopeNonConform
	}

	return nil
}
