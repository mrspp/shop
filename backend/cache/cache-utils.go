package cache

import (
	"context"
	"fmt"
	"log"
	"shopee-crawler/cache/loader"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	ctx = context.Background()
)

const (
	serviceName = "crawler"
)

// Cache ...
type Cache interface {
	Get(key interface{}) (string, error)
	Evict(key interface{})
	Set(key, value string, ttl time.Duration)
}

// BaseCache ...
type baseCache struct {
	cacheManager *redis.Client
	service      string
	cacheName    string
	loader.CacheLoader
}

func (c baseCache) getCacheKey(key string) string {
	return fmt.Sprintf("%s.%s.%s", c.service, c.cacheName, key)
}

func (c baseCache) Get(k interface{}) (string, error) {
	key := fmt.Sprintf("%v", k)
	cacheKey := c.getCacheKey(key)
	result, err := c.cacheManager.Get(ctx, cacheKey).Result()
	switch {
	case err == redis.Nil:
		result, err := c.Load(key)
		if err != nil {
			return "", err
		}
		c.cacheManager.Set(ctx, cacheKey, result, time.Hour)
		return result, err
	case err != nil:
		return "", err
	default:
		return result, nil
	}
}

func (c baseCache) Evict(k interface{}) {
	key := fmt.Sprintf("%v", k)
	cacheKey := c.getCacheKey(key)
	if err := c.cacheManager.Del(ctx, cacheKey).Err(); err != nil {
		log.Println(err.Error())
	}
}

func (c baseCache) Set(key, value string, ttl time.Duration) {
	cacheKey := c.getCacheKey(key)
	c.cacheManager.Set(ctx, cacheKey, value, ttl)
}
