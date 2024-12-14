package handlers

import (
	"e-commerce/internal/models"
	"e-commerce/internal/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type UserAddressHandler struct {
	service services.UserAddressService
}

func NewUserAddressHandler(service services.UserAddressService) *UserAddressHandler {
	return &UserAddressHandler{service: service}
}

func (h *UserAddressHandler) GetAllUserAddresses(w http.ResponseWriter, r *http.Request) {
	addresses, err := h.service.GetAllUserAddresses()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(addresses)
}

func (h *UserAddressHandler) GetUserAddressByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	addressID, _ := strconv.Atoi(vars["id"])

	address, err := h.service.GetUserAddressByID(addressID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if address == nil {
		http.Error(w, "User address not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(address)
}

func (h *UserAddressHandler) CreateUserAddress(w http.ResponseWriter, r *http.Request) {
	var address models.UserAddress
	if err := json.NewDecoder(r.Body).Decode(&address); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.service.CreateUserAddress(&address); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *UserAddressHandler) UpdateUserAddress(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	addressID, _ := strconv.Atoi(vars["id"])

	var address models.UserAddress
	if err := json.NewDecoder(r.Body).Decode(&address); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	address.AddressID = addressID

	if err := h.service.UpdateUserAddress(&address); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *UserAddressHandler) DeleteUserAddress(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	addressID, _ := strconv.Atoi(vars["id"])

	if err := h.service.DeleteUserAddress(addressID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
