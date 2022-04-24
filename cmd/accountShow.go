package cmd

import (
	"github.com/spf13/cobra"
)

var defualt bool

// accountShowCmd represents the accountShow command
var accountShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show all accounts",
	Long:  `Show all accounts`,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		accounts, _ := client.AccountsList()
		for k, v := range accounts {
			cmd.Println(k, v)
		}
	},
}

func init() {
	accountCmd.AddCommand(accountShowCmd)
}
