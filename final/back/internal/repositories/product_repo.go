package repositories

import (
	"e-commerce/internal/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	GetAllProducts() ([]models.Product, error)
	GetProductByID(productID int) (*models.Product, error)
	CreateProduct(product *models.Product) error
	UpdateProduct(product *models.Product) error
	DeleteProduct(productID int) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

// GetAllProducts retrieves all products
func (r *productRepository) GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	err := r.db.Find(&products).Error
	return products, err
}

// GetProductByID retrieves a product by its ID
func (r *productRepository) GetProductByID(productID int) (*models.Product, error) {
	var product models.Product
	err := r.db.First(&product, productID).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &product, err
}

// CreateProduct creates a new product
func (r *productRepository) CreateProduct(product *models.Product) error {
	return r.db.Create(product).Error
}

// UpdateProduct updates an existing product
func (r *productRepository) UpdateProduct(product *models.Product) error {
	return r.db.Save(product).Error
}

// DeleteProduct deletes a product by its ID
func (r *productRepository) DeleteProduct(productID int) error {
	return r.db.Delete(&models.Product{}, productID).Error
}
