package routes

import (
	"e-commerce/internal/handlers"
	"e-commerce/internal/repositories"
	"e-commerce/internal/services"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterUserAddressRoutes(router *mux.Router, db *gorm.DB) {
	repo := repositories.NewUserAddressRepository(db)
	service := services.NewUserAddressService(repo)
	handler := handlers.NewUserAddressHandler(service)

	router.HandleFunc("/user-addresses", handler.GetAllUserAddresses).Methods("GET")
	router.HandleFunc("/user-addresses/{id:[0-9]+}", handler.GetUserAddressByID).Methods("GET")
	router.HandleFunc("/user-addresses", handler.CreateUserAddress).Methods("POST")
	router.HandleFunc("/user-addresses/{id:[0-9]+}", handler.UpdateUserAddress).Methods("PUT")
	router.HandleFunc("/user-addresses/{id:[0-9]+}", handler.DeleteUserAddress).Methods("DELETE")
}
