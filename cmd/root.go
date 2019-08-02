package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "commitsar",
	Short: "Checks if commits comply",
	Long:  "Checks if commits comply with conventional commits",
	Run: (func(cmd *cobra.Command, args []string) {
		log.Println("hi")
	}),
}

// Verbose is used to allow verbose/debug output for any given command
var Verbose bool

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
}

// Execute just executes the rootCmd for Cobra
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
