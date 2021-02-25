package crawler

import (
	"encoding/json"
	"log"
	"shopee-crawler/dto"
	"shopee-crawler/mapper"
	"shopee-crawler/repository"
	"shopee-crawler/utils/constant"
	httpclient "shopee-crawler/utils/http-client"
)

var categoryCrawlerInstance *categoryCrawler

type categoryCrawler struct {
	categoryRepo repository.CategoryRepo
}

// GetCategoryCrawler ...
func GetCategoryCrawler() Crawler {
	if categoryCrawlerInstance == nil {
		categoryCrawlerInstance = &categoryCrawler{
			repository.GetCategoryRepo(),
		}
	}
	return categoryCrawlerInstance
}

// Crawl ...
func (c *categoryCrawler) Crawl() error {
	var cate dto.CategoryResponse
	client := httpclient.NewShopeeClient()
	jsonRespone, err := client.Get(constant.ListAllShopeeCategoryURL)
	if err != nil {
		log.Println(err)
		return err
	}
	err = json.Unmarshal(jsonRespone, &cate)
	if err != nil {
		log.Println(err)
		return err
	}
	return c.categoryRepo.SaveAll(mapper.CategoryDTOsToEntities(cate.Data.Categories))
}
