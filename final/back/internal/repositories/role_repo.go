package repositories

import (
	"e-commerce/internal/models"

	"gorm.io/gorm"
)

type RoleRepository interface {
	GetAllRoles() ([]models.Role, error)
	GetRoleByID(roleID int) (*models.Role, error)
	CreateRole(role *models.Role) error
	UpdateRole(role *models.Role) error
	DeleteRole(roleID int) error
}

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{db: db}
}

func (r *roleRepository) GetAllRoles() ([]models.Role, error) {
	var roles []models.Role
	err := r.db.Find(&roles).Error
	return roles, err
}

func (r *roleRepository) GetRoleByID(roleID int) (*models.Role, error) {
	var role models.Role
	err := r.db.First(&role, roleID).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &role, err
}

func (r *roleRepository) CreateRole(role *models.Role) error {
	return r.db.Create(role).Error
}

func (r *roleRepository) UpdateRole(role *models.Role) error {
	return r.db.Save(role).Error
}

func (r *roleRepository) DeleteRole(roleID int) error {
	return r.db.Delete(&models.Role{}, roleID).Error
}
