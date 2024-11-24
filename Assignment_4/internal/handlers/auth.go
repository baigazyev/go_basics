package handlers

import (
	"encoding/json"
	"net/http"

	"myproject/internal/middleware"

	"github.com/sirupsen/logrus"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var loginReq LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		middleware.LogEvent(logrus.WarnLevel, "Invalid login request", logrus.Fields{
			"error": err.Error(),
		})
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	// Simulate successful login
	middleware.LogEvent(logrus.InfoLevel, "User logged in", logrus.Fields{
		"username": loginReq.Username,
	})

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Login successful"))
}
