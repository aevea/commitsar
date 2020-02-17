package main

import (
	"fmt"
	"os"

	"github.com/outillage/commitsar/internal/runners"
	"github.com/outillage/integrations"
	"github.com/spf13/cobra"
)

// Verbose is used to allow verbose/debug output for any given command
var Verbose bool

// Strict is used to enforce only standard categories
var Strict bool

// Dir is the location of repo to check
var Dir string

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

	return runners.RunCommitsar(".", upstreamBranch, debug, strict, args...)
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

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
