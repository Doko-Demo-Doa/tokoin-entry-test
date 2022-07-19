package cmd

import (
	"github.com/c-bata/go-prompt"
	"github.com/spf13/cobra"
)

var SearchDynamic = func(annotationValue string) []prompt.Suggest {
	if annotationValue != getFoodDynamicAnnotationValue {
		return nil
	}

	return []prompt.Suggest{
		{Text: "apple", Description: "Green apple"},
		{Text: "tomato", Description: "Red tomato"},
	}
}

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search entry",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

func init() {
	RootCmd.AddCommand(searchCmd)
	getCmd.PersistentFlags().BoolP("verbose", "v", false, "Verbose log")
}
