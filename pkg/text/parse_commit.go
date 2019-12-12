package text

import (
	"github.com/outillage/quoad"
	"regexp"
	"strings"
)

var (
	expectedFormatRegex = regexp.MustCompile(`(?s)^(?P<category>\S+?)?(?P<scope>\(\S+\))?(?P<breaking>!?)?: (?P<heading>[^\n\r]+)?([\n\r]{2}(?P<body>.*))?`)
)

// ParseCommit takes a commits message and parses it into usable blocks
func ParseCommit(message string, hash [20]byte) quoad.Commit {
	match := expectedFormatRegex.FindStringSubmatch(message)

	if len(match) == 0 {
		return quoad.Commit{}
	}

	result := make(map[string]string)
	for i, name := range expectedFormatRegex.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}

	scope := result["scope"]

	// strip brackets from scope if present
	if scope != "" {
		scope = strings.Replace(scope, "(", "", 1)
		scope = strings.Replace(scope, ")", "", 1)
	}

	return quoad.Commit{
		Category: result["category"],
		Scope:    scope,
		Breaking: result["breaking"] == "!",
		Heading:  result["heading"],
		Body:     result["body"],
		Hash:     hash,
	}
}
