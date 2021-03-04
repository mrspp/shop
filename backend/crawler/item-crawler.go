package crawler

import (
	"encoding/json"
	"fmt"
	"os"
	"shopee-crawler/dto"
	"shopee-crawler/pubsub"
	"shopee-crawler/utils"
	"shopee-crawler/utils/constant"
	httpclient "shopee-crawler/utils/http-client"
)

var itemCrawlerInstance *itemCrawler

// ItemCrawler ...
type ItemCrawler interface {
	CrawlByShop(url string, shopID int) error
}
type itemCrawler struct {
	publisher pubsub.Publisher
}

// GetItemCrawler ...
func GetItemCrawler() ItemCrawler {
	if itemCrawlerInstance == nil {
		itemCrawlerInstance = &itemCrawler{
			pubsub.GetPublisher(),
		}
	}
	return itemCrawlerInstance
}

func (i *itemCrawler) CrawlByShop(url string, shopID int) error {
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
		i.publisher.Publish(constant.ITEM_TYPE, os.Getenv("KAFKA_TOPIC"), resp)
		stopCrawl = itemRS.Nomore
		offset += step
		utils.RandomSleep()
	}
	return nil
}
