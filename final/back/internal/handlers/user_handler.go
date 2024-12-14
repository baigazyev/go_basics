package handlers

import (
	"e-commerce/internal/models"
	"e-commerce/internal/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	service services.UserService
}

func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetOrderDetails(w http.ResponseWriter, r *http.Request) {
	// Call the service layer to get order details
	details, err := h.service.GetOrderDetails()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(details)
}

func (h *UserHandler) GetTotalRevenue(w http.ResponseWriter, r *http.Request) {
	// Call the service layer to get total revenue
	revenue, err := h.service.GetTotalRevenue()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]float64{"total_revenue": revenue})
}

func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, _ := strconv.Atoi(vars["id"])

	user, err := h.service.GetUserByID(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if user == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	// Decode the JSON request body into the user struct
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	if user.Role == "" {
		user.Role = "user"
	}

	// Validate required fields
	if user.Email == "" || user.PasswordHash == "" {
		http.Error(w, "Email and password are required", http.StatusBadRequest)
		return
	}

	// Hash the password before saving
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}
	user.PasswordHash = string(hashedPassword)

	// Save the user using the service layer
	if err := h.service.CreateUser(&user); err != nil {
		http.Error(w, "Failed to create user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with success
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully"})
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, _ := strconv.Atoi(vars["id"])

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	user.UserID = userID

	if err := h.service.UpdateUser(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, _ := strconv.Atoi(vars["id"])

	if err := h.service.DeleteUser(userID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
