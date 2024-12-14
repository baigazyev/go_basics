package cache

// type HybridCache struct {
// 	redisClient *redis.Client
// 	dbRepo      *repositories.CacheRepository
// }

// func NewHybridCache(redisClient *redis.Client, dbRepo *repositories.CacheRepository) *HybridCache {
// 	return &HybridCache{redisClient: redisClient, dbRepo: dbRepo}
// }

// func (hc *HybridCache) SetCache(key string, value interface{}, ttl time.Duration) error {
// 	data, err := json.Marshal(value)
// 	if err != nil {
// 		return err
// 	}

// 	// Сохранить в Redis
// 	err = hc.redisClient.Set(context.Background(), key, data, ttl).Err()
// 	if err != nil {
// 		// Если Redis недоступен, сохранить в БД
// 		return hc.dbRepo.SetCache(key, string(data), ttl)
// 	}

// 	return nil
// }

// func (hc *HybridCache) GetCache(key string) (string, error) {
// 	// Проверить Redis
// 	data, err := hc.redisClient.Get(context.Background(), key).Result()
// 	if err == redis.Nil {
// 		// Если данных нет в Redis, искать в БД
// 		cache, dbErr := hc.dbRepo.GetCache(key)
// 		if dbErr != nil || cache == nil {
// 			return "", nil // Кэш не найден
// 		}
// 		return cache.CacheValue, nil
// 	}
// 	return data, err
// }
