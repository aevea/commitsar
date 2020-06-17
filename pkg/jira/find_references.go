package jira

import (
	"regexp"
)

// FindReferences scans a given message looking for all issues that match the keys. Default to the JIRA default regex if no keys are provided.
func FindReferences(keys []string, message string) ([]string, error) {
	regex, err := regexp.Compile(buildRegex(keys))

	if err != nil {
		return nil, err
	}

	matches := regex.FindAllString(message, -1)

	return matches, nil
}
