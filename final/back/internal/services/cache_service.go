package services

import (
	"e-commerce/internal/models"
	"e-commerce/internal/repositories"
)

// Определяем интерфейс CacheService
type CacheService interface {
	GetCacheByKey(cacheKey string) (*models.Cache, error)
	SetCache(cache *models.Cache) error
	DeleteCache(cacheKey string) error
	ClearExpiredCaches() error
}

// Структура, реализующая интерфейс CacheService
type cacheService struct {
	repo repositories.CacheRepositoryInterface
}

// Конструктор для CacheService
func NewCacheService(repo repositories.CacheRepositoryInterface) CacheService {
	return &cacheService{repo: repo}
}

// Реализация метода GetCacheByKey
func (s *cacheService) GetCacheByKey(cacheKey string) (*models.Cache, error) {
	return s.repo.GetCacheByKey(cacheKey)
}

// Реализация метода SetCache
func (s *cacheService) SetCache(cache *models.Cache) error {
	return s.repo.SetCache(cache)
}

// Реализация метода DeleteCache
func (s *cacheService) DeleteCache(cacheKey string) error {
	return s.repo.DeleteCache(cacheKey)
}

// Реализация метода ClearExpiredCaches
func (s *cacheService) ClearExpiredCaches() error {
	return s.repo.ClearExpiredCaches()
}
