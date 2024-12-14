package services

import (
	"e-commerce/internal/models"
	"e-commerce/internal/repositories"
)

type CategoryService interface {
	GetAllCategories() ([]models.Category, error)
	GetCategoryByID(categoryID int) (*models.Category, error)
	CreateCategory(category *models.Category) error
	UpdateCategory(category *models.Category) error
	DeleteCategory(categoryID int) error
}

type categoryService struct {
	repo repositories.CategoryRepository
}

func NewCategoryService(repo repositories.CategoryRepository) CategoryService {
	return &categoryService{repo: repo}
}

func (s *categoryService) GetAllCategories() ([]models.Category, error) {
	return s.repo.GetAllCategories()
}

func (s *categoryService) GetCategoryByID(categoryID int) (*models.Category, error) {
	return s.repo.GetCategoryByID(categoryID)
}

func (s *categoryService) CreateCategory(category *models.Category) error {
	return s.repo.CreateCategory(category)
}

func (s *categoryService) UpdateCategory(category *models.Category) error {
	return s.repo.UpdateCategory(category)
}

func (s *categoryService) DeleteCategory(categoryID int) error {
	return s.repo.DeleteCategory(categoryID)
}
