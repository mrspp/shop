package crawler

import (
	"fmt"
	"os"
	"shopee-crawler/pubsub"
	"shopee-crawler/utils/constant"
	httpclient "shopee-crawler/utils/http-client"
)

var shopCrwalerInstance *shopCrwaler

type shopCrwaler struct {
	publisher pubsub.Publisher
}

// GetShopCrawler ...
func GetShopCrawler() Crawler {
	if shopCrwalerInstance == nil {
		shopCrwalerInstance = &shopCrwaler{
			pubsub.GetPublisher(),
		}
	}
	return shopCrwalerInstance
}

// Crawl ...
func (s *shopCrwaler) Crawl() error {
	// TODO: List all
	// categories, err := s.categoryRepo.FindAll()
	// if err != nil {
	// 	return err
	// }
	// for _, c := range categories {
	// 	if err := s.crawlShop(constant.ListAllShopByCategoryURL, c.ID); err != nil {
	// 		return err
	// 	}
	// 	utils.RandomSleep()
	// }
	return nil
}

func (s *shopCrwaler) crawlShop(url string, categoryID int) error {
	client := httpclient.NewShopeeClient()
	url = fmt.Sprintf(url, categoryID)
	resp, err := client.Get(url)
	if err != nil {
		return err
	}

	return s.publisher.Publish(constant.SHOP_TYPE, os.Getenv("KAFKA_TOPIC"), resp)

}
