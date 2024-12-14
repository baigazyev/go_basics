package routes

import (
	"e-commerce/internal/handlers"
	"e-commerce/internal/repositories"
	"e-commerce/internal/services"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterUserRoutes(router *mux.Router, db *gorm.DB) {
	repo := repositories.NewUserRepository(db)
	service := services.NewUserService(repo)
	handler := handlers.NewUserHandler(service)

	router.HandleFunc("/users", handler.GetAllUsers).Methods("GET")
	router.HandleFunc("/users/{id:[0-9]+}", handler.GetUserByID).Methods("GET")
	router.HandleFunc("/users", handler.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id:[0-9]+}", handler.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id:[0-9]+}", handler.DeleteUser).Methods("DELETE")
}
