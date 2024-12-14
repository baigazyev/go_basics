package services

import (
	"e-commerce/internal/models"
	"e-commerce/internal/repositories"
	"time"
)

type OrderService interface {
	GetOrderDetails() ([]models.OrderDetails, error)
	GetTotalRevenue() (float64, error)
	GetRevenueByStatus() ([]models.RevenueByStatus, error)
	GetOrdersInDateRange(startDate, endDate time.Time) ([]models.OrderDetails, error)
	GetAvgOrderValuePerUser() ([]models.AvgOrderValue, error)
	GetOrdersByRole() ([]models.OrdersByRole, error)
	GetDetailedOrders() ([]models.DetailedOrder, error)
}

type orderService struct {
	repo repositories.OrderRepository
}

func NewOrderService(repo repositories.OrderRepository) OrderService {
	return &orderService{repo: repo}
}

func (s *orderService) GetOrderDetails() ([]models.OrderDetails, error) {
	return s.repo.GetOrderDetails()
}

func (s *orderService) GetTotalRevenue() (float64, error) {
	return s.repo.GetTotalRevenue()
}

func (s *orderService) GetRevenueByStatus() ([]models.RevenueByStatus, error) {
	return s.repo.GetRevenueByStatus()
}

func (s *orderService) GetOrdersInDateRange(startDate, endDate time.Time) ([]models.OrderDetails, error) {
	return s.repo.GetOrdersInDateRange(startDate, endDate)
}

func (s *orderService) GetAvgOrderValuePerUser() ([]models.AvgOrderValue, error) {
	return s.repo.GetAvgOrderValuePerUser()
}

func (s *orderService) GetOrdersByRole() ([]models.OrdersByRole, error) {
	return s.repo.GetOrdersByRole()
}

func (s *orderService) GetDetailedOrders() ([]models.DetailedOrder, error) {
	return s.repo.GetDetailedOrders()
}
