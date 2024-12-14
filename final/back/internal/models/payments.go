package models

import "time"

type Payment struct {
	PaymentID     int       `json:"payment_id" db:"payment_id"`
	OrderID       int       `json:"order_id" db:"order_id"`
	Amount        float64   `json:"amount" db:"amount"`
	PaymentDate   time.Time `json:"payment_date" db:"payment_date"`
	PaymentMethod string    `json:"payment_method" db:"payment_method"`
}
