package cmd

import (
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

// voteWithdrawCmd represents the voteWithdraw command
var voteWithdrawCmd = &cobra.Command{
	Use:                   "withdraw <publicAddress> <passphrase> <ID>",
	Short:                 "Withdrawal of deposit commensurate with the vote",
	Long:                  `Withdrawal of deposit commensurate with the vote`,
	Args:                  cobra.ExactArgs(3),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		ID, err := strconv.ParseInt(args[2], 10, 64)
		if err != nil {
			log.Fatalln(err)
		}
		res, err := client.Withdraw(args[0], args[1], ID)
		if err != nil {
			log.Fatalln(err)
		}
		cmd.Println("successful\n", "txID : ", res)
	},
}

func init() {
	voteCmd.AddCommand(voteWithdrawCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// voteWithdrawCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// voteWithdrawCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
