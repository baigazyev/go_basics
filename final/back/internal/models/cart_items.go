package models

type CartItem struct {
	CartItemID int `json:"cart_item_id" db:"cart_item_id"`
	CartID     int `json:"cart_id" db:"cart_id"`
	ProductID  int `json:"product_id" db:"product_id"`
	Quantity   int `json:"quantity" db:"quantity"`
}
