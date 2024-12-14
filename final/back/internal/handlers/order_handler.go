package handlers

import (
	"encoding/json"
	"net/http"

	"e-commerce/internal/services"
)

type OrderHandler struct {
	Service services.OrderService
}

func NewOrderHandler(service services.OrderService) *OrderHandler {
	return &OrderHandler{Service: service}
}

func (h *OrderHandler) GetOrderDetails(w http.ResponseWriter, r *http.Request) {
	orders, err := h.Service.GetOrderDetails()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(orders)
}

func (h *OrderHandler) GetTotalRevenue(w http.ResponseWriter, r *http.Request) {
	revenue, err := h.Service.GetTotalRevenue()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]float64{"total_revenue": revenue})
}

func (h *OrderHandler) GetRevenueByStatus(w http.ResponseWriter, r *http.Request) {
	revenue, err := h.Service.GetRevenueByStatus()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(revenue)
}
