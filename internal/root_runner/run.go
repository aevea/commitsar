package root_runner

import (
	"errors"

	"github.com/aevea/commitsar/internal/commitpipeline"
	"github.com/aevea/commitsar/internal/dispatcher"
	"github.com/aevea/commitsar/internal/prpipeline"
	"github.com/apex/log"
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

	if viper.IsSet("pull_request") {
		log.Debug("PR pipeline")

		prOptions := prpipeline.Options{
			Path: options.Path,
		}

		if viper.IsSet("pull_request.jira_title") {
			prOptions.Style = prpipeline.JiraStyle
			prOptions.Keys = viper.GetStringSlice("pull_request.jira_keys")
		}

		if viper.IsSet("pull_request.conventional") {
			prOptions.Style = prpipeline.ConventionalStyle
		}

		prPipe, err := prpipeline.New(prOptions)

		if err != nil {
			return err
		}

		pipelines = append(pipelines, prPipe)
	}

	if len(pipelines) == 0 {
		return errors.New("no pipelines defined")
	}

	result := dispatch.RunPipelines(pipelines)

	if len(result.Errors) != 0 {
		for _, err := range result.Errors {
			log.Errorf("pipeline: %s, error: %s", err.PipelineName, err.Error)
		}
		return errors.New("some errors were found")
	}

	for _, successMessage := range result.SuccessfulPipelines {
		log.Info(successMessage.Message)
	}

	return nil
}
