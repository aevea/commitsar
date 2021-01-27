package text

import (
	"errors"
)

var errMissingRequiredScope = errors.New("required scope missing")

// RequiredScopeChecker creates case sensitive strict checker, validating that
// provided scope matches one from required scopes list
func RequiredScopeChecker(requiredScopes []string) func(string) error {
	return func(scope string) error {
		if len(requiredScopes) < 1 {
			return nil
		}

		for _, rs := range requiredScopes {
			if scope == rs {
				return nil
			}
		}

		return errMissingRequiredScope
	}
}
