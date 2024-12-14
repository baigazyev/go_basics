package services

import (
	"e-commerce/internal/models"
	"e-commerce/internal/repositories"
)

type ShoppingCartService interface {
	GetAllCarts() ([]models.ShoppingCart, error)
	GetCartByID(cartID int) (*models.ShoppingCart, error)
	GetCartByUserID(userID int) (*models.ShoppingCart, error)
	CreateCart(cart *models.ShoppingCart) error
	DeleteCart(cartID int) error
}

type shoppingCartService struct {
	repo repositories.ShoppingCartRepository
}

func NewShoppingCartService(repo repositories.ShoppingCartRepository) ShoppingCartService {
	return &shoppingCartService{repo: repo}
}

func (s *shoppingCartService) GetAllCarts() ([]models.ShoppingCart, error) {
	return s.repo.GetAllCarts()
}

func (s *shoppingCartService) GetCartByID(cartID int) (*models.ShoppingCart, error) {
	return s.repo.GetCartByID(cartID)
}

func (s *shoppingCartService) GetCartByUserID(userID int) (*models.ShoppingCart, error) {
	return s.repo.GetCartByUserID(userID)
}

func (s *shoppingCartService) CreateCart(cart *models.ShoppingCart) error {
	return s.repo.CreateCart(cart)
}

func (s *shoppingCartService) DeleteCart(cartID int) error {
	return s.repo.DeleteCart(cartID)
}
