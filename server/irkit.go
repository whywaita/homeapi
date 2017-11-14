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
	s := vars["switch"] // on or off
	deviceName := vars["device"]

end:
	for _, device := range deviceList {
		if deviceName == device.Name {
			if s == "on" {
				err = device.OnCommand()
				device.Status = true
			} else {
				err = device.OffCommand()
				device.Status = false
			}

			if err != nil {
				logger.Warn("msg", zap.Error(err))
			} else {
				logger.Info("msg", zap.Bool("status", device.Status))
			}

			break end
		}
	}
}
