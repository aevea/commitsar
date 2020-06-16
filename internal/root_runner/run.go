package root_runner

import (
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

	return nil
}
