package main

import (
	"github.com/whywaita/homeapi/cli"
	"github.com/whywaita/homeapi/server"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()

	go server.Run(logger)

	cli.Start(logger)
}
