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

func irkitHTTPHandle(w http.ResponseWriter, r *http.Request, logger *zap.Logger, manager irkit.Manager) {
	vars := mux.Vars(r)
	s := vars["switch"] // on or off
	deviceName := vars["device"]

	deviceIndex := irkit.GetIndexValue(deviceName, &manager)

	if (strings.EqualFold(s, "on")) || (strings.EqualFold(s, "off")) {
		switchIRKitDevice(w, deviceIndex, s, manager, logger)
	}
}

func switchIRKitDevice(w http.ResponseWriter, deviceIndex int, sw string, manager irkit.Manager, logger *zap.Logger) {
	w.Header().Set("Content-Type", "application/json")

	err := manager.SwitchDeviceStatus(deviceIndex, sw)

	if err != nil {
		logger.Warn("msg", zap.Error(err))
		ej := ErrorJSON{
			Error: err.Error(),
		}
		jsonBytes, _ := json.Marshal(ej)
		fmt.Fprintf(w, string(jsonBytes))
	} else {
		logger.Info("device status change", zap.String("device", manager.Devices[deviceIndex].Name), zap.Bool("status", manager.Devices[deviceIndex].Status))
		jsonBytes, _ := manager.Devices[deviceIndex].UserMarshalJSON()
		fmt.Fprintf(w, string(jsonBytes))
	}
}

func showIRKitDevices(w http.ResponseWriter, r *http.Request, logger *zap.Logger, manager irkit.Manager) {
	var list irkit.UserDeviceList
	w.Header().Set("Content-Type", "application/json")

	for _, d := range manager.Devices {
		u := d.ToUserDevice()
		list.Devices = append(list.Devices, u)
	}
	jb, err := json.Marshal(list)
	if err != nil {
		logger.Warn("msg", zap.Error(err))
	}
	fmt.Fprintf(w, string(jb))
}
