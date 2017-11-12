package cli

import (
	"fmt"
	"os"
	"strings"

	"go.uber.org/zap"

	"github.com/whywaita/homeapi/irkit"
	"github.com/whywaita/homeapi/version"
)

func Executor(s string) {
	s = strings.TrimSpace(s)
	ss := strings.Split(s, " ")
	if s == "version" {
		version.Display()
	} else if s == "quit" || s == "exit" {
		fmt.Println("Bye!")
		os.Exit(0)
		return
	} else if len(ss) == 3 && ss[0] == "irkit" {
		irkitCommandExecutor(ss)
	} else {
	}

	return
}

func irkitCommandExecutor(s []string) {
	logger, _ := zap.NewProduction()

	if len(s) != 3 {
		fmt.Println("irkit's command length must be 3")
		os.Exit(1)
	}

	commandType := s[1]
	switchType := s[2] // maybe on or off

	if commandType == "light" {
		if switchType == "on" {
			fmt.Println("HomeLight On...")
			irkit.HomeLightOn(logger)
		} else if switchType == "off" {
			fmt.Println("HomeLight Off...")
			irkit.HomeLightOff(logger)
		} else {
			fmt.Print("invaild command")
		}
	} else if commandType == "tvpower" {
		fmt.Println("TV Power change...")
		irkit.TVPowerToggle(logger)
	}

}
