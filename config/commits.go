package config

import (
	"github.com/aevea/commitsar/internal/root_runner"
	"github.com/spf13/viper"
)

// CommitConfig will return the RunnerOptions using defaults unless overriden in config or flags
func CommitConfig() root_runner.RunnerOptions {
	// defaults
	strict := true
	limit := 0
	all := false
	requiredScopes := []string{}

	if viper.IsSet("commits.strict") {
		strict = viper.GetBool("commits.strict")
	}

	if viper.IsSet("commits.limit") {
		limit = viper.GetInt("commits.limit")
	}

	if viper.IsSet("commits.all") {
		all = viper.GetBool("commits.all")
	}

	if viper.IsSet("commits.required-scopes") {
		requiredScopes = viper.GetStringSlice("commits.required-scopes")
	}

	return root_runner.RunnerOptions{
		Strict:         strict,
		Limit:          limit,
		AllCommits:     all,
		RequiredScopes: requiredScopes,
	}
}
