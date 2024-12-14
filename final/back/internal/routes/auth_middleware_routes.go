package routes

import (
	"e-commerce/internal/handlers"
	"e-commerce/internal/middleware"
	"e-commerce/internal/repositories"
	"e-commerce/internal/services"
	"log"
	"os"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterAuthMiddlewareRoutes(router *mux.Router, db *gorm.DB) {
	// Load JWT secret securely
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET environment variable is not set")
	}

	// Initialize repositories
	userRepo := repositories.NewUserRepository(db)

	// Initialize services
	adminService := services.NewAdminService(userRepo)
	sessionRepo := repositories.NewSessionRepository(db)
	authService := services.NewAuthService(userRepo, sessionRepo, jwtSecret)
	userService := services.NewUserService(userRepo) // Add UserService for customer routes

	// Initialize handlers
	adminHandler := handlers.NewAdminHandler(adminService)
	authHandler := handlers.NewAuthHandler(authService)
	userHandler := handlers.NewUserHandler(userService) // Add UserHandler for customer routes

	// Public Routes
	router.HandleFunc("/login", authHandler.Login).Methods("POST")

	// Admin Routes (Protected)
	adminRouter := router.PathPrefix("/admin").Subrouter()
	adminRouter.Use(middleware.RoleMiddleware("admin")) // Only admin users can access
	adminRouter.HandleFunc("/manage-users", adminHandler.ManageUsers).Methods("GET")

	// Customer Routes (Protected)
	customerRouter := router.PathPrefix("/customer").Subrouter()
	customerRouter.Use(middleware.RoleMiddleware("customer", "admin"))                    // Both customers and admins can access
	customerRouter.HandleFunc("/view-orders", userHandler.GetOrderDetails).Methods("GET") // Use UserHandler for customer actions
}
