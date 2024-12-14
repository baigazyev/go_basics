package services

import (
	"e-commerce/internal/models"
	"e-commerce/internal/repositories"
	"errors"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

func init() {
	if len(jwtKey) == 0 {
		panic("JWT_SECRET environment variable is not set")
	}
}

type AuthService interface {
	Authenticate(email, password string) (string, error)
	ValidateToken(tokenString string) (*models.Session, error)
	Logout(token string) error
}

type authService struct {
	repo        repositories.UserRepository
	sessionRepo repositories.SessionRepository
	jwtKey      []byte
}

func NewAuthService(userRepo repositories.UserRepository, sessionRepo repositories.SessionRepository, jwtSecret string) AuthService {
	return &authService{
		repo:        userRepo,
		sessionRepo: sessionRepo,
		jwtKey:      []byte(jwtSecret), // Convert string to byte slice
	}
}

func (s *authService) Authenticate(email, password string) (string, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		log.Printf("Error fetching user by email: %v", err)
		return "", err
	}

	if user == nil {
		log.Printf("User not found for email: %s", email)
		return "", errors.New("user not found")
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		log.Printf("Password comparison failed for user %s: %v", user.Email, err)
		return "", errors.New("invalid credentials")
	}

	// Generate JWT token
	expirationTime := time.Now().Add(24 * time.Hour)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"role":  user.Role,
		"exp":   expirationTime.Unix(),
	})

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		log.Printf("Error generating token: %v", err)
		return "", err
	}

	// Save session
	session := models.Session{
		SessionID: uuid.New().String(),
		UserID:    user.UserID,
		Token:     tokenString,
		CreatedAt: time.Now(),
		ExpiresAt: expirationTime,
	}

	if err := s.sessionRepo.SaveSession(&session); err != nil {
		log.Printf("Error saving session: %v", err)
		return "", err
	}

	return tokenString, nil
}

func (s *authService) ValidateToken(tokenString string) (*models.Session, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		log.Printf("Invalid token: %v", err)
		return nil, errors.New("invalid token")
	}

	session, err := s.sessionRepo.GetSessionByToken(tokenString)
	if err != nil || session == nil || session.ExpiresAt.Before(time.Now()) {
		log.Printf("Session expired or not found: %v", err)
		return nil, errors.New("session expired or not found")
	}

	return session, nil
}

func (s *authService) Logout(token string) error {
	return s.sessionRepo.DeleteSessionByToken(token)
}
