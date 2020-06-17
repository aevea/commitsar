package root_runner

import (
	"errors"

	"github.com/logrusorgru/aurora"
	"github.com/spf13/viper"
)

// Run executes the base command for Commitsar
func (runner *Runner) Run(options RunnerOptions, args ...string) error {
	if !viper.GetBool("commits.disabled") {
		err := runner.runCommits(options, args...)

		if err != nil {
			return err
		}
	} else {
		runner.Logger.Println("Commit section skipped due to commits.disabled set to true in .commitsar.yaml")
	}

	if viper.GetBool("pull_request.jira_title") {
		jiraKeys := viper.GetStringSlice("pull_request.jira_keys")
		references, err := runner.RunPullRequest(jiraKeys)

		if err != nil {
			return err
		}

		if len(references) == 0 {
			return errors.New("no references found in Pull Request title")
		}

		successMessage := aurora.Sprintf(aurora.Green("Success! Found the following JIRA issue references: %v \n"), references)

		runner.Logger.Print(successMessage)
	}

	return nil
}
