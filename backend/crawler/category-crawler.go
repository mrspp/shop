package crawler

import (
	"log"
	"os"
	"shopee-crawler/pubsub"
	"shopee-crawler/utils/constant"
	httpclient "shopee-crawler/utils/http-client"
)

var categoryCrawlerInstance *categoryCrawler

// CategoryCrawler ...
type CategoryCrawler interface {
	Crawl() error
}
type categoryCrawler struct {
	publisher pubsub.Publisher
}

// GetCategoryCrawler ...
func GetCategoryCrawler() CategoryCrawler {
	if categoryCrawlerInstance == nil {
		categoryCrawlerInstance = &categoryCrawler{
			pubsub.GetPublisher(),
		}
	}
	return categoryCrawlerInstance
}

// Crawl ...
func (c *categoryCrawler) Crawl() error {
	client := httpclient.NewShopeeClient()
	jsonRespone, err := client.Get(constant.ListAllCategoryURL)
	if err != nil {
		log.Println(err)
		return err
	}
	return c.publisher.Publish(constant.CATEGORY_TYPE, os.Getenv("KAFKA_TOPIC"), jsonRespone)
}
