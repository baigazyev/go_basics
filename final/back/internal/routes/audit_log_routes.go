package routes

import (
	"e-commerce/internal/handlers"
	"e-commerce/internal/repositories"
	"e-commerce/internal/services"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterAuditLogRoutes(router *mux.Router, db *gorm.DB) {
	repo := repositories.NewAuditLogRepository(db)
	service := services.NewAuditLogService(repo)
	handler := handlers.NewAuditLogHandler(service)

	router.HandleFunc("/audit-logs", handler.GetAllAuditLogs).Methods("GET")
	router.HandleFunc("/audit-logs/user/{user_id:[0-9]+}", handler.GetAuditLogsByUserID).Methods("GET")
	router.HandleFunc("/audit-logs", handler.CreateAuditLog).Methods("POST")
}
