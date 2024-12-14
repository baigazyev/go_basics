package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"e-commerce/internal/models"
	"e-commerce/internal/services"

	"github.com/gorilla/mux"
)

type ProductHandler struct {
	Service services.ProductService
}

func NewProductHandler(service services.ProductService) *ProductHandler {
	return &ProductHandler{Service: service}
}

func (h *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.Service.GetAllProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(products)
}

func (h *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID, _ := strconv.Atoi(vars["id"])

	product, err := h.Service.GetProductByID(productID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if product == nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.Service.CreateProduct(&product); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID, _ := strconv.Atoi(vars["id"])

	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	product.ProductID = productID

	if err := h.Service.UpdateProduct(&product); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID, _ := strconv.Atoi(vars["id"])

	if err := h.Service.DeleteProduct(productID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
