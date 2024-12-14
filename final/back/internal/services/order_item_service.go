package services

import (
	"e-commerce/internal/models"
	"e-commerce/internal/repositories"
)

type OrderItemService interface {
	GetAllOrderItems() ([]models.OrderItem, error)
	GetOrderItemsByOrderID(orderID int) ([]models.OrderItem, error)
	GetOrderItemByID(orderItemID int) (*models.OrderItem, error)
	CreateOrderItem(orderItem *models.OrderItem) error
	UpdateOrderItem(orderItem *models.OrderItem) error
	DeleteOrderItem(orderItemID int) error
}

type orderItemService struct {
	repo repositories.OrderItemRepository
}

func NewOrderItemService(repo repositories.OrderItemRepository) OrderItemService {
	return &orderItemService{repo: repo}
}

func (s *orderItemService) GetAllOrderItems() ([]models.OrderItem, error) {
	return s.repo.GetAllOrderItems()
}

func (s *orderItemService) GetOrderItemsByOrderID(orderID int) ([]models.OrderItem, error) {
	return s.repo.GetOrderItemsByOrderID(orderID)
}

func (s *orderItemService) GetOrderItemByID(orderItemID int) (*models.OrderItem, error) {
	return s.repo.GetOrderItemByID(orderItemID)
}

func (s *orderItemService) CreateOrderItem(orderItem *models.OrderItem) error {
	return s.repo.CreateOrderItem(orderItem)
}

func (s *orderItemService) UpdateOrderItem(orderItem *models.OrderItem) error {
	return s.repo.UpdateOrderItem(orderItem)
}

func (s *orderItemService) DeleteOrderItem(orderItemID int) error {
	return s.repo.DeleteOrderItem(orderItemID)
}
