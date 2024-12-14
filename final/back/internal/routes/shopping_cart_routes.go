package routes

import (
	"e-commerce/internal/handlers"
	"e-commerce/internal/repositories"
	"e-commerce/internal/services"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterShoppingCartRoutes(router *mux.Router, db *gorm.DB) {
	repo := repositories.NewShoppingCartRepository(db)
	service := services.NewShoppingCartService(repo)
	handler := handlers.NewShoppingCartHandler(service)

	router.HandleFunc("/shopping-carts", handler.GetAllCarts).Methods("GET")
	router.HandleFunc("/shopping-carts/{id:[0-9]+}", handler.GetCartByID).Methods("GET")
	router.HandleFunc("/shopping-carts/user/{user_id:[0-9]+}", handler.GetCartByUserID).Methods("GET")
	router.HandleFunc("/shopping-carts", handler.CreateCart).Methods("POST")
	router.HandleFunc("/shopping-carts/{id:[0-9]+}", handler.DeleteCart).Methods("DELETE")
}
