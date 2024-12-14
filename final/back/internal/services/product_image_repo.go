package services

import (
	"e-commerce/internal/models"
	"e-commerce/internal/repositories"
)

type ProductImageService interface {
	GetAllProductImages() ([]models.ProductImage, error)
	GetProductImageByID(imageID int) (*models.ProductImage, error)
	CreateProductImage(image *models.ProductImage) error
	UpdateProductImage(image *models.ProductImage) error
	DeleteProductImage(imageID int) error
}

type productImageService struct {
	repo repositories.ProductImageRepository
}

func NewProductImageService(repo repositories.ProductImageRepository) ProductImageService {
	return &productImageService{repo: repo}
}

func (s *productImageService) GetAllProductImages() ([]models.ProductImage, error) {
	return s.repo.GetAllProductImages()
}

func (s *productImageService) GetProductImageByID(imageID int) (*models.ProductImage, error) {
	return s.repo.GetProductImageByID(imageID)
}

func (s *productImageService) CreateProductImage(image *models.ProductImage) error {
	return s.repo.CreateProductImage(image)
}

func (s *productImageService) UpdateProductImage(image *models.ProductImage) error {
	return s.repo.UpdateProductImage(image)
}

func (s *productImageService) DeleteProductImage(imageID int) error {
	return s.repo.DeleteProductImage(imageID)
}
