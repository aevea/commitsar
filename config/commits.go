package config

import (
	"github.com/aevea/commitsar/internal/root_runner"
	"github.com/aevea/integrations"
	"github.com/spf13/viper"
	"fmt"
	"os"
	"github.com/apex/log"
	"log"
)

const (
	// CommitsarConfigPath is used an env variable to override the default location of the config file.
	CommitsarConfigPath = "COMMITSAR_CONFIG_PATH"
)

// CommitConfig will return the RunnerOptions using defaults unless overridden in config or flags
func CommitConfig() root_runner.RunnerOptions {
	// defaults
	strict := true
	limit := 0
	all := false
	upstreamBranch := integrations.FindCompareBranch()
	requiredScopes := []string{}

	if err := LoadConfig(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if viper.GetBool("verbose") {
		log.SetLevel(log.DebugLevel)
	}

	if viper.IsSet("commits.strict") {
		strict = viper.GetBool("commits.strict")
	}

	if viper.IsSet("commits.limit") {
		limit = viper.GetInt("commits.limit")
	}

	if viper.IsSet("commits.all") {
		all = viper.GetBool("commits.all")
	}

	if viper.IsSet("commits.upstreamBranch") {
		upstreamBranch = viper.GetString("commits.upstreamBranch")
	}

	if viper.IsSet("commits.required-scopes") {
		requiredScopes = viper.GetStringSlice("commits.required-scopes")
	}

	return root_runner.RunnerOptions{
		Strict:         strict,
		Limit:          limit,
		AllCommits:     all,
		UpstreamBranch: upstreamBranch,
		RequiredScopes: requiredScopes,
	}
}

// LoadConfig iterates through possible config paths. No config will be loaded if no files are present.
func LoadConfig() error {
	viper.AutomaticEnv()
	viper.SetConfigName(".commitsar")
	viper.SetConfigType("yaml")

	if viper.IsSet(CommitsarConfigPath) {
		viper.AddConfigPath(viper.GetString(CommitsarConfigPath))
	}

	viper.AddConfigPath(viper.GetString("config-path"))

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			log.Warn("config file not found, using defaults")
		} else {
			// Config file was found but another error was produced
			return err
		}
	}
	return nil
}
