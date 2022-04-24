package cmd

import (
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

// voteDepositCmd represents the voteDeposit command
var voteDepositCmd = &cobra.Command{
	Use:                   "deposit <publicAddress> <passphrase> <ID> <name> <depositValue-wei>",
	Short:                 "Deposit for a vote",
	Long:                  `Deposit for a vote`,
	Args:                  cobra.ExactArgs(5),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		ID, err := strconv.ParseInt(args[2], 10, 64)
		if err != nil {
			log.Fatalln(err)
		}
		depositAmount, err := strconv.ParseInt(args[4], 10, 64)
		if err != nil {
			log.Fatalln(err)
		}
		res, err := client.Deposit(args[0], args[1], ID, args[3], depositAmount)
		if err != nil {
			log.Fatalln(err)
		}
		cmd.Println("successful\n", "txID : ", res)
	},
}

func init() {
	voteCmd.AddCommand(voteDepositCmd)
}
