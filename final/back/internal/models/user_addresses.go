package models

type UserAddress struct {
	AddressID int    `json:"address_id" db:"address_id"`
	UserID    int    `json:"user_id" db:"user_id"`
	Street    string `json:"street" db:"street"`
	City      string `json:"city" db:"city"`
	State     string `json:"state" db:"state"`
	ZipCode   string `json:"zip_code" db:"zip_code"`
}
