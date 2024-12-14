package models

import "time"

type User struct {
	UserID       int       `json:"user_id" gorm:"primaryKey"`
	Username     string    `json:"username" db:"username"`
	PasswordHash string    `json:"password_hash" db:"password_hash"`
	Email        string    `json:"email" db:"email"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	Role         string    `json:"role" db:"role"`
}
