package cmd

import (
	"github.com/spf13/cobra"
)

// accountCmd represents the account command
var accountCmd = &cobra.Command{
	Use:   "account",
	Short: "Account management",
	Long:  `Account management`,
}

func init() {
	rootCmd.AddCommand(accountCmd)
}
