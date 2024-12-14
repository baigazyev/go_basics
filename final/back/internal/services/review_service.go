package services

import (
	"e-commerce/internal/models"
	"e-commerce/internal/repositories"
)

type ReviewService interface {
	GetAllReviews() ([]models.Review, error)
	GetReviewsByProductID(productID int) ([]models.Review, error)
	GetReviewsByUserID(userID int) ([]models.Review, error)
	GetReviewByID(reviewID int) (*models.Review, error)
	CreateReview(review *models.Review) error
	UpdateReview(review *models.Review) error
	DeleteReview(reviewID int) error
}

type reviewService struct {
	repo repositories.ReviewRepository
}

func NewReviewService(repo repositories.ReviewRepository) ReviewService {
	return &reviewService{repo: repo}
}

func (s *reviewService) GetAllReviews() ([]models.Review, error) {
	return s.repo.GetAllReviews()
}

func (s *reviewService) GetReviewsByProductID(productID int) ([]models.Review, error) {
	return s.repo.GetReviewsByProductID(productID)
}

func (s *reviewService) GetReviewsByUserID(userID int) ([]models.Review, error) {
	return s.repo.GetReviewsByUserID(userID)
}

func (s *reviewService) GetReviewByID(reviewID int) (*models.Review, error) {
	return s.repo.GetReviewByID(reviewID)
}

func (s *reviewService) CreateReview(review *models.Review) error {
	return s.repo.CreateReview(review)
}

func (s *reviewService) UpdateReview(review *models.Review) error {
	return s.repo.UpdateReview(review)
}

func (s *reviewService) DeleteReview(reviewID int) error {
	return s.repo.DeleteReview(reviewID)
}
