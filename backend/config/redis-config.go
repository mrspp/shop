package config

import (
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
)

var rdm *redis.Client

type redisConfig struct {
	addr     string
	password string
	db       int
}

func buildRedisConfig() redisConfig {
	return redisConfig{
		addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		password: os.Getenv("REDIS_PASSWORD"),
	}
}

// GetCacheManager ...
func GetCacheManager() *redis.Client {
	if rdm == nil {
		cacheConfig := buildRedisConfig()
		rdm = redis.NewClient(&redis.Options{
			Addr:     cacheConfig.addr,
			Password: cacheConfig.password,
			DB:       cacheConfig.db,
		})
	}
	return rdm
}
