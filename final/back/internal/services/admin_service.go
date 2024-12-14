package services

import (
	"e-commerce/internal/models"
	"e-commerce/internal/repositories"
)

type AdminService struct {
	repo repositories.UserRepository
}

func NewAdminService(repo repositories.UserRepository) *AdminService {
	return &AdminService{repo: repo}
}

// GetAllUsers retrieves all users from the repository
func (s *AdminService) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAllUsers()
}

// GetTotalRevenue calculates the total revenue from the repository
func (s *AdminService) GetTotalRevenue() (float64, error) {
	return s.repo.GetTotalRevenue()
}
