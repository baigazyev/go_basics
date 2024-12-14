package handlers

import (
	"e-commerce/internal/models"
	"e-commerce/internal/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type RoleHandler struct {
	service services.RoleService
}

func NewRoleHandler(service services.RoleService) *RoleHandler {
	return &RoleHandler{service: service}
}

func (h *RoleHandler) GetAllRoles(w http.ResponseWriter, r *http.Request) {
	roles, err := h.service.GetAllRoles()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(roles)
}

func (h *RoleHandler) GetRoleByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	roleID, _ := strconv.Atoi(vars["id"])

	role, err := h.service.GetRoleByID(roleID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if role == nil {
		http.Error(w, "Role not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(role)
}

func (h *RoleHandler) CreateRole(w http.ResponseWriter, r *http.Request) {
	var role models.Role
	if err := json.NewDecoder(r.Body).Decode(&role); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.service.CreateRole(&role); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *RoleHandler) UpdateRole(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	roleID, _ := strconv.Atoi(vars["id"])

	var role models.Role
	if err := json.NewDecoder(r.Body).Decode(&role); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	role.RoleID = roleID

	if err := h.service.UpdateRole(&role); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *RoleHandler) DeleteRole(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	roleID, _ := strconv.Atoi(vars["id"])

	if err := h.service.DeleteRole(roleID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
