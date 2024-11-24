package main

import (
	"log"
	"net/http"

	"myproject/internal/handlers"
	"myproject/internal/middleware"

	"github.com/gorilla/mux"
)

func main() {
	middleware.InitLogger()  // Initialize logger
	middleware.InitMetrics() // Initialize Prometheus metrics

	r := mux.NewRouter()

	// Register routes
	r.HandleFunc("/", handlers.HomeHandler).Methods("GET")
	r.HandleFunc("/login", handlers.LoginHandler).Methods("POST")

	// Secure routes
	secure := r.PathPrefix("/secure").Subrouter()
	secure.Use(middleware.JWTMiddleware)
	secure.HandleFunc("/dashboard", handlers.DashboardHandler).Methods("GET")

	// Metrics endpoint
	r.Handle("/metrics", middleware.MetricsHandler())

	// Add logging and monitoring middleware
	r.Use(middleware.RequestLoggingMiddleware)
	r.Use(middleware.MetricsMiddleware)

	// Start server
	log.Println("Server is running on https://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

// package main

// import (
// 	"log"
// 	"net/http"

// 	"myproject/internal/handlers"
// 	"myproject/internal/middleware"

// 	"github.com/gorilla/mux"
// )

// func main() {
// 	r := mux.NewRouter()

// 	// Register routes
// 	r.HandleFunc("/", handlers.HomeHandler).Methods("GET")
// 	r.HandleFunc("/login", handlers.LoginHandler).Methods("POST")

// 	// Secure routes
// 	secure := r.PathPrefix("/secure").Subrouter()
// 	secure.Use(middleware.JWTMiddleware)
// 	secure.HandleFunc("/dashboard", handlers.DashboardHandler).Methods("GET")

// 	// Add CSRF Protection
// 	csrfMiddleware := middleware.CSRFProtection(r)

// 	// Start server with TLS
// 	server := &http.Server{
// 		Addr:    ":8080",
// 		Handler: csrfMiddleware,
// 	}
// 	log.Println("Server is running on https://localhost:8080")
// 	log.Fatal(server.ListenAndServeTLS("cert.pem", "key.pem"))
// }
