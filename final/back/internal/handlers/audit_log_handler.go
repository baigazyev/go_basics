package handlers

import (
	"e-commerce/internal/models"
	"e-commerce/internal/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type AuditLogHandler struct {
	service services.AuditLogService
}

func NewAuditLogHandler(service services.AuditLogService) *AuditLogHandler {
	return &AuditLogHandler{service: service}
}

// GetAllAuditLogs retrieves all audit logs
func (h *AuditLogHandler) GetAllAuditLogs(w http.ResponseWriter, r *http.Request) {
	logs, err := h.service.GetAllAuditLogs()
	if err != nil {
		http.Error(w, "Failed to retrieve audit logs: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(logs)
}

// GetAuditLogsByUserID retrieves audit logs for a specific user by user ID
func (h *AuditLogHandler) GetAuditLogsByUserID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["user_id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	logs, err := h.service.GetAuditLogsByUserID(userID)
	if err != nil {
		http.Error(w, "Failed to retrieve audit logs: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if len(logs) == 0 {
		http.Error(w, "No audit logs found for the specified user", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(logs)
}

// CreateAuditLog creates a new audit log
func (h *AuditLogHandler) CreateAuditLog(w http.ResponseWriter, r *http.Request) {
	var log models.AuditLog
	if err := json.NewDecoder(r.Body).Decode(&log); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Validate required fields
	if log.Action == "" || log.UserID == 0 {
		http.Error(w, "Action and UserID are required fields", http.StatusBadRequest)
		return
	}

	if err := h.service.CreateAuditLog(&log); err != nil {
		http.Error(w, "Failed to create audit log: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Audit log created successfully"))
}
