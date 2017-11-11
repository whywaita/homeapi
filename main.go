package main

import (
	"os"

	"github.com/whywaita/homeapi/cli"
	"github.com/whywaita/homeapi/server"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()

	go server.Run(logger)

	cli.Start(os.Args, logger)
}
