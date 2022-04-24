package cmd

import (
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

// voteSubmitCmd represents the voteSubmit command
var voteSubmitCmd = &cobra.Command{
	Use:                   "submit <publicAddress> <passphrase> <ID> <name>",
	Short:                 "Submit a vote in a voting",
	Long:                  `Submit a vote in a voting`,
	Args:                  cobra.ExactArgs(4),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		ID, err := strconv.ParseInt(args[2], 10, 64)
		if err != nil {
			log.Fatalln(err)
		}
		res, err := client.SubmitVote(args[0], args[1], ID, args[3])
		if err != nil {
			log.Fatalln(err)
		}
		cmd.Println("successful\n", "txID : ", res)
	},
}

func init() {
	voteCmd.AddCommand(voteSubmitCmd)
}
