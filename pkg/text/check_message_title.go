package text

import (
	"errors"
	"regexp"
)

var (
	errCategoryMissing     = errors.New("category missing")
	errCategoryWrongFormat = errors.New("category wrong format")
	errScopeNonConform     = errors.New("malformed scope")
	errMissingBCBody       = errors.New("breaking change must contain commit body")
	errBCMissingText       = errors.New("breaking change commit body must start with BREAKING CHANGE: ")

	// Fields such as category and chore should contain only word characters
	fieldRegex = regexp.MustCompile(`^\w+$`)

	// Commits with breaking changes should contain text with BREAKING CHANGE: at start
	bcRegex = regexp.MustCompile(`^BREAKING CHANGE: `)
)

// CheckMessageTitle verifies that the message title conforms to
// conventional commmit standard https://www.conventionalcommits.org/en/v1.0.0-beta.4/#summary
func CheckMessageTitle(commit Commit) error {
	if commit.Category == "" {
		return errCategoryMissing
	}
	categoryMatch := fieldRegex.FindStringSubmatch(commit.Category)

	if categoryMatch == nil {
		return errCategoryWrongFormat
	}

	scopeMatch := fieldRegex.FindStringSubmatch(commit.Scope)
	if commit.Scope != "" && scopeMatch == nil {
		return errScopeNonConform
	}

	if commit.Breaking {
		if commit.Body == "" {
			return errMissingBCBody
		}
		bcMatch := bcRegex.FindStringSubmatch(commit.Body)
		if bcMatch == nil {
			return errBCMissingText
		}
	}

	return nil
}
