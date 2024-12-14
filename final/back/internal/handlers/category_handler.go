package handlers

import (
	"e-commerce/internal/models"
	"e-commerce/internal/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type CategoryHandler struct {
	service services.CategoryService
}

func NewCategoryHandler(service services.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: service}
}

func (h *CategoryHandler) GetAllCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := h.service.GetAllCategories()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(categories)
}

func (h *CategoryHandler) GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	categoryID, _ := strconv.Atoi(vars["id"])

	category, err := h.service.GetCategoryByID(categoryID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if category == nil {
		http.Error(w, "Category not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(category)
}

func (h *CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category models.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.service.CreateCategory(&category); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *CategoryHandler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	categoryID, _ := strconv.Atoi(vars["id"])

	var category models.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	category.CategoryID = categoryID

	if err := h.service.UpdateCategory(&category); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *CategoryHandler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	categoryID, _ := strconv.Atoi(vars["id"])

	if err := h.service.DeleteCategory(categoryID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
