package main

import (
	"fmt"
	"github.com/outillage/commitsar/internal/version_runner"
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

// Version is a global variable passed during build time
var Version string

// Commit is a global variable passed during build time. Should be used if version is not available.
var Commit string

// BuildTime is a global variable passed during build time
var BuildTime string

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

	return root_runner.RunCommitsar(".", upstreamBranch, debug, strict, args...)
}

func main() {
	var rootCmd = &cobra.Command{
		Use:           "commitsar <from?>...<to>",
		Short:         "Checks if commits comply",
		Long:          "Checks if commits comply with conventional commits",
		RunE:          runRoot,
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().BoolVarP(&Strict, "strict", "s", true, "strict mode")
	rootCmd.PersistentFlags().StringVarP(&Dir, "path", "d", ".", "dir points to the path of the repository")

	version := Version

	// Version returns undefined if not on a tag. This needs to reset it.
	if Version == "undefined" {
		version = ""
	}

	if Version == "" && Commit != "" {
		version = Commit
	}
	if Version == "" && Commit == "" {
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
					Date: BuildTime,
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
