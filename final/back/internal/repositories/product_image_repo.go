package repositories

import (
	"e-commerce/internal/models"

	"gorm.io/gorm"
)

type ProductImageRepository interface {
	GetAllProductImages() ([]models.ProductImage, error)
	GetProductImageByID(imageID int) (*models.ProductImage, error)
	CreateProductImage(image *models.ProductImage) error
	UpdateProductImage(image *models.ProductImage) error
	DeleteProductImage(imageID int) error
}

type productImageRepository struct {
	db *gorm.DB
}

func NewProductImageRepository(db *gorm.DB) ProductImageRepository {
	return &productImageRepository{db: db}
}

func (r *productImageRepository) GetAllProductImages() ([]models.ProductImage, error) {
	var images []models.ProductImage
	err := r.db.Find(&images).Error
	return images, err
}

func (r *productImageRepository) GetProductImageByID(imageID int) (*models.ProductImage, error) {
	var image models.ProductImage
	err := r.db.First(&image, imageID).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &image, err
}

func (r *productImageRepository) CreateProductImage(image *models.ProductImage) error {
	return r.db.Create(image).Error
}

func (r *productImageRepository) UpdateProductImage(image *models.ProductImage) error {
	return r.db.Save(image).Error
}

func (r *productImageRepository) DeleteProductImage(imageID int) error {
	return r.db.Delete(&models.ProductImage{}, imageID).Error
}
