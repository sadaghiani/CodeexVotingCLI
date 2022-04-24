package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:               "CodeexVoting",
	Short:             "A command line application for account management and voting based on CodeexVotingContract",
	Long:              `A command line application for account management and voting based on CodeexVotingContract`,
	CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	Version:           "0.1.0",
	DisableAutoGenTag: true,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// err := doc.GenMarkdownTree(rootCmd, "./doc/")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
