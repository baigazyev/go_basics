package routes

import (
	"e-commerce/internal/handlers"
	"e-commerce/internal/repositories"
	"e-commerce/internal/services"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterOrderRoutes(router *mux.Router, db *gorm.DB) {
	repo := repositories.NewOrderRepository(db)
	service := services.NewOrderService(repo)
	handler := handlers.NewOrderHandler(service)

	router.HandleFunc("/orders/details", handler.GetOrderDetails).Methods("GET")
	router.HandleFunc("/revenue/total", handler.GetTotalRevenue).Methods("GET")
	router.HandleFunc("/revenue/status", handler.GetRevenueByStatus).Methods("GET")
}
