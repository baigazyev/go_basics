package repositories

import (
	"e-commerce/internal/models"

	"gorm.io/gorm"
)

type ShoppingCartRepository interface {
	GetAllCarts() ([]models.ShoppingCart, error)
	GetCartByID(cartID int) (*models.ShoppingCart, error)
	GetCartByUserID(userID int) (*models.ShoppingCart, error)
	CreateCart(cart *models.ShoppingCart) error
	DeleteCart(cartID int) error
}

type shoppingCartRepository struct {
	db *gorm.DB
}

func NewShoppingCartRepository(db *gorm.DB) ShoppingCartRepository {
	return &shoppingCartRepository{db: db}
}

func (r *shoppingCartRepository) GetAllCarts() ([]models.ShoppingCart, error) {
	var carts []models.ShoppingCart
	err := r.db.Find(&carts).Error
	return carts, err
}

func (r *shoppingCartRepository) GetCartByID(cartID int) (*models.ShoppingCart, error) {
	var cart models.ShoppingCart
	err := r.db.First(&cart, cartID).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &cart, err
}

func (r *shoppingCartRepository) GetCartByUserID(userID int) (*models.ShoppingCart, error) {
	var cart models.ShoppingCart
	err := r.db.Where("user_id = ?", userID).First(&cart).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &cart, err
}

func (r *shoppingCartRepository) CreateCart(cart *models.ShoppingCart) error {
	return r.db.Create(cart).Error
}

func (r *shoppingCartRepository) DeleteCart(cartID int) error {
	return r.db.Delete(&models.ShoppingCart{}, cartID).Error
}
