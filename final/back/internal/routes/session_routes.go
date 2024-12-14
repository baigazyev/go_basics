package routes

import (
	"e-commerce/internal/handlers"
	"e-commerce/internal/repositories"
	"e-commerce/internal/services"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterSessionRoutes(router *mux.Router, db *gorm.DB) {
	repo := repositories.NewSessionRepository(db)
	service := services.NewSessionService(repo)
	handler := handlers.NewSessionHandler(service)

	router.HandleFunc("/sessions/{id}", handler.GetSessionByID).Methods("GET")
	router.HandleFunc("/sessions", handler.CreateSession).Methods("POST")
	router.HandleFunc("/sessions/{id}", handler.DeleteSession).Methods("DELETE")
	router.HandleFunc("/sessions/clear-expired", handler.DeleteExpiredSessions).Methods("POST")
}
