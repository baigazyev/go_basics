package models

import "time"

type ShoppingCart struct {
	CartID    int       `json:"cart_id" db:"cart_id"`
	UserID    int       `json:"user_id" db:"user_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
