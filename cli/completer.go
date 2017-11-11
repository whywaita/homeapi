package cli

import (
	"strings"

	prompt "github.com/c-bata/go-prompt"
)

var commands = []prompt.Suggest{
	{Text: "light", Description: ""},
	{Text: "aircon", Description: ""},
	{Text: "version", Description: "Display version"},

	{Text: "exit", Description: "Exit this program"},
}

func Completer(d prompt.Document) []prompt.Suggest {
	if d.TextBeforeCursor() == "" {
		return []prompt.Suggest{}
	}
	args := strings.Split(d.TextBeforeCursor(), " ")
	// w := d.GetWordBeforeCursor()
	// if strings.HasPrefix(w, "-") {
	// 	return optionCompleter(args, strings.HasPrefix(w, "--"))
	// }

	return argumentsCompleter(excludeOptions(args))

}

func argumentsCompleter(args []string) []prompt.Suggest {
	if len(args) <= 1 {
		return prompt.FilterHasPrefix(commands, args[0], true)
	}

	first := args[0]
	switch first {
	case "light":
		second := args[1]
		if len(args) == 2 {
			subcommands := []prompt.Suggest{
				{Text: "on"},
				{Text: "off"},
			}
			return prompt.FilterHasPrefix(subcommands, second, true)
		}

	case "aircon":
		second := args[1]
		if len(args) == 2 {
			subcommands := []prompt.Suggest{
				{Text: "on"},
				{Text: "off"},
			}
			return prompt.FilterHasPrefix(subcommands, second, true)
		}

	default:
		return []prompt.Suggest{}
	}

	return []prompt.Suggest{}

}
