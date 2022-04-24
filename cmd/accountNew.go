package cmd

import (
	"github.com/spf13/cobra"
)

// accountNewCmd represents the accountNew command
var accountNewCmd = &cobra.Command{
	Use:                   "new <passphrase>",
	Short:                 "Create new account",
	Long:                  `Create new account`,
	Args:                  cobra.ExactArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		client.CreateAccount(args[0])
	},
}

func init() {
	accountCmd.AddCommand(accountNewCmd)
}
