package server

import (
	"net/http"

	"go.uber.org/zap"

	"github.com/gorilla/mux"
	"github.com/whywaita/homeapi/irkit"
)

func AirCon(w http.ResponseWriter, r *http.Request, logger *zap.Logger) {
	vars := mux.Vars(r)
	s := vars["switch"]

	if s == "on" {
		irkit.AirConOn(logger)
	} else {
		irkit.AirConOff(logger)
	}
}

func HomeLight(w http.ResponseWriter, r *http.Request, logger *zap.Logger) {
	vars := mux.Vars(r)
	s := vars["switch"]

	if s == "on" {
		irkit.HomeLightOn(logger)
	} else {
		irkit.HomeLightOff(logger)
	}
}

func TVPower(w http.ResponseWriter, r *http.Request, logger *zap.Logger) {
	vars := mux.Vars(r)
	s := vars["switch"]

	if s == "on" {
		irkit.TVPowerToggle(logger)
	} else {
		irkit.TVPowerToggle(logger)
	}
}
