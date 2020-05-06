package cmd

import (
	"PwnBuddy/consts"
	"fmt"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "pwnbuddy>",
	Short: "PwnBuddy Framework",
	Long: `
	PwnBuddy is blah blah blah
	`,
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		//os.Exit(1) //Commented out, breaks the prompt loop
	}
}

func init() {
	RootCmd.InitDefaultHelpCmd()
	RootCmd.InitDefaultHelpFlag()

	RootCmd.SetUsageTemplate(consts.USAGE_TEMPLATE)
	RootCmd.SetHelpTemplate(consts.HELP_TEMPLATE)
	fmt.Println(exitCmd.UseLine())
	fmt.Println(RootCmd.UseLine())

	// Add Persistent flags for all subcommands.
	// Or you can add local flags for only root command
}
