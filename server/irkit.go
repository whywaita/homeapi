package server

import (
	"encoding/json"
	"fmt"
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
				ej := ErrorJSON{
					Error: err.Error(),
				}
				jsonBytes, _ := json.Marshal(ej)
				fmt.Fprintf(w, string(jsonBytes))
			} else {
				logger.Info("device status change", zap.String("device", device.Name), zap.Bool("status", device.Status))
				jsonBytes, _ := device.UserMarshalJSON()
				fmt.Fprintf(w, string(jsonBytes))
			}

			break end
		}
	}
}

func showIRKitDevices(w http.ResponseWriter, r *http.Request, logger *zap.Logger, deviceList []irkit.Device) {
	var list irkit.UserDeviceList
	for _, d := range deviceList {
		u := d.ToUserDevice()
		list.Devices = append(list.Devices, u)
	}
	jb, err := json.Marshal(list)
	if err != nil {
		logger.Warn("msg", zap.Error(err))
	}
	fmt.Fprintf(w, string(jb))
}
