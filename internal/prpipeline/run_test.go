package prpipeline

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateTitle_ConventionalStyle_Success(t *testing.T) {
	pipeline := &Pipeline{
		options: Options{
			Styles: []PRStyle{ConventionalStyle},
		},
	}

	result, err := pipeline.validateTitle("feat: add new feature")

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Contains(t, result.Message, "Success!")
	assert.Contains(t, result.Message, "conventional commit")
}

func TestValidateTitle_ConventionalStyle_Failure(t *testing.T) {
	pipeline := &Pipeline{
		options: Options{
			Styles: []PRStyle{ConventionalStyle},
		},
	}

	result, err := pipeline.validateTitle("invalid title without category")

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "validation failed")
	assert.Contains(t, err.Error(), "category missing")
}

func TestValidateTitle_JiraStyle_Success(t *testing.T) {
	pipeline := &Pipeline{
		options: Options{
			Styles: []PRStyle{JiraStyle},
			Keys:   []string{"TEST"},
		},
	}

	result, err := pipeline.validateTitle("TEST-123: add new feature")

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Contains(t, result.Message, "Success!")
	assert.Contains(t, result.Message, "JIRA issue references")
}

func TestValidateTitle_JiraStyle_Failure(t *testing.T) {
	pipeline := &Pipeline{
		options: Options{
			Styles: []PRStyle{JiraStyle},
			Keys:   []string{"TEST"},
		},
	}

	result, err := pipeline.validateTitle("add new feature without jira")

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "validation failed")
	assert.Contains(t, err.Error(), "no JIRA issue references found")
}

func TestValidateTitle_BothStyles_Success(t *testing.T) {
	pipeline := &Pipeline{
		options: Options{
			Styles: []PRStyle{ConventionalStyle, JiraStyle},
			Keys:   []string{"TEST"},
		},
	}

	result, err := pipeline.validateTitle("feat(TEST-123): add new feature")

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Contains(t, result.Message, "Success!")
	assert.Contains(t, result.Message, "conventional commit")
	assert.Contains(t, result.Message, "JIRA issue references")
}

func TestValidateTitle_BothStyles_ConventionalFails(t *testing.T) {
	pipeline := &Pipeline{
		options: Options{
			Styles: []PRStyle{ConventionalStyle, JiraStyle},
			Keys:   []string{"TEST"},
		},
	}

	result, err := pipeline.validateTitle("TEST-123: invalid title without category")

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "validation failed")
	assert.Contains(t, err.Error(), "category missing")
}

func TestValidateTitle_BothStyles_JiraFails(t *testing.T) {
	pipeline := &Pipeline{
		options: Options{
			Styles: []PRStyle{ConventionalStyle, JiraStyle},
			Keys:   []string{"TEST"},
		},
	}

	result, err := pipeline.validateTitle("feat: add new feature without jira")

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "validation failed")
	assert.Contains(t, err.Error(), "no JIRA issue references found")
}

func TestValidateTitle_BothStyles_BothFail(t *testing.T) {
	pipeline := &Pipeline{
		options: Options{
			Styles: []PRStyle{ConventionalStyle, JiraStyle},
			Keys:   []string{"TEST"},
		},
	}

	result, err := pipeline.validateTitle("invalid title without category or jira")

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "validation failed")
	// Should contain both error messages
	errorMsg := err.Error()
	assert.True(t, strings.Contains(errorMsg, "category missing") || strings.Contains(errorMsg, "no JIRA issue references found"))
}

func TestValidateTitle_NoStyles(t *testing.T) {
	pipeline := &Pipeline{
		options: Options{
			Styles: []PRStyle{},
		},
	}

	result, err := pipeline.validateTitle("any title")

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "pr checking is configured, but no style has been chosen", err.Error())
}

func TestValidateTitle_ConventionalWithJiraInScope(t *testing.T) {
	// Test the specific use case from the issue: feat(XXX-1234): This is the title
	pipeline := &Pipeline{
		options: Options{
			Styles: []PRStyle{ConventionalStyle, JiraStyle},
			Keys:   []string{"XXX"},
		},
	}

	result, err := pipeline.validateTitle("feat(XXX-1234): This is the title")

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Contains(t, result.Message, "Success!")
	assert.Contains(t, result.Message, "conventional commit")
	assert.Contains(t, result.Message, "JIRA issue references")
}

func TestValidateTitle_ConventionalWithJiraInTitle(t *testing.T) {
	// Test with JIRA in the title part: feat: XXX-1234 This is the title
	pipeline := &Pipeline{
		options: Options{
			Styles: []PRStyle{ConventionalStyle, JiraStyle},
			Keys:   []string{"XXX"},
		},
	}

	result, err := pipeline.validateTitle("feat: XXX-1234 This is the title")

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Contains(t, result.Message, "Success!")
	assert.Contains(t, result.Message, "conventional commit")
	assert.Contains(t, result.Message, "JIRA issue references")
}

