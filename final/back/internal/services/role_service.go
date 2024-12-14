package services

import (
	"e-commerce/internal/models"
	"e-commerce/internal/repositories"
)

type RoleService interface {
	GetAllRoles() ([]models.Role, error)
	GetRoleByID(roleID int) (*models.Role, error)
	CreateRole(role *models.Role) error
	UpdateRole(role *models.Role) error
	DeleteRole(roleID int) error
}

type roleService struct {
	repo repositories.RoleRepository
}

func NewRoleService(repo repositories.RoleRepository) RoleService {
	return &roleService{repo: repo}
}

func (s *roleService) GetAllRoles() ([]models.Role, error) {
	return s.repo.GetAllRoles()
}

func (s *roleService) GetRoleByID(roleID int) (*models.Role, error) {
	return s.repo.GetRoleByID(roleID)
}

func (s *roleService) CreateRole(role *models.Role) error {
	return s.repo.CreateRole(role)
}

func (s *roleService) UpdateRole(role *models.Role) error {
	return s.repo.UpdateRole(role)
}

func (s *roleService) DeleteRole(roleID int) error {
	return s.repo.DeleteRole(roleID)
}
