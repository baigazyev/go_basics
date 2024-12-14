package routes

import (
	"e-commerce/internal/handlers"
	"e-commerce/internal/repositories"
	"e-commerce/internal/services"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterCategoryRoutes(router *mux.Router, db *gorm.DB) {
	repo := repositories.NewCategoryRepository(db)
	service := services.NewCategoryService(repo)
	handler := handlers.NewCategoryHandler(service)

	router.HandleFunc("/categories", handler.GetAllCategories).Methods("GET")
	router.HandleFunc("/categories/{id:[0-9]+}", handler.GetCategoryByID).Methods("GET")
	router.HandleFunc("/categories", handler.CreateCategory).Methods("POST")
	router.HandleFunc("/categories/{id:[0-9]+}", handler.UpdateCategory).Methods("PUT")
	router.HandleFunc("/categories/{id:[0-9]+}", handler.DeleteCategory).Methods("DELETE")
}
