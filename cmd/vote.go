package cmd

import (
	"github.com/spf13/cobra"
)

// voteCmd represents the vote command
var voteCmd = &cobra.Command{
	Use:   "vote",
	Short: "Voting management",
	Long:  `Voting management`,
}

func init() {
	rootCmd.AddCommand(voteCmd)
}
