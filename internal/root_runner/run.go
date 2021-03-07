package root_runner

import (
	"errors"

	"github.com/aevea/commitsar/internal/commitpipeline"
	"github.com/aevea/commitsar/internal/dispatcher"
	"github.com/apex/log"
	"github.com/logrusorgru/aurora"
	"github.com/spf13/viper"
)

// Run executes the base command for Commitsar
func (runner *Runner) Run(options RunnerOptions, args ...string) error {
	dispatch := dispatcher.New()

	var pipelines []dispatcher.Pipeliner

	if !viper.GetBool("commits.disabled") {
		commitOptions := commitpipeline.Options{
			AllCommits:     options.AllCommits,
			Limit:          options.Limit,
			Path:           options.Path,
			Strict:         options.Strict,
			UpstreamBranch: options.UpstreamBranch,
			RequiredScopes: options.RequiredScopes,
		}

		commitPipe, err := commitpipeline.New(&commitOptions, args...)

		if err != nil {
			return err
		}

		pipelines = append(pipelines, commitPipe)
	} else {
		log.Info("Commit section skipped due to commits.disabled set to true in .commitsar.yaml")
	}

	if viper.GetBool("pull_request.jira_title") {
		jiraKeys := viper.GetStringSlice("pull_request.jira_keys")
		references, err := runner.RunPullRequest(jiraKeys)

		if err != nil {
			return err
		}

		if len(references) == 0 {
			return errors.New("No JIRA references found in Pull Request title")
		}

		successMessage := aurora.Sprintf(aurora.Green("Success! Found the following JIRA issue references: %v"), references)

		log.Info(successMessage)
	}

	result := dispatch.RunPipelines(pipelines)

	if len(result.Errors) != 0 {
		return errors.New("Some errors were found")
	}

	for _, successMessage := range result.SuccessfulPipelines {
		log.Info(successMessage.Message)
	}

	return nil
}
