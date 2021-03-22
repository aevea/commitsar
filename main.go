package main

import (
	"fmt"
	"os"

	"github.com/aevea/commitsar/config"
	"github.com/aevea/commitsar/internal/version_runner"
	"github.com/logrusorgru/aurora"

	"github.com/aevea/commitsar/internal/root_runner"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
)

// version is a global variable passed during build time
var version string

// commit is a global variable passed during build time. Should be used if version is not available.
var commit string

// date is a global variable passed during build time
var date string

func runRoot(cmd *cobra.Command, args []string) error {
	if viper.GetBool("verbose") {
		log.SetLevel(log.DebugLevel)
	}

	runner := root_runner.New()

	commitConfig := config.CommitConfig()

	commitConfig.Path = "."

	return runner.Run(commitConfig, args...)
}

func bindRootFlags(rootCmd *cobra.Command) error {
	rootCmd.Flags().BoolP("verbose", "v", false, "verbose output")
	err := viper.BindPFlag("verbose", rootCmd.Flags().Lookup("verbose"))
	if err != nil {
		return err
	}
	rootCmd.Flags().BoolP("strict", "s", true, "strict mode")
	err = viper.BindPFlag("commits.strict", rootCmd.Flags().Lookup("strict"))
	if err != nil {
		return err
	}
	rootCmd.Flags().BoolP("all", "a", false, "iterate through all the commits on the branch")
	err = viper.BindPFlag("commits.all", rootCmd.Flags().Lookup("all"))
	if err != nil {
		return err
	}
	rootCmd.Flags().StringSlice("required-scopes", nil, "forces scope to match one of the provided values")
	err = viper.BindPFlag("commits.required-scopes", rootCmd.Flags().Lookup("required-scopes"))
	if err != nil {
		return err
	}

	// Not used. TODO: Documentation
	rootCmd.Flags().StringP("path", "d", ".", "dir points to the path of the repository")
	err = viper.BindPFlag("path", rootCmd.Flags().Lookup("path"))
	if err != nil {
		return err
	}
	return nil
}

func main() {
	log.SetHandler(cli.Default)

	if err := config.LoadConfig(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var rootCmd = &cobra.Command{
		Use:           "commitsar <from?>...<to>",
		Short:         "Checks if commits comply",
		Long:          "Checks if commits comply with conventional commits",
		RunE:          runRoot,
		SilenceUsage:  true,
		SilenceErrors: true,
		Args:          cobra.MinimumNArgs(0),
	}

	if err := bindRootFlags(rootCmd); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Version returns undefined if not on a tag. This needs to reset it.
	if version == "undefined" {
		version = ""
	}

	if version == "" && commit != "" {
		version = commit
	}
	if version == "" && commit == "" {
		version = "development"
	}

	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of Commitsar",
		Long:  `All software has versions. This is Commitsars.`,
		RunE: func(cmd *cobra.Command, args []string) error {

			err := version_runner.Run(
				version_runner.VersionInfo{
					Version: version,
					Date:    date,
				},
			)
			return err
		},
	}

	rootCmd.AddCommand(versionCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(aurora.Red(err))
		os.Exit(1)
	}
}
