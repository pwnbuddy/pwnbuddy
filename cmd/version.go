package cmd

import (
	"PwnBuddy/utils"
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints version of PwnBuddy",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(utils.GetAsciiString())
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
