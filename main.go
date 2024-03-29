package main

import (
	"os"
	"strings"

	"quan/tokoin-test/cmd"

	"github.com/c-bata/go-prompt"
	cobraprompt "github.com/stromland/cobra-prompt"
)

var advancedPrompt = &cobraprompt.CobraPrompt{
	RootCmd:                  cmd.RootCmd,
	PersistFlagValues:        true,
	ShowHelpCommandAndFlags:  true,
	DisableCompletionCommand: true,
	AddDefaultExitCommand:    true,
	GoPromptOptions: []prompt.Option{
		prompt.OptionTitle("tokoin-search"),
		prompt.OptionPrefix(">(^!^)> "),
		prompt.OptionMaxSuggestion(10),
	},

	DynamicSuggestionsFunc: func(annotationValue string, document *prompt.Document) []prompt.Suggest {
		return []prompt.Suggest{}
	},

	OnErrorFunc: func(err error) {
		if strings.Contains(err.Error(), "unknown command") {
			cmd.RootCmd.PrintErrln(err)
			return
		}

		cmd.RootCmd.PrintErr(err)
		os.Exit(1)
	},
}

var simplePrompt = &cobraprompt.CobraPrompt{
	RootCmd:                  cmd.RootCmd,
	AddDefaultExitCommand:    true,
	DisableCompletionCommand: true,
}

func main() {
	// Change to simplePrompt to see the difference
	simplePrompt.Run()
}
