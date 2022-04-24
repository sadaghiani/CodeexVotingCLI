package cmd

import (
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

// voteNewCmd represents the voteNew command
var voteNewCmd = &cobra.Command{
	Use:                   "new <publicAddress> <passphrase> <ID> <minimumDeposit> <endDepositTime> <endVotingTime>",
	Short:                 "Create a new voting",
	Long:                  `Create a new voting`,
	Args:                  cobra.ExactArgs(6),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {

		ID, err := strconv.ParseInt(args[2], 10, 64)
		if err != nil {
			log.Fatalln(err)
		}
		minimumDeposit, err := strconv.ParseInt(args[3], 10, 64)
		if err != nil {
			log.Fatalln(err)
		}
		endDepositTime, err := strconv.ParseInt(args[4], 10, 64)
		if err != nil {
			log.Fatalln(err)
		}
		endVotingTime, err := strconv.ParseInt(args[5], 10, 64)
		if err != nil {
			log.Fatalln(err)
		}
		res, err := client.NewVoting(args[0], args[1], ID, minimumDeposit, endDepositTime, endVotingTime)
		if err != nil {
			log.Fatalln(err)
		}
		cmd.Println("Add new voting\n", "txID : ", res)
	},
}

func init() {
	voteCmd.AddCommand(voteNewCmd)
}
