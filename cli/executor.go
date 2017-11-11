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
	} else if len(ss) == 2 {
		commandExecutor(ss)
	} else {
	}

	return
}

func commandExecutor(s []string) {
	logger, _ := zap.NewProduction()

	if len(s) != 2 {
		fmt.Println("command length must be 2")
		os.Exit(1)
	}

	commandType := s[0]
	switchType := s[1] // maybe on or off

	if commandType == "light" {
		if switchType == "on" {
			fmt.Println("HomeLight On...")
			irkit.HomeLightOn(logger)
		} else if switchType == "off" {
			fmt.Println("HomeLight Off...")
			irkit.HomeLightOff(logger)
		}
	}

}
