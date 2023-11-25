package prpipeline

import (
	"errors"

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

	switch pipeline.options.Style {
	case JiraStyle:
		references, err := jira.FindReferences(pipeline.options.Keys, *title)

		if err != nil {
			return nil, err
		}

		if len(references) > 0 {
			return &dispatcher.PipelineSuccess{
				PipelineName: pipeline.Name(),
				Message:      aurora.Sprintf(aurora.Green("Success! Found the following JIRA issue references: %v"), references),
			}, nil
		}
	case ConventionalStyle:
		commit := quoad.ParseCommitMessage(*title)

		err := text.CheckMessageTitle(commit, true)

		if err != nil {
			return nil, err
		}

		return &dispatcher.PipelineSuccess{
			PipelineName: pipeline.Name(),
			Message:      aurora.Sprintf(aurora.Green("Success! PR title is compliant with conventional commit")),
		}, nil
	}

	return nil, errors.New("pr checking is configured, but no style has been chosen")
}
