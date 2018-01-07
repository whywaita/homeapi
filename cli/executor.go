package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/whywaita/yayoi/irkit"
	"github.com/whywaita/yayoi/version"
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
	if len(s) != 3 {
		fmt.Println("irkit's command length must be 3")
		os.Exit(1)
	}

	commandType := s[1]
	switchType := s[2] // maybe on or off

	if commandType == "light" {
		if switchType == "on" {
			fmt.Println("HomeLight On...")
			irkit.HomeLightOn()
		} else if switchType == "off" {
			fmt.Println("HomeLight Off...")
			irkit.HomeLightOff()
		} else {
			fmt.Print("invaild command")
		}
	} else if commandType == "tvpower" {
		fmt.Println("TV Power change...")
		irkit.TVPowerToggle()
	} else if commandType == "aircon" {
		if switchType == "on" {
			fmt.Println("Aircon On...")
			irkit.AirConOn()
		} else if switchType == "off" {
			fmt.Println("Aircon Off...")
			irkit.AirConOff()
		} else {
			fmt.Print("invaild command")
		}
	}

}
