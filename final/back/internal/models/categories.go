package models

type Category struct {
	CategoryID  int    `json:"category_id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
