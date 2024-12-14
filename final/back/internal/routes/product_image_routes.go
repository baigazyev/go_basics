package routes

import (
	"e-commerce/internal/handlers"
	"e-commerce/internal/repositories"
	"e-commerce/internal/services"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterProductImageRoutes(router *mux.Router, db *gorm.DB) {
	repo := repositories.NewProductImageRepository(db)
	service := services.NewProductImageService(repo)
	handler := handlers.NewProductImageHandler(service)

	router.HandleFunc("/product-images", handler.GetAllProductImages).Methods("GET")
	router.HandleFunc("/product-images/{id:[0-9]+}", handler.GetProductImageByID).Methods("GET")
	router.HandleFunc("/product-images", handler.CreateProductImage).Methods("POST")
	router.HandleFunc("/product-images/{id:[0-9]+}", handler.UpdateProductImage).Methods("PUT")
	router.HandleFunc("/product-images/{id:[0-9]+}", handler.DeleteProductImage).Methods("DELETE")
}
