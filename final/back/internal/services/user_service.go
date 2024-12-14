package services

import (
	"e-commerce/internal/models"
	"e-commerce/internal/repositories"
	"errors"
)

type UserService interface {
	GetOrderDetails() ([]models.OrderDetails, error)
	GetTotalRevenue() (float64, error)
	GetAllUsers() ([]models.User, error)
	GetUserByID(userID int) (*models.User, error)
	CreateUser(user *models.User) error
	UpdateUser(user *models.User) error
	DeleteUser(userID int) error
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) GetOrderDetails() ([]models.OrderDetails, error) {
	return s.repo.GetOrderDetails()
}

func (s *userService) GetTotalRevenue() (float64, error) {
	return s.repo.GetTotalRevenue()
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAllUsers()
}

func (s *userService) GetUserByID(userID int) (*models.User, error) {
	return s.repo.GetUserByID(userID)
}

func (s *userService) CreateUser(user *models.User) error {
	// Check if the email is already in use
	existingUser, err := s.repo.GetUserByEmail(user.Email)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return errors.New("email already in use")
	}

	// Save the user to the database
	return s.repo.CreateUser(user)
}

func (s *userService) UpdateUser(user *models.User) error {
	return s.repo.UpdateUser(user)
}

func (s *userService) DeleteUser(userID int) error {
	return s.repo.DeleteUser(userID)
}
