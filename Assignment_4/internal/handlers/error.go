package handlers

import (
	"log"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, status int, message string) {
	log.Printf("Error: %d - %s", status, message)
	http.Error(w, http.StatusText(status), status)
}
