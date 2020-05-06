package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var exitCmd = &cobra.Command{
	Use:   "exit",
	Short: "Exit from PwnBuddy",
	Run: func(cmd *cobra.Command, args []string) {
		os.Exit(0)
	},
}

func init() {
	RootCmd.AddCommand(exitCmd)
}
