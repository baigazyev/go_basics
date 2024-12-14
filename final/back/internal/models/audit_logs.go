package models

import "time"

type AuditLog struct {
	LogID     int       `json:"log_id" db:"log_id"`
	Action    string    `json:"action" db:"action"`
	UserID    int       `json:"user_id" db:"user_id"`
	Timestamp time.Time `json:"timestamp" db:"timestamp"`
}
