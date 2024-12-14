package repositories

import (
	"e-commerce/internal/models"

	"gorm.io/gorm"
)

type OrderItemRepository interface {
	GetAllOrderItems() ([]models.OrderItem, error)
	GetOrderItemsByOrderID(orderID int) ([]models.OrderItem, error)
	GetOrderItemByID(orderItemID int) (*models.OrderItem, error)
	CreateOrderItem(orderItem *models.OrderItem) error
	UpdateOrderItem(orderItem *models.OrderItem) error
	DeleteOrderItem(orderItemID int) error
}

type orderItemRepository struct {
	db *gorm.DB
}

func NewOrderItemRepository(db *gorm.DB) OrderItemRepository {
	return &orderItemRepository{db: db}
}

func (r *orderItemRepository) GetAllOrderItems() ([]models.OrderItem, error) {
	var orderItems []models.OrderItem
	err := r.db.Find(&orderItems).Error
	return orderItems, err
}

func (r *orderItemRepository) GetOrderItemsByOrderID(orderID int) ([]models.OrderItem, error) {
	var orderItems []models.OrderItem
	err := r.db.Where("order_id = ?", orderID).Find(&orderItems).Error
	return orderItems, err
}

func (r *orderItemRepository) GetOrderItemByID(orderItemID int) (*models.OrderItem, error) {
	var orderItem models.OrderItem
	err := r.db.First(&orderItem, orderItemID).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &orderItem, err
}

func (r *orderItemRepository) CreateOrderItem(orderItem *models.OrderItem) error {
	return r.db.Create(orderItem).Error
}

func (r *orderItemRepository) UpdateOrderItem(orderItem *models.OrderItem) error {
	return r.db.Save(orderItem).Error
}

func (r *orderItemRepository) DeleteOrderItem(orderItemID int) error {
	return r.db.Delete(&models.OrderItem{}, orderItemID).Error
}
