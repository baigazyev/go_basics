package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"e-commerce/internal/models"
	"e-commerce/internal/services"

	"github.com/gorilla/mux"
)

type PaymentHandler struct {
	Service services.PaymentService
}

func NewPaymentHandler(service services.PaymentService) *PaymentHandler {
	return &PaymentHandler{Service: service}
}

// GetAllPayments handles GET /payments
func (h *PaymentHandler) GetAllPayments(w http.ResponseWriter, r *http.Request) {
	payments, err := h.Service.GetAllPayments()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(payments)
}

// GetPaymentByID handles GET /payments/{id}
func (h *PaymentHandler) GetPaymentByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	paymentID, _ := strconv.Atoi(vars["id"])

	payment, err := h.Service.GetPaymentByID(paymentID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if payment == nil {
		http.Error(w, "Payment not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(payment)
}

// CreatePayment handles POST /payments
func (h *PaymentHandler) CreatePayment(w http.ResponseWriter, r *http.Request) {
	var payment models.Payment
	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.Service.CreatePayment(&payment); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// UpdatePayment handles PUT /payments/{id}
func (h *PaymentHandler) UpdatePayment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	paymentID, _ := strconv.Atoi(vars["id"])

	var payment models.Payment
	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	payment.PaymentID = paymentID

	if err := h.Service.UpdatePayment(&payment); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// DeletePayment handles DELETE /payments/{id}
func (h *PaymentHandler) DeletePayment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	paymentID, _ := strconv.Atoi(vars["id"])

	if err := h.Service.DeletePayment(paymentID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
