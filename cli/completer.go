package cli

import (
	"strings"

	prompt "github.com/c-bata/go-prompt"
)

var commands = []prompt.Suggest{
	{Text: "irkit", Description: "irkit commands"},
	{Text: "version", Description: "Display version"},

	{Text: "exit", Description: "Exit this program"},
}

var irkitCommands = []prompt.Suggest{
	{Text: "light", Description: ""},
	{Text: "aircon", Description: ""},
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
	if len(args) == 2 {
		return prompt.FilterHasPrefix(irkitCommands, args[1], true)
	}

	first := args[0]
	switch first {
	case "irkit":
		second := args[1]
		switch second {
		case "light":
			third := args[2]
			if len(args) == 3 {
				subcommands := []prompt.Suggest{
					{Text: "on"},
					{Text: "off"},
				}
				return prompt.FilterHasPrefix(subcommands, third, true)
			}

		case "aircon":
			third := args[2]
			if len(args) == 3 {
				subcommands := []prompt.Suggest{
					{Text: "on"},
					{Text: "off"},
				}
				return prompt.FilterHasPrefix(subcommands, third, true)
			}

		default:
			return []prompt.Suggest{}
		}

	default:
		return []prompt.Suggest{}
	}

	return []prompt.Suggest{}
}
