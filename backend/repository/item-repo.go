package repository

import (
	"shopee-crawler/config"
	"shopee-crawler/entity"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var itemRepoInstance *itemRepo

// ItemRepo ...
type ItemRepo interface {
	Save(item entity.Item) error
	SaveAll(items []entity.Item) error
}
type itemRepo struct {
	db *gorm.DB
}

func (i *itemRepo) Save(item entity.Item) error {
	return i.db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&item).Error
}
func (i *itemRepo) SaveAll(items []entity.Item) error {
	return i.db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&items).Error
}

// GetItemRepo ...
func GetItemRepo() ItemRepo {
	return &itemRepo{
		config.GetDbConnection(),
	}
}
