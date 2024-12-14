package handlers

import (
	"e-commerce/internal/models"
	"e-commerce/internal/services"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type CacheHandler struct {
	service services.CacheService
}

func NewCacheHandler(service services.CacheService) *CacheHandler {
	return &CacheHandler{service: service}
}

func (h *CacheHandler) GetCacheByKey(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cacheKey := vars["key"]

	cache, err := h.service.GetCacheByKey(cacheKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if cache == nil {
		http.Error(w, "Cache not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(cache)
}

func (h *CacheHandler) SetCache(w http.ResponseWriter, r *http.Request) {
	var cache models.Cache
	if err := json.NewDecoder(r.Body).Decode(&cache); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	cache.ExpirationTime = time.Now().Add(24 * time.Hour) // Default expiration time
	if err := h.service.SetCache(&cache); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *CacheHandler) DeleteCache(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cacheKey := vars["key"]

	if err := h.service.DeleteCache(cacheKey); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *CacheHandler) ClearExpiredCaches(w http.ResponseWriter, r *http.Request) {
	if err := h.service.ClearExpiredCaches(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
