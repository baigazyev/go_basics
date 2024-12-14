package repositories

import (
	"e-commerce/internal/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetOrderDetails() ([]models.OrderDetails, error)
	GetTotalRevenue() (float64, error)
	GetAllUsers() ([]models.User, error)
	GetUserByID(userID int) (*models.User, error)
	CreateUser(user *models.User) error
	UpdateUser(user *models.User) error
	DeleteUser(userID int) error
	GetUserByEmail(email string) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetOrderDetails() ([]models.OrderDetails, error) {
	var details []models.OrderDetails
	err := r.db.Table("orders").
		Select("orders.order_id, users.username, orders.order_date, orders.status, orders.total_amount").
		Joins("join users on users.user_id = orders.user_id").
		Scan(&details).Error
	return details, err
}

func (r *userRepository) GetTotalRevenue() (float64, error) {
	var totalRevenue float64
	err := r.db.Table("orders").
		Select("SUM(total_amount) as total_revenue").
		Scan(&totalRevenue).Error
	return totalRevenue, err
}

func (r *userRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) GetUserByID(userID int) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, userID).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &user, err
}

func (r *userRepository) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) UpdateUser(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) DeleteUser(userID int) error {
	return r.db.Delete(&models.User{}, userID).Error
}

func (r *userRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &user, err
}
