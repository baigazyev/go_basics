package handlers

import (
	"e-commerce/internal/models"
	"e-commerce/internal/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ReviewHandler struct {
	service services.ReviewService
}

func NewReviewHandler(service services.ReviewService) *ReviewHandler {
	return &ReviewHandler{service: service}
}

func (h *ReviewHandler) GetAllReviews(w http.ResponseWriter, r *http.Request) {
	reviews, err := h.service.GetAllReviews()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(reviews)
}

func (h *ReviewHandler) GetReviewsByProductID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID, _ := strconv.Atoi(vars["product_id"])

	reviews, err := h.service.GetReviewsByProductID(productID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(reviews)
}

func (h *ReviewHandler) GetReviewsByUserID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, _ := strconv.Atoi(vars["user_id"])

	reviews, err := h.service.GetReviewsByUserID(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(reviews)
}

func (h *ReviewHandler) GetReviewByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	reviewID, _ := strconv.Atoi(vars["id"])

	review, err := h.service.GetReviewByID(reviewID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if review == nil {
		http.Error(w, "Review not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(review)
}

func (h *ReviewHandler) CreateReview(w http.ResponseWriter, r *http.Request) {
	var review models.Review
	if err := json.NewDecoder(r.Body).Decode(&review); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.service.CreateReview(&review); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *ReviewHandler) UpdateReview(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	reviewID, _ := strconv.Atoi(vars["id"])

	var review models.Review
	if err := json.NewDecoder(r.Body).Decode(&review); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	review.ReviewID = reviewID

	if err := h.service.UpdateReview(&review); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *ReviewHandler) DeleteReview(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	reviewID, _ := strconv.Atoi(vars["id"])

	if err := h.service.DeleteReview(reviewID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
