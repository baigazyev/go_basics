package handlers

import (
	"e-commerce/internal/models"
	"e-commerce/internal/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type CartItemHandler struct {
	service services.CartItemService
}

func NewCartItemHandler(service services.CartItemService) *CartItemHandler {
	return &CartItemHandler{service: service}
}

func (h *CartItemHandler) GetAllCartItems(w http.ResponseWriter, r *http.Request) {
	cartItems, err := h.service.GetAllCartItems()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(cartItems)
}

func (h *CartItemHandler) GetCartItemsByCartID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cartID, _ := strconv.Atoi(vars["cart_id"])

	cartItems, err := h.service.GetCartItemsByCartID(cartID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(cartItems)
}

func (h *CartItemHandler) GetCartItemByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cartItemID, _ := strconv.Atoi(vars["id"])

	cartItem, err := h.service.GetCartItemByID(cartItemID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if cartItem == nil {
		http.Error(w, "Cart item not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(cartItem)
}

func (h *CartItemHandler) AddCartItem(w http.ResponseWriter, r *http.Request) {
	var cartItem models.CartItem
	if err := json.NewDecoder(r.Body).Decode(&cartItem); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.service.AddCartItem(&cartItem); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *CartItemHandler) UpdateCartItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cartItemID, _ := strconv.Atoi(vars["id"])

	var cartItem models.CartItem
	if err := json.NewDecoder(r.Body).Decode(&cartItem); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	cartItem.CartItemID = cartItemID

	if err := h.service.UpdateCartItem(&cartItem); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *CartItemHandler) DeleteCartItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cartItemID, _ := strconv.Atoi(vars["id"])

	if err := h.service.DeleteCartItem(cartItemID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *CartItemHandler) ClearCart(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cartID, _ := strconv.Atoi(vars["cart_id"])

	if err := h.service.ClearCart(cartID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
