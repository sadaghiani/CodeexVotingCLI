package cmd

import (
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

// voteShowCmd represents the voteShow command
var voteShowCmd = &cobra.Command{
	Use:                   "show <ID>",
	Short:                 "Show the names of a voting",
	Long:                  `Show the names of a voting`,
	Args:                  cobra.ExactArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		ID, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Fatalln(err)
		}
		res, err := client.GetListNameFromID(ID)
		if err != nil {
			log.Fatalln(err)
		}
		cmd.Println("Names : ", res)
	},
}

func init() {
	voteCmd.AddCommand(voteShowCmd)
}
