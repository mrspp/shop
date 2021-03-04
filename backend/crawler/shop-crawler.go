package crawler

import (
	"os"
	"shopee-crawler/pubsub"
	"shopee-crawler/utils/constant"
	httpclient "shopee-crawler/utils/http-client"
)

var shopCrwalerInstance *shopCrwaler

type shopCrwaler struct {
	publisher pubsub.Publisher
}

// ShopCrawler ...
type ShopCrawler interface {
	Crawl() error
}

// GetShopCrawler ...
func GetShopCrawler() ShopCrawler {
	if shopCrwalerInstance == nil {
		shopCrwalerInstance = &shopCrwaler{
			pubsub.GetPublisher(),
		}
	}
	return shopCrwalerInstance
}

// Crawl ...
func (s *shopCrwaler) Crawl() error {
	client := httpclient.NewShopeeClient()
	resp, err := client.Get(constant.ListAllShopURL)
	if err != nil {
		return err
	}

	return s.publisher.Publish(constant.SHOP_TYPE, os.Getenv("KAFKA_TOPIC"), resp)
}
