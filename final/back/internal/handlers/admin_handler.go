package handlers

import (
	"e-commerce/internal/services"
	"encoding/json"
	"net/http"
)

type AdminHandler struct {
	service *services.AdminService
}

func NewAdminHandler(service *services.AdminService) *AdminHandler {
	return &AdminHandler{service: service}
}

func (h *AdminHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (h *AdminHandler) GetTotalRevenue(w http.ResponseWriter, r *http.Request) {
	revenue, err := h.service.GetTotalRevenue()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]float64{"total_revenue": revenue})
}

func (h *AdminHandler) ManageUsers(w http.ResponseWriter, r *http.Request) {
	// Logic for managing users (e.g., fetch, update, delete)
	users, err := h.service.GetAllUsers() // Example service method
	if err != nil {
		http.Error(w, "Failed to fetch users: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
