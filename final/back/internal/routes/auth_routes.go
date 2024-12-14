package routes

import (
	"e-commerce/internal/handlers"
	"e-commerce/internal/repositories"
	"e-commerce/internal/services"
	"log"
	"os"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterAuthRoutes(router *mux.Router, db *gorm.DB) {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET environment variable is not set")
	}
	userRepo := repositories.NewUserRepository(db)
	sessionRepo := repositories.NewSessionRepository(db)                     // Create the session repository
	authService := services.NewAuthService(userRepo, sessionRepo, jwtSecret) // Pass both repositories
	authHandler := handlers.NewAuthHandler(authService)

	router.HandleFunc("/login", authHandler.Login).Methods("POST")
}
