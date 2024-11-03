package auth

import (
	"bookstore-api/internal/user"
	"errors"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtSecret = []byte("your_secret_key") // Change to a secure key

// GenerateJWT generates a JWT token for the given username
func GenerateJWT(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(), // Token valid for 72 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ValidateJWT validates the JWT token
func ValidateJWT(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.NewValidationError("invalid signing method", jwt.ValidationErrorMalformed)
		}
		return jwtSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["username"].(string), nil
	}
	return "", err
}

func GetUserFromToken(ctx *gin.Context) (user.User, error) {
	// Get the token from the "Authorization" header
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		return user.User{}, errors.New("authorization header not found")
	}

	// Split the header to get the token
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return user.User{}, errors.New("invalid authorization format")
	}

	tokenString := parts[1]

	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return user.User{}, errors.New("invalid token")
	}

	// Extract claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		username := claims["username"].(string) // Adjust as necessary
		// Lookup user by username in your user store
		for _, u := range user.GetAllUsers() {
			if u.Username == username {
				return u, nil
			}
		}
	}

	return user.User{}, errors.New("user not found")
}
