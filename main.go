package main

import (
	"flag"
	"fmt"

	"github.com/whywaita/yayoi/cli"
	"github.com/whywaita/yayoi/irkit"
	"github.com/whywaita/yayoi/server"
	"go.uber.org/zap"
)

var (
	mode       = flag.String("m", "server", "change mode, cli / server")
	deviceList = []irkit.Device{}
)

func main() {
	logger, _ := zap.NewProduction()
	flag.Parse()
	deviceList = irkit.MakeDeviceList()
	manager := irkit.NewManager()

	if *mode == "cli" {
		cli.Start(logger)
	} else {
		fmt.Println("Server mode start...")
		server.Run(logger, manager)
	}

}
