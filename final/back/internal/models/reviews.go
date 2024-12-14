package models

import "time"

type Review struct {
	ReviewID  int       `json:"review_id" db:"review_id"`
	ProductID int       `json:"product_id" db:"product_id"`
	UserID    int       `json:"user_id" db:"user_id"`
	Rating    int       `json:"rating" db:"rating"`
	Comment   string    `json:"comment" db:"comment"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
