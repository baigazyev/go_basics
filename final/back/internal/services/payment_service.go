package services

import (
	"e-commerce/internal/models"
	"e-commerce/internal/repositories"
)

type PaymentService interface {
	GetAllPayments() ([]models.Payment, error)
	GetPaymentByID(paymentID int) (*models.Payment, error)
	CreatePayment(payment *models.Payment) error
	UpdatePayment(payment *models.Payment) error
	DeletePayment(paymentID int) error
}

type paymentService struct {
	repo repositories.PaymentRepository
}

func NewPaymentService(repo repositories.PaymentRepository) PaymentService {
	return &paymentService{repo: repo}
}

func (s *paymentService) GetAllPayments() ([]models.Payment, error) {
	return s.repo.GetAllPayments()
}

func (s *paymentService) GetPaymentByID(paymentID int) (*models.Payment, error) {
	return s.repo.GetPaymentByID(paymentID)
}

func (s *paymentService) CreatePayment(payment *models.Payment) error {
	return s.repo.CreatePayment(payment)
}

func (s *paymentService) UpdatePayment(payment *models.Payment) error {
	return s.repo.UpdatePayment(payment)
}

func (s *paymentService) DeletePayment(paymentID int) error {
	return s.repo.DeletePayment(paymentID)
}
