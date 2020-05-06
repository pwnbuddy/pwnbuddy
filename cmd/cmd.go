package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const CALLBACK_ANNOTATION = "pwnbuddy"

type CMD struct {
	RootCommand        *cobra.Command
	PromptOptions      []prompt.Option
	SuggestionFunction func(annotation string, document prompt.Document) []prompt.Suggest
}

func (c *CMD) Run() {
	p := prompt.New(
		func(in string) {
			promptArgs := strings.Fields(in)
			os.Args = append([]string{os.Args[0]}, promptArgs...)
			err := c.RootCommand.Execute()
			if err != nil {
				fmt.Println(err)
			}
		},
		func(d prompt.Document) []prompt.Suggest {
			return findSuggestions(c, d)
		},
		c.PromptOptions...,
	)
	p.Run()
}

func findSuggestions(c *CMD, d prompt.Document) []prompt.Suggest {
	command := c.RootCommand
	args := strings.Fields(d.CurrentLine())

	if found, _, err := command.Find(args); err == nil {
		command = found
	}

	var suggestions []prompt.Suggest
	resetFlags, _ := command.Flags().GetBool("flags-no-reset")
	addFlags := func(flag *pflag.Flag) {
		if flag.Changed && !resetFlags {
			flag.Value.Set(flag.DefValue)
		}
		if flag.Hidden {
			return
		}
		if strings.HasPrefix(d.GetWordBeforeCursor(), "--") {
			suggestions = append(suggestions, prompt.Suggest{Text: "--" + flag.Name, Description: flag.Usage})
		} else if strings.HasPrefix(d.GetWordBeforeCursor(), "-") && flag.Shorthand != "" {
			suggestions = append(suggestions, prompt.Suggest{Text: "-" + flag.Shorthand, Description: flag.Usage})
		}
	}

	command.LocalFlags().VisitAll(addFlags)
	command.InheritedFlags().VisitAll(addFlags)

	if command.HasAvailableSubCommands() {
		for _, c := range command.Commands() {
			if !c.Hidden {
				suggestions = append(suggestions, prompt.Suggest{Text: c.Name(), Description: c.Short})
			}
		}
	}

	annotation := command.Annotations[CALLBACK_ANNOTATION]
	if c.SuggestionFunction != nil && annotation != "" {
		suggestions = append(suggestions, c.SuggestionFunction(annotation, d)...)
	}
	return prompt.FilterHasPrefix(suggestions, d.GetWordBeforeCursor(), true)
}
