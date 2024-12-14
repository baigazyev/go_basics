package handlers

import (
	"e-commerce/internal/models"
	"e-commerce/internal/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type OrderItemHandler struct {
	service services.OrderItemService
}

func NewOrderItemHandler(service services.OrderItemService) *OrderItemHandler {
	return &OrderItemHandler{service: service}
}

func (h *OrderItemHandler) GetAllOrderItems(w http.ResponseWriter, r *http.Request) {
	orderItems, err := h.service.GetAllOrderItems()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(orderItems)
}

func (h *OrderItemHandler) GetOrderItemsByOrderID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderID, _ := strconv.Atoi(vars["order_id"])

	orderItems, err := h.service.GetOrderItemsByOrderID(orderID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(orderItems)
}

func (h *OrderItemHandler) GetOrderItemByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderItemID, _ := strconv.Atoi(vars["id"])

	orderItem, err := h.service.GetOrderItemByID(orderItemID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if orderItem == nil {
		http.Error(w, "Order item not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(orderItem)
}

func (h *OrderItemHandler) CreateOrderItem(w http.ResponseWriter, r *http.Request) {
	var orderItem models.OrderItem
	if err := json.NewDecoder(r.Body).Decode(&orderItem); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.service.CreateOrderItem(&orderItem); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *OrderItemHandler) UpdateOrderItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderItemID, _ := strconv.Atoi(vars["id"])

	var orderItem models.OrderItem
	if err := json.NewDecoder(r.Body).Decode(&orderItem); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	orderItem.OrderItemID = orderItemID

	if err := h.service.UpdateOrderItem(&orderItem); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *OrderItemHandler) DeleteOrderItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderItemID, _ := strconv.Atoi(vars["id"])

	if err := h.service.DeleteOrderItem(orderItemID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
