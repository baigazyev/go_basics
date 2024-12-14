package services

import (
	"e-commerce/internal/models"
	"e-commerce/internal/repositories"
)

type ProductService interface {
	GetAllProducts() ([]models.Product, error)
	GetProductByID(productID int) (*models.Product, error)
	CreateProduct(product *models.Product) error
	UpdateProduct(product *models.Product) error
	DeleteProduct(productID int) error
}

type productService struct {
	repo repositories.ProductRepository
}

func NewProductService(repo repositories.ProductRepository) ProductService {
	return &productService{repo: repo}
}

func (s *productService) GetAllProducts() ([]models.Product, error) {
	return s.repo.GetAllProducts()
}

func (s *productService) GetProductByID(productID int) (*models.Product, error) {
	return s.repo.GetProductByID(productID)
}

func (s *productService) CreateProduct(product *models.Product) error {
	return s.repo.CreateProduct(product)
}

func (s *productService) UpdateProduct(product *models.Product) error {
	return s.repo.UpdateProduct(product)
}

func (s *productService) DeleteProduct(productID int) error {
	return s.repo.DeleteProduct(productID)
}
