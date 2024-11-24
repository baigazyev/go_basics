package handlers

import (
	"fmt"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorHandler(w, r, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to the Home Page!"))
}

// DashboardHandler handles requests to the secure dashboard route.
func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve user ID from context (set by JWT middleware)
	userID := r.Context().Value("userID")
	if userID == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Example response for the dashboard
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome to the secure dashboard, User ID: %v", userID)
}
