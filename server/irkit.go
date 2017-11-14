package server

import (
	"net/http"

	"go.uber.org/zap"

	"github.com/gorilla/mux"
	"github.com/whywaita/yayoi/irkit"
)

func irkitHTTPHandle(w http.ResponseWriter, r *http.Request, logger *zap.Logger, deviceList []irkit.Device) {
	var err error
	vars := mux.Vars(r)
	s := vars["switch"]
	deviceName := vars["device"]

	for _, device := range deviceList {
		if deviceName == device.Name {
			if s == "on" {
				err = device.OnCommand()
			} else {
				err = device.OffCommand()
			}
		}
	}

	if err != nil {
		logger.Warn("[ERROR]", zap.Error(err))
	}
}
