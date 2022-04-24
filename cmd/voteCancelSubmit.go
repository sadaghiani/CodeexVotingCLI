package cmd

import (
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

// voteCancelSubmitCmd represents the voteCancelSubmit command
var voteCancelSubmitCmd = &cobra.Command{
	Use:                   "cancelsubmit <publicAddress> <passphrase> <ID> <name>",
	Short:                 "Cancel a vote",
	Long:                  `Cancel a vote`,
	Args:                  cobra.ExactArgs(4),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		ID, err := strconv.ParseInt(args[2], 10, 64)
		if err != nil {
			log.Fatalln(err)
		}
		res, err := client.ClearVote(args[0], args[1], ID, args[3])
		if err != nil {
			log.Fatalln(err)
		}
		cmd.Println("successful\n", "txID : ", res)
	},
}

func init() {
	voteCmd.AddCommand(voteCancelSubmitCmd)
}
