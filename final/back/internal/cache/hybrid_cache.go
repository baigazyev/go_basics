package cache

import (
	"context"
	"encoding/json"
	"time"

	"e-commerce/internal/models"
	"e-commerce/internal/repositories"

	"github.com/redis/go-redis/v9"
)

type HybridCache struct {
	redisClient *redis.Client
	dbRepo      repositories.CacheRepositoryInterface
}

func NewHybridCache(redisClient *redis.Client, dbRepo repositories.CacheRepositoryInterface) *HybridCache {
	return &HybridCache{redisClient: redisClient, dbRepo: dbRepo}
}

func (hc *HybridCache) SetCache(key string, value interface{}, ttl time.Duration) error {
	// Сериализация значения в JSON
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	// Сохранение в Redis
	err = hc.redisClient.Set(context.Background(), key, data, ttl).Err()
	if err != nil {
		// Если Redis недоступен, сохранить в БД
		return hc.dbRepo.SetCache(&models.Cache{
			CacheKey:       key,
			CacheValue:     string(data),
			ExpirationTime: time.Now().Add(ttl),
		})
	}

	return nil
}

func (hc *HybridCache) GetCache(key string) (string, error) {
	// Проверка в Redis
	data, err := hc.redisClient.Get(context.Background(), key).Result()
	if err == redis.Nil {
		// Если данных нет в Redis, ищем в БД
		cache, dbErr := hc.dbRepo.GetCacheByKey(key)
		if dbErr != nil || cache == nil {
			return "", nil
		}
		return cache.CacheValue, nil
	}
	return data, err
}
