package cache

import (
	"shopee-crawler/cache/loader"
	"shopee-crawler/config"
)

var shopCahceInstance *shopCache

type shopCache struct {
	baseCache
}

var cacheName = "shop"

// GetShopCache ...
func GetShopCache() Cache {
	if shopCahceInstance == nil {
		shopCahceInstance = &shopCache{
			baseCache{
				cacheManager: config.GetCacheManager(),
				service:      serviceName,
				cacheName:    cacheName,
				CacheLoader:  loader.GetShopLoader(),
			},
		}
	}
	return shopCahceInstance
}
