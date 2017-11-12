package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"go.uber.org/zap"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
	jsonStr := `{"health": "true"}`

	fmt.Fprintf(w, jsonStr)
}

func Run(logger *zap.Logger) {
	router := mux.NewRouter()
	router.Path("/health").HandlerFunc(healthCheck)

	rAPI := router.PathPrefix("/api").Subrouter()
	rIrkit := rAPI.PathPrefix("/irkit").Subrouter()

	rIrkit.HandleFunc("/aircon/{switch:on|off}", func(w http.ResponseWriter, r *http.Request) {
		AirCon(w, r, logger)
	})
	rIrkit.HandleFunc("/light/{switch:on|off}", func(w http.ResponseWriter, r *http.Request) {
		HomeLight(w, r, logger)
	})
	rIrkit.HandleFunc("/tv/{switch:on|off}", func(w http.ResponseWriter, r *http.Request) {
		TVPower(w, r, logger)
	})

	if err := http.ListenAndServe(":8080", router); err != nil {
		logger.Fatal("[ERROR]", zap.Error(err))
	}
}
