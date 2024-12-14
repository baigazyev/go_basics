package handlers

import (
	"e-commerce/internal/models"
	"e-commerce/internal/services"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type SessionHandler struct {
	service services.SessionService
}

func NewSessionHandler(service services.SessionService) *SessionHandler {
	return &SessionHandler{service: service}
}

func (h *SessionHandler) GetSessionByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sessionID := vars["id"]

	session, err := h.service.GetSessionByID(sessionID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if session == nil {
		http.Error(w, "Session not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(session)
}

func (h *SessionHandler) CreateSession(w http.ResponseWriter, r *http.Request) {
	var session models.Session
	if err := json.NewDecoder(r.Body).Decode(&session); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Set expiration time if not provided
	if session.ExpiresAt.IsZero() {
		session.ExpiresAt = time.Now().Add(24 * time.Hour)
	}

	if err := h.service.CreateSession(&session); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *SessionHandler) DeleteSession(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sessionID := vars["id"]

	if err := h.service.DeleteSession(sessionID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *SessionHandler) DeleteExpiredSessions(w http.ResponseWriter, r *http.Request) {
	if err := h.service.DeleteExpiredSessions(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
