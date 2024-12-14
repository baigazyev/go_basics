package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	OrderID     int       `json:"order_id" gorm:"primaryKey"`
	UserID      int       `json:"user_id" gorm:"index"`                    // Поле для связи с таблицей users
	User        User      `json:"user,omitempty" gorm:"foreignKey:UserID"` // Внешний ключ
	OrderDate   time.Time `json:"order_date"`
	Status      string    `json:"status"`
	TotalAmount float64   `json:"total_amount"`
}

type OrderDetails struct {
	OrderID     int       `json:"order_id"`
	Username    string    `json:"username"`
	OrderDate   time.Time `json:"order_date"`
	Status      string    `json:"status"`
	TotalAmount float64   `json:"total_amount"`
}

func GetOrderDetails(db *gorm.DB) ([]OrderDetails, error) {
	var details []OrderDetails

	err := db.Table("orders").
		Select("orders.order_id, users.username, orders.order_date, orders.status, orders.total_amount").
		Joins("join users on users.user_id = orders.user_id").
		Scan(&details).Error

	return details, err
}

func GetTotalRevenue(db *gorm.DB) (float64, error) {
	var totalRevenue float64

	err := db.Table("orders").
		Select("SUM(total_amount) as total_revenue").
		Scan(&totalRevenue).Error

	return totalRevenue, err
}

type RevenueByStatus struct {
	Status       string  `json:"status"`
	TotalRevenue float64 `json:"total_revenue"`
}

func GetRevenueByStatus(db *gorm.DB) ([]RevenueByStatus, error) {
	var revenue []RevenueByStatus

	err := db.Table("orders").
		Select("status, SUM(total_amount) as total_revenue").
		Group("status").
		Scan(&revenue).Error

	return revenue, err
}

func GetOrdersInDateRange(db *gorm.DB, startDate, endDate time.Time) ([]OrderDetails, error) {
	var orders []OrderDetails

	err := db.Table("orders").
		Select("orders.order_id, users.username, orders.order_date, orders.status, orders.total_amount").
		Joins("join users on users.user_id = orders.user_id").
		Where("orders.order_date BETWEEN ? AND ?", startDate, endDate).
		Scan(&orders).Error

	return orders, err
}

type AvgOrderValue struct {
	Username          string  `json:"username"`
	AverageOrderValue float64 `json:"average_order_value"`
}

func GetAvgOrderValuePerUser(db *gorm.DB) ([]AvgOrderValue, error) {
	var result []AvgOrderValue

	err := db.Table("orders").
		Select("users.username, AVG(orders.total_amount) as average_order_value").
		Joins("join users on users.user_id = orders.user_id").
		Group("users.username").
		Scan(&result).Error

	return result, err
}

type OrdersByRole struct {
	Role       string `json:"role"`
	OrderCount int    `json:"order_count"`
}

func GetOrdersByRole(db *gorm.DB) ([]OrdersByRole, error) {
	var result []OrdersByRole

	err := db.Table("orders").
		Select("users.role, COUNT(orders.order_id) as order_count").
		Joins("join users on users.user_id = orders.user_id").
		Group("users.role").
		Scan(&result).Error

	return result, err
}

type DetailedOrder struct {
	OrderID      int       `json:"order_id"`
	Username     string    `json:"username"`
	CategoryName string    `json:"category_name"`
	OrderDate    time.Time `json:"order_date"`
	TotalAmount  float64   `json:"total_amount"`
}

func GetDetailedOrders(db *gorm.DB) ([]DetailedOrder, error) {
	var details []DetailedOrder

	err := db.Table("orders").
		Select("orders.order_id, users.username, categories.name as category_name, orders.order_date, orders.total_amount").
		Joins("join users on users.user_id = orders.user_id").
		Joins("join categories on categories.category_id = orders.category_id").
		Scan(&details).Error

	return details, err
}
