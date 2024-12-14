package services

import (
	"e-commerce/internal/models"
	"e-commerce/internal/repositories"
)

type CartItemService interface {
	GetAllCartItems() ([]models.CartItem, error)
	GetCartItemsByCartID(cartID int) ([]models.CartItem, error)
	GetCartItemByID(cartItemID int) (*models.CartItem, error)
	AddCartItem(cartItem *models.CartItem) error
	UpdateCartItem(cartItem *models.CartItem) error
	DeleteCartItem(cartItemID int) error
	ClearCart(cartID int) error
}

type cartItemService struct {
	repo repositories.CartItemRepository
}

func NewCartItemService(repo repositories.CartItemRepository) CartItemService {
	return &cartItemService{repo: repo}
}

func (s *cartItemService) GetAllCartItems() ([]models.CartItem, error) {
	return s.repo.GetAllCartItems()
}

func (s *cartItemService) GetCartItemsByCartID(cartID int) ([]models.CartItem, error) {
	return s.repo.GetCartItemsByCartID(cartID)
}

func (s *cartItemService) GetCartItemByID(cartItemID int) (*models.CartItem, error) {
	return s.repo.GetCartItemByID(cartItemID)
}

func (s *cartItemService) AddCartItem(cartItem *models.CartItem) error {
	return s.repo.AddCartItem(cartItem)
}

func (s *cartItemService) UpdateCartItem(cartItem *models.CartItem) error {
	return s.repo.UpdateCartItem(cartItem)
}

func (s *cartItemService) DeleteCartItem(cartItemID int) error {
	return s.repo.DeleteCartItem(cartItemID)
}

func (s *cartItemService) ClearCart(cartID int) error {
	return s.repo.ClearCart(cartID)
}
