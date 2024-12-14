package routes

import (
	"e-commerce/internal/handlers"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func RegisterDataRoutes(router *mux.Router) {
	// Register the fetch-data route
	router.HandleFunc("/fetch-data", handlers.FetchDataHandler).Methods("GET")
	router.Handle("/metrics", promhttp.Handler())
}
