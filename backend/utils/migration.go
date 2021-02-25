package utils

import (
	"shopee-crawler/config"
	"shopee-crawler/entity"
)

// Mirgrate ...
func Mirgrate() {
	config.GetDbConnection().AutoMigrate(&entity.Category{}, &entity.Shop{}, &entity.Item{}, &entity.ItemRating{})
}
