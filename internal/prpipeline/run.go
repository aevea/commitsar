package prpipeline

import (
	"errors"
	"fmt"
	"strings"

	"github.com/aevea/commitsar/internal/dispatcher"
	"github.com/aevea/commitsar/pkg/jira"
	"github.com/aevea/commitsar/pkg/text"
	"github.com/aevea/quoad"
	"github.com/logrusorgru/aurora"
)

func (pipeline *Pipeline) Run() (*dispatcher.PipelineSuccess, error) {
	title, err := getPRTitle(pipeline.options.Path)

	if err != nil {
		return nil, err
	}

	return pipeline.validateTitle(*title)
}

// validateTitle validates a PR title against the configured styles
func (pipeline *Pipeline) validateTitle(title string) (*dispatcher.PipelineSuccess, error) {
	if len(pipeline.options.Styles) == 0 {
		return nil, errors.New("pr checking is configured, but no style has been chosen")
	}

	var successMessages []string
	var allErrors []error

	for _, style := range pipeline.options.Styles {
		switch style {
		case JiraStyle:
			references, err := jira.FindReferences(pipeline.options.Keys, title)

			if err != nil {
				allErrors = append(allErrors, err)
				continue
			}

			if len(references) > 0 {
				successMessages = append(successMessages, aurora.Sprintf(aurora.Green("Found the following JIRA issue references: %v"), references))
			} else {
				allErrors = append(allErrors, errors.New("no JIRA issue references found in PR title"))
			}
		case ConventionalStyle:
			commit := quoad.ParseCommitMessage(title)

			err := text.CheckMessageTitle(commit, true)

			if err != nil {
				allErrors = append(allErrors, err)
				continue
			}

			successMessages = append(successMessages, aurora.Sprintf(aurora.Green("PR title is compliant with conventional commit")))
		}
	}

	if len(allErrors) > 0 {
		// Combine all errors into a single error message
		errorMessages := make([]string, len(allErrors))
		for i, err := range allErrors {
			errorMessages[i] = err.Error()
		}
		return nil, fmt.Errorf("validation failed: %s", strings.Join(errorMessages, "; "))
	}

	// Combine all success messages
	combinedMessage := "Success! " + strings.Join(successMessages, " ")

	return &dispatcher.PipelineSuccess{
		PipelineName: pipeline.Name(),
		Message:      combinedMessage,
	}, nil
}
