package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:           "commitsar",
	Short:         "Checks if commits comply",
	Long:          "Checks if commits comply with conventional commits",
	RunE:          (runRoot),
	SilenceUsage:  true,
	SilenceErrors: true,
}

// Verbose is used to allow verbose/debug output for any given command
var Verbose bool

// Strict is used to enforce only standard categories
var Strict bool

// Dir is the location of repo to check
var Dir string

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().BoolVarP(&Strict, "strict", "s", true, "strict mode")
	rootCmd.PersistentFlags().StringVarP(&Dir, "path", "d", ".", "dir points to the path of the repository")
}

// Execute just executes the rootCmd for Cobra
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
