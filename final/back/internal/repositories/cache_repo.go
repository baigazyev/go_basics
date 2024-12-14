package repositories

import (
	"e-commerce/internal/models"
	"time"

	"gorm.io/gorm"
)

type CacheRepositoryInterface interface {
	GetCacheByKey(cacheKey string) (*models.Cache, error)
	SetCache(cache *models.Cache) error
	DeleteCache(cacheKey string) error
	ClearExpiredCaches() error
}

type CacheRepository struct {
	db *gorm.DB
}

func NewCacheRepository(db *gorm.DB) CacheRepositoryInterface {
	return &CacheRepository{db: db}
}

// Реализация метода SetCache
func (r *CacheRepository) SetCache(cache *models.Cache) error {
	return r.db.Save(cache).Error
}

// Реализация метода GetCacheByKey
func (r *CacheRepository) GetCacheByKey(cacheKey string) (*models.Cache, error) {
	var cache models.Cache
	err := r.db.Where("cache_key = ?", cacheKey).First(&cache).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	// Проверить срок действия кэша
	if time.Now().After(cache.ExpirationTime) {
		_ = r.db.Delete(&cache) // Удалить устаревший кэш
		return nil, nil
	}

	return &cache, nil
}

// Реализация метода DeleteCache
func (r *CacheRepository) DeleteCache(cacheKey string) error {
	return r.db.Where("cache_key = ?", cacheKey).Delete(&models.Cache{}).Error
}

// Реализация метода ClearExpiredCaches
func (r *CacheRepository) ClearExpiredCaches() error {
	// Удалить записи, срок действия которых истёк
	return r.db.Where("expiration_time <= ?", time.Now()).Delete(&models.Cache{}).Error
}
