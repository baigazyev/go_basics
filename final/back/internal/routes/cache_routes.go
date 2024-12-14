package routes

import (
	"e-commerce/internal/handlers"
	"e-commerce/internal/repositories"
	"e-commerce/internal/services"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterCacheRoutes(router *mux.Router, db *gorm.DB) {
	repo := repositories.NewCacheRepository(db)
	service := services.NewCacheService(repo)
	handler := handlers.NewCacheHandler(service)

	router.HandleFunc("/cache/{key}", handler.GetCacheByKey).Methods("GET")
	router.HandleFunc("/cache", handler.SetCache).Methods("POST")
	router.HandleFunc("/cache/{key}", handler.DeleteCache).Methods("DELETE")
	router.HandleFunc("/cache/clear-expired", handler.ClearExpiredCaches).Methods("POST")
}
