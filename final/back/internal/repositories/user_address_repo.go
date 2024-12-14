package repositories

import (
	"e-commerce/internal/models"

	"gorm.io/gorm"
)

type UserAddressRepository interface {
	GetAllUserAddresses() ([]models.UserAddress, error)
	GetUserAddressByID(addressID int) (*models.UserAddress, error)
	CreateUserAddress(address *models.UserAddress) error
	UpdateUserAddress(address *models.UserAddress) error
	DeleteUserAddress(addressID int) error
}

type userAddressRepository struct {
	db *gorm.DB
}

func NewUserAddressRepository(db *gorm.DB) UserAddressRepository {
	return &userAddressRepository{db: db}
}

func (r *userAddressRepository) GetAllUserAddresses() ([]models.UserAddress, error) {
	var addresses []models.UserAddress
	err := r.db.Find(&addresses).Error
	return addresses, err
}

func (r *userAddressRepository) GetUserAddressByID(addressID int) (*models.UserAddress, error) {
	var address models.UserAddress
	err := r.db.First(&address, addressID).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &address, err
}

func (r *userAddressRepository) CreateUserAddress(address *models.UserAddress) error {
	return r.db.Create(address).Error
}

func (r *userAddressRepository) UpdateUserAddress(address *models.UserAddress) error {
	return r.db.Save(address).Error
}

func (r *userAddressRepository) DeleteUserAddress(addressID int) error {
	return r.db.Delete(&models.UserAddress{}, addressID).Error
}
