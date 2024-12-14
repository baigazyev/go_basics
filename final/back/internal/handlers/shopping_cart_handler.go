package handlers

import (
	"e-commerce/internal/models"
	"e-commerce/internal/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ShoppingCartHandler struct {
	service services.ShoppingCartService
}

func NewShoppingCartHandler(service services.ShoppingCartService) *ShoppingCartHandler {
	return &ShoppingCartHandler{service: service}
}

func (h *ShoppingCartHandler) GetAllCarts(w http.ResponseWriter, r *http.Request) {
	carts, err := h.service.GetAllCarts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(carts)
}

func (h *ShoppingCartHandler) GetCartByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cartID, _ := strconv.Atoi(vars["id"])

	cart, err := h.service.GetCartByID(cartID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if cart == nil {
		http.Error(w, "Cart not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(cart)
}

func (h *ShoppingCartHandler) GetCartByUserID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, _ := strconv.Atoi(vars["user_id"])

	cart, err := h.service.GetCartByUserID(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if cart == nil {
		http.Error(w, "Cart not found for this user", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(cart)
}

func (h *ShoppingCartHandler) CreateCart(w http.ResponseWriter, r *http.Request) {
	var cart models.ShoppingCart
	if err := json.NewDecoder(r.Body).Decode(&cart); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.service.CreateCart(&cart); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *ShoppingCartHandler) DeleteCart(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cartID, _ := strconv.Atoi(vars["id"])

	if err := h.service.DeleteCart(cartID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
