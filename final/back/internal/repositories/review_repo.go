package repositories

import (
	"e-commerce/internal/models"

	"gorm.io/gorm"
)

type ReviewRepository interface {
	GetAllReviews() ([]models.Review, error)
	GetReviewsByProductID(productID int) ([]models.Review, error)
	GetReviewsByUserID(userID int) ([]models.Review, error)
	GetReviewByID(reviewID int) (*models.Review, error)
	CreateReview(review *models.Review) error
	UpdateReview(review *models.Review) error
	DeleteReview(reviewID int) error
}

type reviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) ReviewRepository {
	return &reviewRepository{db: db}
}

func (r *reviewRepository) GetAllReviews() ([]models.Review, error) {
	var reviews []models.Review
	err := r.db.Find(&reviews).Error
	return reviews, err
}

func (r *reviewRepository) GetReviewsByProductID(productID int) ([]models.Review, error) {
	var reviews []models.Review
	err := r.db.Where("product_id = ?", productID).Find(&reviews).Error
	return reviews, err
}

func (r *reviewRepository) GetReviewsByUserID(userID int) ([]models.Review, error) {
	var reviews []models.Review
	err := r.db.Where("user_id = ?", userID).Find(&reviews).Error
	return reviews, err
}

func (r *reviewRepository) GetReviewByID(reviewID int) (*models.Review, error) {
	var review models.Review
	err := r.db.First(&review, reviewID).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &review, err
}

func (r *reviewRepository) CreateReview(review *models.Review) error {
	return r.db.Create(review).Error
}

func (r *reviewRepository) UpdateReview(review *models.Review) error {
	return r.db.Save(review).Error
}

func (r *reviewRepository) DeleteReview(reviewID int) error {
	return r.db.Delete(&models.Review{}, reviewID).Error
}
