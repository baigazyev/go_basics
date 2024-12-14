package routes

import (
	"e-commerce/internal/handlers"
	"e-commerce/internal/repositories"
	"e-commerce/internal/services"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterRoleRoutes(router *mux.Router, db *gorm.DB) {
	repo := repositories.NewRoleRepository(db)
	service := services.NewRoleService(repo)
	handler := handlers.NewRoleHandler(service)

	router.HandleFunc("/roles", handler.GetAllRoles).Methods("GET")
	router.HandleFunc("/roles/{id:[0-9]+}", handler.GetRoleByID).Methods("GET")
	router.HandleFunc("/roles", handler.CreateRole).Methods("POST")
	router.HandleFunc("/roles/{id:[0-9]+}", handler.UpdateRole).Methods("PUT")
	router.HandleFunc("/roles/{id:[0-9]+}", handler.DeleteRole).Methods("DELETE")
}
