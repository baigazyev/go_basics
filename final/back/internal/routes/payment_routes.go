package routes

import (
	"e-commerce/internal/handlers"
	"e-commerce/internal/repositories"
	"e-commerce/internal/services"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterPaymentRoutes(router *mux.Router, db *gorm.DB) {
	repo := repositories.NewPaymentRepository(db)
	service := services.NewPaymentService(repo)
	handler := handlers.NewPaymentHandler(service)

	router.HandleFunc("/payments", handler.GetAllPayments).Methods("GET")
	router.HandleFunc("/payments/{id:[0-9]+}", handler.GetPaymentByID).Methods("GET")
	router.HandleFunc("/payments", handler.CreatePayment).Methods("POST")
	router.HandleFunc("/payments/{id:[0-9]+}", handler.UpdatePayment).Methods("PUT")
	router.HandleFunc("/payments/{id:[0-9]+}", handler.DeletePayment).Methods("DELETE")
}
