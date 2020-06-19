package jira

import "strings"

const (
	defaultJiraRegex = `([A-Z][A-Z0-9]{1,10}-[0-9]+)`
)

func buildRegex(keys []string) string {
	if len(keys) == 0 {
		return defaultJiraRegex
	}

	result := strings.Builder{}

	result.WriteString("(")

	for index, key := range keys {
		result.WriteString(key)

		if (index + 1) != len(keys) {
			result.WriteString("|")
		}
	}

	result.WriteString(")")

	result.WriteString("-[0-9]+")

	return result.String()
}
