package cmd

import (
	"github.com/spf13/cobra"
)

// accountImportCmd represents the accountImport command
var accountImportCmd = &cobra.Command{
	Use:                   "import <privatekey> <passphrase>",
	Short:                 "Import account from private key",
	Long:                  `Import account from private key`,
	Args:                  cobra.ExactArgs(2),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		client.NewAccountFromPrivateKey(args[0], args[1])
	},
}

func init() {
	accountCmd.AddCommand(accountImportCmd)
}
