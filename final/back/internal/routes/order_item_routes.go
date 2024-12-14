package routes

import (
	"e-commerce/internal/handlers"
	"e-commerce/internal/repositories"
	"e-commerce/internal/services"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterOrderItemRoutes(router *mux.Router, db *gorm.DB) {
	repo := repositories.NewOrderItemRepository(db)
	service := services.NewOrderItemService(repo)
	handler := handlers.NewOrderItemHandler(service)

	router.HandleFunc("/order-items", handler.GetAllOrderItems).Methods("GET")
	router.HandleFunc("/order-items/order/{order_id:[0-9]+}", handler.GetOrderItemsByOrderID).Methods("GET")
	router.HandleFunc("/order-items/{id:[0-9]+}", handler.GetOrderItemByID).Methods("GET")
	router.HandleFunc("/order-items", handler.CreateOrderItem).Methods("POST")
	router.HandleFunc("/order-items/{id:[0-9]+}", handler.UpdateOrderItem).Methods("PUT")
	router.HandleFunc("/order-items/{id:[0-9]+}", handler.DeleteOrderItem).Methods("DELETE")
}
