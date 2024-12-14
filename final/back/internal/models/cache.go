package models

import "time"

type Cache struct {
	CacheKey       string    `json:"cache_key" db:"cache_key"`
	CacheValue     string    `json:"cache_value" db:"cache_value"`
	ExpirationTime time.Time `json:"expiration_time" db:"expiration_time"`
}
