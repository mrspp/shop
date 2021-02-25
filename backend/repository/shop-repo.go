package repository

import (
	"shopee-crawler/config"
	"shopee-crawler/entity"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var shopRepoInstance *shopRepo

// ShopRepo ...
type ShopRepo interface {
	Save(shop entity.Shop) error
	SaveAll(shops []entity.Shop) error
	FindAll() ([]entity.Shop, error)
}

type shopRepo struct {
	db *gorm.DB
}

func (s *shopRepo) Save(shop entity.Shop) error {
	return s.db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&shop).Error
}

func (s *shopRepo) SaveAll(shops []entity.Shop) error {
	for _, shop := range shops {
		if err := s.Save(shop); err != nil {
			return err
		}
	}
	return nil
}

func (s *shopRepo) FindAll() ([]entity.Shop, error) {
	var shops []entity.Shop
	if err := s.db.Find(&shops).Error; err != nil {
		return nil, err
	}
	return shops, nil
}

// GetShopRepo ...
func GetShopRepo() ShopRepo {
	if shopRepoInstance == nil {
		shopRepoInstance = &shopRepo{
			config.GetDbConnection(),
		}
	}
	return shopRepoInstance
}
