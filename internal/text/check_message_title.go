package text

import (
	"errors"
	"regexp"
)

var (
	titleRegex         = regexp.MustCompile(`^(?P<type>\w+?)(?P<scope>\(\w+\)?)*!*:\s(?P<message>.+)$`)
	errTitleNonConform = errors.New("message title does not conform to conventional commits")
)

// CheckMessageTitle verifies that the message title conforms to
// conventional commmit standard https://www.conventionalcommits.org/en/v1.0.0-beta.4/#summary
func CheckMessageTitle(message string) error {
	match := titleRegex.FindStringSubmatch(message)
	if match == nil {
		return errTitleNonConform
	}
	return nil
}
