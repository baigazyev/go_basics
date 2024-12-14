package models

import "time"

type Session struct {
	SessionID string    `json:"session_id" db:"session_id"` // UUID for session tracking
	UserID    int       `json:"user_id" db:"user_id"`
	Token     string    `json:"token" db:"token"` // JWT token
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	ExpiresAt time.Time `json:"expires_at" db:"expires_at"`
}
