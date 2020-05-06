package main

import (
	"PwnBuddy/cmd"
	"PwnBuddy/utils"
	"fmt"

	"github.com/c-bata/go-prompt"
)

func main() {
	shell := &cmd.CMD{
		RootCommand:        cmd.RootCmd,
		SuggestionFunction: nil,
		PromptOptions: []prompt.Option{
			prompt.OptionTitle("PwnBuddy"),
			prompt.OptionPrefix("pwnbuddy> "),
			prompt.OptionShowCompletionAtStart(),
			prompt.OptionMaxSuggestion(10),
		},
	}

	fmt.Print(utils.GetAsciiString())

	shell.Run()
}
