package loader

import (
	"encoding/json"
	"shopee-crawler/mapper"
	"shopee-crawler/repository"
	"strconv"
)

var shopLoaderInstance *shopLoader

type shopLoader struct {
	shopRepo repository.ShopRepo
}

func (c *shopLoader) Load(key string) (string, error) {
	shopID, err := strconv.Atoi(key)
	if err != nil {
		return "", err
	}

	shop, err := c.shopRepo.FindByIDWithItems(shopID)
	if err != nil {
		return "", err
	}
	shopCacheDTO := mapper.ShopEntityToCacheDTO(shop)
	data, err := json.Marshal(shopCacheDTO)
	return string(data), err
}

// GetShopLoader ...
func GetShopLoader() CacheLoader {
	if shopLoaderInstance == nil {
		shopLoaderInstance = &shopLoader{
			repository.GetShopRepo(),
		}
	}
	return shopLoaderInstance
}
