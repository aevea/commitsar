package jira

import "strings"

func buildRegex(keys []string) string {
	// This is the default regex for JIRA tickets. Please see https://community.atlassian.com/t5/Bitbucket-questions/Regex-pattern-to-match-JIRA-issue-key/qaq-p/233319 for reference
	if len(keys) == 0 {
		return `([A-Z]+-[\d]+)`
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
