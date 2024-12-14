package routes

import (
	"e-commerce/internal/handlers"
	"e-commerce/internal/repositories"
	"e-commerce/internal/services"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterProductRoutes(router *mux.Router, db *gorm.DB) {
	repo := repositories.NewProductRepository(db)
	service := services.NewProductService(repo)
	handler := handlers.NewProductHandler(service)

	router.HandleFunc("/api/products", handler.GetAllProducts).Methods("GET")
	router.HandleFunc("/api/products/{id:[0-9]+}", handler.GetProductByID).Methods("GET")
	router.HandleFunc("/api/products", handler.CreateProduct).Methods("POST")
	router.HandleFunc("/api/products/{id:[0-9]+}", handler.UpdateProduct).Methods("PUT")
	router.HandleFunc("/api/products/{id:[0-9]+}", handler.DeleteProduct).Methods("DELETE")
}
