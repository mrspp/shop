package mapper

import (
	"shopee-crawler/dto"
	"shopee-crawler/entity"
)

var priceUnit int64 = 100000

// ItemDTOToEntity ...
func ItemDTOToEntity(itemDTO dto.ItemDTO) entity.Item {
	return entity.Item{
		ID:                     itemDTO.Itemid,
		PriceMaxBeforeDiscount: itemDTO.PriceBeforeDiscount / priceUnit,
		Image:                  itemDTO.Image,
		ShopID:                 itemDTO.Shopid,
		ReferenceItemID:        itemDTO.ReferenceItemID,
		Currency:               itemDTO.Currency,
		RawDiscount:            itemDTO.RawDiscount,
		PriceBeforeDiscount:    itemDTO.PriceMaxBeforeDiscount / priceUnit,
		CmtCount:               itemDTO.CmtCount,
		ViewCount:              itemDTO.ViewCount,
		Catid:                  itemDTO.Catid,
		IsOfficialShop:         itemDTO.IsOfficialShop,
		Brand:                  itemDTO.Brand,
		PriceMin:               itemDTO.PriceMin / priceUnit,
		PriceMinBeforeDiscount: itemDTO.PriceMinBeforeDiscount / priceUnit,
		Sold:                   itemDTO.Sold,
		Stock:                  itemDTO.Stock,
		PriceMax:               itemDTO.PriceMax / priceUnit,
		Price:                  itemDTO.Price / priceUnit,
		ItemRating: entity.ItemRating{
			RatingStar:        itemDTO.ItemRating.RatingStar,
			RatingCount:       itemDTO.ItemRating.RatingCount[0],
			OneStarRate:       itemDTO.ItemRating.RatingCount[1],
			TwoStarRate:       itemDTO.ItemRating.RatingCount[2],
			ThreeStarRate:     itemDTO.ItemRating.RatingCount[3],
			FourStarRate:      itemDTO.ItemRating.RatingCount[4],
			FiveStarRate:      itemDTO.ItemRating.RatingCount[5],
			RcountWithImage:   itemDTO.ItemRating.RcountWithImage,
			RcountWithContext: itemDTO.ItemRating.RcountWithContext,
			ItemID:            int(itemDTO.Itemid),
		},
		Name:           itemDTO.Name,
		Ctime:          itemDTO.Ctime,
		ItemStatus:     itemDTO.ItemStatus,
		HistoricalSold: itemDTO.HistoricalSold,
	}
}

// ItemDTOsToEntities ...
func ItemDTOsToEntities(itemDTOs []dto.ItemDTO) []entity.Item {
	items := make([]entity.Item, 0)
	for _, i := range itemDTOs {
		if i.Price%priceUnit != 0 {
			continue
		}
		items = append(items, ItemDTOToEntity(i))
	}
	return items
}
