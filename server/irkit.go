package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"go.uber.org/zap"

	"github.com/gorilla/mux"
	"github.com/whywaita/yayoi/irkit"
)

func irkitHTTPHandle(w http.ResponseWriter, r *http.Request, logger *zap.Logger, deviceList []irkit.Device) {
	vars := mux.Vars(r)
	s := vars["switch"] // on or off
	deviceName := vars["device"]

	device := irkit.SearchDeviceByName(deviceName, deviceList)

	if (strings.EqualFold(s, "on")) || (strings.EqualFold(s, "off")) {
		switchIRKitDevice(w, device, s, logger)
	}
}

func switchIRKitDevice(w http.ResponseWriter, device irkit.Device, s string, logger *zap.Logger) {
	var err error
	w.Header().Set("Content-Type", "application/json")

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
}

func showIRKitDevices(w http.ResponseWriter, r *http.Request, logger *zap.Logger, deviceList []irkit.Device) {
	var list irkit.UserDeviceList
	w.Header().Set("Content-Type", "application/json")

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
