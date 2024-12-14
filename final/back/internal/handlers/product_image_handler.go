package handlers

import (
	"e-commerce/internal/models"
	"e-commerce/internal/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ProductImageHandler struct {
	service services.ProductImageService
}

func NewProductImageHandler(service services.ProductImageService) *ProductImageHandler {
	return &ProductImageHandler{service: service}
}

func (h *ProductImageHandler) GetAllProductImages(w http.ResponseWriter, r *http.Request) {
	images, err := h.service.GetAllProductImages()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(images)
}

func (h *ProductImageHandler) GetProductImageByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	imageID, _ := strconv.Atoi(vars["id"])

	image, err := h.service.GetProductImageByID(imageID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if image == nil {
		http.Error(w, "Product image not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(image)
}

func (h *ProductImageHandler) CreateProductImage(w http.ResponseWriter, r *http.Request) {
	var image models.ProductImage
	if err := json.NewDecoder(r.Body).Decode(&image); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.service.CreateProductImage(&image); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *ProductImageHandler) UpdateProductImage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	imageID, _ := strconv.Atoi(vars["id"])

	var image models.ProductImage
	if err := json.NewDecoder(r.Body).Decode(&image); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	image.ImageID = imageID

	if err := h.service.UpdateProductImage(&image); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *ProductImageHandler) DeleteProductImage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	imageID, _ := strconv.Atoi(vars["id"])

	if err := h.service.DeleteProductImage(imageID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
