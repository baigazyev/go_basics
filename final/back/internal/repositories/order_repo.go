package repositories

import (
	"e-commerce/internal/models"
	"time"

	"gorm.io/gorm"
)

type OrderRepository interface {
	GetOrderDetails() ([]models.OrderDetails, error)
	GetTotalRevenue() (float64, error)
	GetRevenueByStatus() ([]models.RevenueByStatus, error)
	GetOrdersInDateRange(startDate, endDate time.Time) ([]models.OrderDetails, error)
	GetAvgOrderValuePerUser() ([]models.AvgOrderValue, error)
	GetOrdersByRole() ([]models.OrdersByRole, error)
	GetDetailedOrders() ([]models.DetailedOrder, error)
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) GetOrderDetails() ([]models.OrderDetails, error) {
	var details []models.OrderDetails
	err := r.db.Table("orders").
		Select("orders.order_id, users.username, orders.order_date, orders.status, orders.total_amount").
		Joins("join users on users.user_id = orders.user_id").
		Scan(&details).Error
	return details, err
}

func (r *orderRepository) GetTotalRevenue() (float64, error) {
	var totalRevenue float64
	err := r.db.Table("orders").
		Select("SUM(total_amount) as total_revenue").
		Scan(&totalRevenue).Error
	return totalRevenue, err
}

func (r *orderRepository) GetRevenueByStatus() ([]models.RevenueByStatus, error) {
	var revenue []models.RevenueByStatus
	err := r.db.Table("orders").
		Select("status, SUM(total_amount) as total_revenue").
		Group("status").
		Scan(&revenue).Error
	return revenue, err
}

func (r *orderRepository) GetOrdersInDateRange(startDate, endDate time.Time) ([]models.OrderDetails, error) {
	var orders []models.OrderDetails
	err := r.db.Table("orders").
		Select("orders.order_id, users.username, orders.order_date, orders.status, orders.total_amount").
		Joins("join users on users.user_id = orders.user_id").
		Where("orders.order_date BETWEEN ? AND ?", startDate, endDate).
		Scan(&orders).Error
	return orders, err
}

func (r *orderRepository) GetAvgOrderValuePerUser() ([]models.AvgOrderValue, error) {
	var result []models.AvgOrderValue
	err := r.db.Table("orders").
		Select("users.username, AVG(orders.total_amount) as average_order_value").
		Joins("join users on users.user_id = orders.user_id").
		Group("users.username").
		Scan(&result).Error
	return result, err
}

func (r *orderRepository) GetOrdersByRole() ([]models.OrdersByRole, error) {
	var result []models.OrdersByRole
	err := r.db.Table("orders").
		Select("users.role, COUNT(orders.order_id) as order_count").
		Joins("join users on users.user_id = orders.user_id").
		Group("users.role").
		Scan(&result).Error
	return result, err
}

func (r *orderRepository) GetDetailedOrders() ([]models.DetailedOrder, error) {
	var details []models.DetailedOrder
	err := r.db.Table("orders").
		Select("orders.order_id, users.username, categories.name as category_name, orders.order_date, orders.total_amount").
		Joins("join users on users.user_id = orders.user_id").
		Joins("join categories on categories.category_id = orders.category_id").
		Scan(&details).Error
	return details, err
}
