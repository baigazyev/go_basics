package models

type OrderItem struct {
	OrderItemID int     `json:"order_item_id" db:"order_item_id"`
	OrderID     int     `json:"order_id" db:"order_id"`
	ProductID   int     `json:"product_id" db:"product_id"`
	Quantity    int     `json:"quantity" db:"quantity"`
	Price       float64 `json:"price" db:"price"`
}
