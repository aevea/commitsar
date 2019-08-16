package text

import (
	"regexp"
	"strings"
)

var titleRegex = regexp.MustCompile("^.*\n")

// MessageTitle returns only the first line of commit message
func MessageTitle(message string) string {
	match := titleRegex.FindString(message)

	match = strings.Replace(match, "\n", "", 1)

	return match
}
