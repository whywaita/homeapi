package main

import (
	"flag"

	"github.com/whywaita/homeapi/cli"
	"github.com/whywaita/homeapi/server"
	"go.uber.org/zap"
)

var (
	mode = flag.String("m", "server", "change mode, cli / server")
)

func main() {
	logger, _ := zap.NewProduction()
	flag.Parse()

	if *mode == "cli" {
		cli.Start(logger)
	} else {
		server.Run(logger)
	}

}
