package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("your_secret_key") // Replace with your actual secret key

// RoleMiddleware validates the JWT token and checks if the user's role is allowed.
func RoleMiddleware(allowedRoles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Extract the Authorization header
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Missing token", http.StatusUnauthorized)
				return
			}

			// Remove "Bearer " prefix from the token
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")

			// Parse the JWT token
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				// Validate the signing method
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, errors.New("Invalid signing method")
				}
				return jwtKey, nil
			})

			if err != nil || !token.Valid {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			// Extract claims from the token
			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				http.Error(w, "Invalid token claims", http.StatusUnauthorized)
				return
			}

			// Get the user's role from the token
			role, ok := claims["role"].(string)
			if !ok {
				http.Error(w, "Role not found in token", http.StatusForbidden)
				return
			}

			// Check if the user's role is allowed
			for _, allowedRole := range allowedRoles {
				if role == allowedRole {
					next.ServeHTTP(w, r) // Proceed to the next handler
					return
				}
			}

			// If the role is not allowed
			http.Error(w, "Forbidden: insufficient permissions", http.StatusForbidden)
		})
	}
}
