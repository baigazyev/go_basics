package models

import "time"

type ProductImage struct {
	ImageID   int       `json:"image_id" db:"image_id"`
	ProductID int       `json:"product_id" db:"product_id"`
	ImageURL  string    `json:"image_url" db:"image_url"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
