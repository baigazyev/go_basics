package routes

import (
	"e-commerce/internal/handlers"
	"e-commerce/internal/repositories"
	"e-commerce/internal/services"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterCartItemRoutes(router *mux.Router, db *gorm.DB) {
	repo := repositories.NewCartItemRepository(db)
	service := services.NewCartItemService(repo)
	handler := handlers.NewCartItemHandler(service)

	router.HandleFunc("/cart-items", handler.GetAllCartItems).Methods("GET")
	router.HandleFunc("/cart-items/cart/{cart_id:[0-9]+}", handler.GetCartItemsByCartID).Methods("GET")
	router.HandleFunc("/cart-items/{id:[0-9]+}", handler.GetCartItemByID).Methods("GET")
	router.HandleFunc("/cart-items", handler.AddCartItem).Methods("POST")
	router.HandleFunc("/cart-items/{id:[0-9]+}", handler.UpdateCartItem).Methods("PUT")
	router.HandleFunc("/cart-items/{id:[0-9]+}", handler.DeleteCartItem).Methods("DELETE")
	router.HandleFunc("/cart-items/cart/{cart_id:[0-9]+}/clear", handler.ClearCart).Methods("POST")
}
