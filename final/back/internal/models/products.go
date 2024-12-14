package models

import "time"

type Product struct {
	ProductID   int       `json:"product_id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	CategoryID  int       `json:"category_id"`                      // Foreign Key, points to Category
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"` // Automatically set timestamp
}
