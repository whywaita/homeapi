package main

import (
	"flag"

	"github.com/whywaita/yayoi/cli"
	"github.com/whywaita/yayoi/server"
	"go.uber.org/zap"
)

var (
	mode = flag.String("m", "server", "change mode, cli / server")
)

func main() {
	logger, _ := zap.NewProduction()
	flag.Parse()

	_ = Init()

	if *mode == "cli" {
		cli.Start(logger)
	} else {
		server.Run(logger)
	}

}
