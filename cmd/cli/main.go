package main

import (
	"fmt"
	"github.com/outillage/commitsar/internal/version_runner"
	"io/ioutil"
	"log"
	"os"

	"github.com/outillage/commitsar/internal/root_runner"
	"github.com/outillage/integrations"
	"github.com/spf13/cobra"
)

// Verbose is used to allow verbose/debug output for any given command
var Verbose bool

// Strict is used to enforce only standard categories
var Strict bool

// Dir is the location of repo to check
var Dir string

// AllCommits will iterate through all the commits on a branch
var AllCommits bool

// version is a global variable passed during build time
var version string

// commit is a global variable passed during build time. Should be used if version is not available.
var commit string

// date is a global variable passed during build time
var date string

func runRoot(cmd *cobra.Command, args []string) error {
	debug := false
	if cmd.Flag("verbose").Value.String() == "true" {
		debug = true
	}

	strict := true
	if cmd.Flag("strict").Value.String() == "false" {
		strict = false
	}

	upstreamBranch := integrations.FindCompareBranch()

	debugLogger := log.Logger{}
	debugLogger.SetPrefix("[DEBUG] ")
	debugLogger.SetOutput(os.Stdout)

	if !debug {
		debugLogger.SetOutput(ioutil.Discard)
		debugLogger.SetPrefix("")
	}

	logger := log.New(os.Stdout, "", 0)

	runner := root_runner.New(logger, &debugLogger, strict)

	options := root_runner.RunnerOptions{
		Path:           ".",
		UpstreamBranch: upstreamBranch,
		Limit:          0,
		AllCommits:     AllCommits,
	}

	return runner.Run(options, args...)
}

func main() {
	var rootCmd = &cobra.Command{
		Use:           "commitsar <from?>...<to>",
		Short:         "Checks if commits comply",
		Long:          "Checks if commits comply with conventional commits",
		RunE:          runRoot,
		SilenceUsage:  true,
		SilenceErrors: true,
		Args: cobra.MinimumNArgs(0),
	}

	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().BoolVarP(&Strict, "strict", "s", true, "strict mode")
	rootCmd.PersistentFlags().StringVarP(&Dir, "path", "d", ".", "dir points to the path of the repository")
	rootCmd.PersistentFlags().BoolVarP(&AllCommits, "all", "a", false, "iterate through all the commits on the branch")

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
			logger := log.New(os.Stdout, "", 0)

			err := version_runner.Run(
				version_runner.VersionInfo{
					Version: version,
					Date: date,
				},
				logger,
				)
			return err
		},
	}

	rootCmd.AddCommand(versionCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
