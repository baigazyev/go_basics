package services

import (
	"e-commerce/internal/models"
	"e-commerce/internal/repositories"
)

type UserAddressService interface {
	GetAllUserAddresses() ([]models.UserAddress, error)
	GetUserAddressByID(addressID int) (*models.UserAddress, error)
	CreateUserAddress(address *models.UserAddress) error
	UpdateUserAddress(address *models.UserAddress) error
	DeleteUserAddress(addressID int) error
}

type userAddressService struct {
	repo repositories.UserAddressRepository
}

func NewUserAddressService(repo repositories.UserAddressRepository) UserAddressService {
	return &userAddressService{repo: repo}
}

func (s *userAddressService) GetAllUserAddresses() ([]models.UserAddress, error) {
	return s.repo.GetAllUserAddresses()
}

func (s *userAddressService) GetUserAddressByID(addressID int) (*models.UserAddress, error) {
	return s.repo.GetUserAddressByID(addressID)
}

func (s *userAddressService) CreateUserAddress(address *models.UserAddress) error {
	return s.repo.CreateUserAddress(address)
}

func (s *userAddressService) UpdateUserAddress(address *models.UserAddress) error {
	return s.repo.UpdateUserAddress(address)
}

func (s *userAddressService) DeleteUserAddress(addressID int) error {
	return s.repo.DeleteUserAddress(addressID)
}
