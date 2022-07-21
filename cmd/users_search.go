package cmd

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/buger/jsonparser"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func jsonSearch(key string, value string) {
	data, err := os.ReadFile("./docs/data/users.json")
	check(err)

	jsonparser.ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		fmt.Println(jsonparser.Get(value, "url"))
	}, "_id")

	// if er != nil {
	// 	return
	// }
	fmt.Println(string(value), value)
}

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for users in json file",
	Run: func(cmd *cobra.Command, args []string) {
		collection := []string{
			"_id",
			"url",
			"external_id",
			"name",
			"alias",
			"created_at",
			"active",
			"verified",
			"shared",
			"locale",
			"timezone",
			"last_login_at",
			"email",
			"phone",
			"signature",
			"organization_id",
			"tags",
			"suspended",
			"role",
		}

		targetPrompt := promptui.Select{
			Label: "Please choose a field to search for",
			Items: collection,
		}

		_, result, err := targetPrompt.Run()

		switch result {
		case "_id":
			validate := func(input string) error {
				_, err := strconv.ParseFloat(input, 64)
				if err != nil {
					return errors.New("sorry but the id must be a number")
				}
				return nil
			}

			idPrompt := promptui.Prompt{
				Label:    "Please enter the id of user you want to search, it must be a number. Use backspace to clear your input.",
				Validate: validate,
			}

			idResult, err := idPrompt.Run()

			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}

			fmt.Printf("You chose id %q\n", idResult)

			// jsonSearch("", "")
			return
		}

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
	},
}

func init() {
	RootCmd.AddCommand(searchCmd)
}
