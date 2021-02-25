package crawler

import (
	"encoding/json"
	"fmt"
	"shopee-crawler/dto"
	"shopee-crawler/mapper"
	"shopee-crawler/repository"
	"shopee-crawler/utils"
	"shopee-crawler/utils/constant"
	httpclient "shopee-crawler/utils/http-client"
)

var itemCrawlerInstance *itemCrawler

type itemCrawler struct {
	shopRepo repository.ShopRepo
	itemRepo repository.ItemRepo
}

// GetItemCrawler ...
func GetItemCrawler() Crawler {
	if itemCrawlerInstance == nil {
		itemCrawlerInstance = &itemCrawler{
			repository.GetShopRepo(),
			repository.GetItemRepo(),
		}
	}
	return itemCrawlerInstance
}
func (i *itemCrawler) Crawl() error {
	shops, err := i.shopRepo.FindAll()
	if err != nil {
		return err
	}
	for _, s := range shops {
		if err := i.crawlItemByShop(constant.ListAllItemByShopURL, s.ID); err != nil {
			return err
		}
		utils.RandomSleep()
	}
	return nil
}
func (i *itemCrawler) crawlItemByShop(url string, shopID int) error {
	offset := 0
	step := 100
	stopCrawl := false
	for !stopCrawl {
		crawlURL := fmt.Sprintf(url, shopID, offset)
		client := httpclient.NewShopeeClient()
		resp, err := client.Get(crawlURL)
		if err != nil {
			return err
		}
		var itemRS dto.ItemResponse
		err = json.Unmarshal(resp, &itemRS)
		if err != nil {
			return err
		}
		i.itemRepo.SaveAll(mapper.ItemDTOsToEntities(itemRS.Items))
		stopCrawl = itemRS.Nomore
		offset += step
		utils.RandomSleep()
	}
	return nil
}
