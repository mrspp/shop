package mapper

import (
	"shopee-crawler/dto"
	"shopee-crawler/entity"
)

// ShopDTOToShopEntity ...
func ShopDTOToShopEntity(shopDTO dto.ShopDTO) entity.Shop {
	return entity.Shop{
		Userid:           shopDTO.Userid,
		Username:         shopDTO.Username,
		ID:               shopDTO.Shopid,
		ShopName:         shopDTO.ShopName,
		Logo:             shopDTO.Logo,
		LogoPc:           shopDTO.LogoPc,
		ShopCollectionID: shopDTO.ShopCollectionID,
		Ctime:            shopDTO.Ctime,
		BrandLabel:       shopDTO.BrandLabel,
	}
}

// ShopDTOsToShopEntities ...
func ShopDTOsToShopEntities(shopDTOs []dto.ShopDTO) []entity.Shop {
	shopEntities := make([]entity.Shop, 0)
	for _, s := range shopDTOs {
		shopEntities = append(shopEntities, ShopDTOToShopEntity(s))
	}
	return shopEntities
}
