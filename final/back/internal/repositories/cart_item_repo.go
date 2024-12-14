package repositories

import (
	"e-commerce/internal/models"

	"gorm.io/gorm"
)

type CartItemRepository interface {
	GetAllCartItems() ([]models.CartItem, error)
	GetCartItemsByCartID(cartID int) ([]models.CartItem, error)
	GetCartItemByID(cartItemID int) (*models.CartItem, error)
	AddCartItem(cartItem *models.CartItem) error
	UpdateCartItem(cartItem *models.CartItem) error
	DeleteCartItem(cartItemID int) error
	ClearCart(cartID int) error
}

type cartItemRepository struct {
	db *gorm.DB
}

func NewCartItemRepository(db *gorm.DB) CartItemRepository {
	return &cartItemRepository{db: db}
}

func (r *cartItemRepository) GetAllCartItems() ([]models.CartItem, error) {
	var cartItems []models.CartItem
	err := r.db.Find(&cartItems).Error
	return cartItems, err
}

func (r *cartItemRepository) GetCartItemsByCartID(cartID int) ([]models.CartItem, error) {
	var cartItems []models.CartItem
	err := r.db.Where("cart_id = ?", cartID).Find(&cartItems).Error
	return cartItems, err
}

func (r *cartItemRepository) GetCartItemByID(cartItemID int) (*models.CartItem, error) {
	var cartItem models.CartItem
	err := r.db.First(&cartItem, cartItemID).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &cartItem, err
}

func (r *cartItemRepository) AddCartItem(cartItem *models.CartItem) error {
	return r.db.Create(cartItem).Error
}

func (r *cartItemRepository) UpdateCartItem(cartItem *models.CartItem) error {
	return r.db.Save(cartItem).Error
}

func (r *cartItemRepository) DeleteCartItem(cartItemID int) error {
	return r.db.Delete(&models.CartItem{}, cartItemID).Error
}

func (r *cartItemRepository) ClearCart(cartID int) error {
	return r.db.Where("cart_id = ?", cartID).Delete(&models.CartItem{}).Error
}
