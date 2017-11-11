package cli

import (
	"fmt"

	prompt "github.com/c-bata/go-prompt"

	"go.uber.org/zap"
)

func Start(logger *zap.Logger) {
	defer fmt.Println("Goodbye!")

	p := prompt.New(
		Executor,
		Completer,
		prompt.OptionTitle("home-api: interactive home-api client"),
		prompt.OptionPrefix(">>> "),
		prompt.OptionInputTextColor(prompt.Yellow),
		prompt.OptionSelectedSuggestionBGColor(prompt.LightGray),
		prompt.OptionSuggestionBGColor(prompt.DarkGray),
	)

	p.Run()
}
