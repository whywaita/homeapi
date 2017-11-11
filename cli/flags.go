package cli

import (
	"flag"
)

type CommandLineOptions struct {
	ShowVersion bool
}

func SetupFlagSet(name string, options *CommandLineOptions) *flag.FlagSet {
	flagSet := flag.NewFlagSet(name, flag.ContinueOnError)
	flagSet.BoolVar(&options.ShowVersion, "version", false, "show version")
	return flagSet
}
