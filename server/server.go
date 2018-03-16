package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/whywaita/yayoi/irkit"

	"go.uber.org/zap"
)

type ErrorJSON struct {
	Error string `json:"error"`
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	jsonStr := `{"health": "true"}`

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, jsonStr)
}

func Run(logger *zap.Logger, manager irkit.Manager) {
	router := mux.NewRouter()
	router.Path("/health").HandlerFunc(healthCheck)

	rIrkit := router.PathPrefix("/irkit").Subrouter()
	rIrkit.HandleFunc("/{device}/{switch:on|off}", func(w http.ResponseWriter, r *http.Request) {
		irkitHTTPHandle(w, r, logger, manager)
	})
	rIrkit.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
		showIRKitDevices(w, r, logger, manager)
	})

	if err := http.ListenAndServe(":8080", router); err != nil {
		logger.Fatal("[ERROR]", zap.Error(err))
	}
}
