package routes

import (
	"e-commerce/internal/handlers"
	"e-commerce/internal/repositories"
	"e-commerce/internal/services"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterReviewRoutes(router *mux.Router, db *gorm.DB) {
	repo := repositories.NewReviewRepository(db)
	service := services.NewReviewService(repo)
	handler := handlers.NewReviewHandler(service)

	router.HandleFunc("/reviews", handler.GetAllReviews).Methods("GET")
	router.HandleFunc("/reviews/product/{product_id:[0-9]+}", handler.GetReviewsByProductID).Methods("GET")
	router.HandleFunc("/reviews/user/{user_id:[0-9]+}", handler.GetReviewsByUserID).Methods("GET")
	router.HandleFunc("/reviews/{id:[0-9]+}", handler.GetReviewByID).Methods("GET")
	router.HandleFunc("/reviews", handler.CreateReview).Methods("POST")
	router.HandleFunc("/reviews/{id:[0-9]+}", handler.UpdateReview).Methods("PUT")
	router.HandleFunc("/reviews/{id:[0-9]+}", handler.DeleteReview).Methods("DELETE")
}
