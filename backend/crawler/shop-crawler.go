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

var shopCrwalerInstance *shopCrwaler

type shopCrwaler struct {
	categoryRepo repository.CategoryRepo
	shopRepo     repository.ShopRepo
}

// GetShopCrawler ...
func GetShopCrawler() Crawler {
	if shopCrwalerInstance == nil {
		shopCrwalerInstance = &shopCrwaler{
			repository.GetCategoryRepo(),
			repository.GetShopRepo(),
		}
	}
	return shopCrwalerInstance
}

// Crawl ...
func (s *shopCrwaler) Crawl() error {
	categories, err := s.categoryRepo.FindAll()
	if err != nil {
		return err
	}
	for _, c := range categories {
		if err := s.crawlShop(constant.ListAllShopByCategoryURL, c.ID); err != nil {
			return err
		}
		utils.RandomSleep()
	}
	return nil
}

func (s *shopCrwaler) crawlShop(url string, categoryID int) error {
	client := httpclient.NewShopeeClient()
	url = fmt.Sprintf(url, categoryID)
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	fmt.Println(string(resp))
	var shopRS dto.ShopResponse
	err = json.Unmarshal(resp, &shopRS)
	category, err := s.categoryRepo.FindByID(categoryID)
	if err != nil {
		return err
	}

	category.Shops = append(category.Shops, mapper.ShopDTOsToShopEntities(shopRS.Data.OfficialShops)...)
	s.categoryRepo.Save(*category)

	return nil
}
