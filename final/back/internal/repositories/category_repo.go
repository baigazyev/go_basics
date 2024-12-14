package repositories

import (
	"e-commerce/internal/models"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetAllCategories() ([]models.Category, error)
	GetCategoryByID(categoryID int) (*models.Category, error)
	CreateCategory(category *models.Category) error
	UpdateCategory(category *models.Category) error
	DeleteCategory(categoryID int) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) GetAllCategories() ([]models.Category, error) {
	var categories []models.Category
	err := r.db.Find(&categories).Error
	return categories, err
}

func (r *categoryRepository) GetCategoryByID(categoryID int) (*models.Category, error) {
	var category models.Category
	err := r.db.First(&category, categoryID).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &category, err
}

func (r *categoryRepository) CreateCategory(category *models.Category) error {
	return r.db.Create(category).Error
}

func (r *categoryRepository) UpdateCategory(category *models.Category) error {
	return r.db.Save(category).Error
}

func (r *categoryRepository) DeleteCategory(categoryID int) error {
	return r.db.Delete(&models.Category{}, categoryID).Error
}
