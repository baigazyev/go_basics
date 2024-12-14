package repositories

import (
	"e-commerce/internal/models"

	"gorm.io/gorm"
)

type PaymentRepository interface {
	GetAllPayments() ([]models.Payment, error)
	GetPaymentByID(paymentID int) (*models.Payment, error)
	CreatePayment(payment *models.Payment) error
	UpdatePayment(payment *models.Payment) error
	DeletePayment(paymentID int) error
}

type paymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepository{db: db}
}

// GetAllPayments retrieves all payments
func (r *paymentRepository) GetAllPayments() ([]models.Payment, error) {
	var payments []models.Payment
	err := r.db.Find(&payments).Error
	return payments, err
}

// GetPaymentByID retrieves a payment by its ID
func (r *paymentRepository) GetPaymentByID(paymentID int) (*models.Payment, error) {
	var payment models.Payment
	err := r.db.First(&payment, paymentID).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &payment, err
}

// CreatePayment creates a new payment record
func (r *paymentRepository) CreatePayment(payment *models.Payment) error {
	return r.db.Create(payment).Error
}

// UpdatePayment updates an existing payment record
func (r *paymentRepository) UpdatePayment(payment *models.Payment) error {
	return r.db.Save(payment).Error
}

// DeletePayment deletes a payment record by its ID
func (r *paymentRepository) DeletePayment(paymentID int) error {
	return r.db.Delete(&models.Payment{}, paymentID).Error
}
