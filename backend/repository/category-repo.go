package repository

import (
	"shopee-crawler/config"
	"shopee-crawler/entity"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var categoryRepoInstance *categoryRepo

// CategoryRepo ...
type CategoryRepo interface {
	Save(category entity.Category) error
	SaveAll(categories []entity.Category) error
	FindByID(categoryID int) (*entity.Category, error)
	FindAll() ([]entity.Category, error)
}

type categoryRepo struct {
	db *gorm.DB
}

func (c *categoryRepo) Save(category entity.Category) error {
	return c.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"name"}),
	}).Create(&category).Error
}
func (c *categoryRepo) SaveAll(categories []entity.Category) error {
	return c.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"name"}),
	}).Create(&categories).Error
}

// FindByID ...
func (c *categoryRepo) FindByID(categoryID int) (*entity.Category, error) {
	var category entity.Category
	if err := c.db.Where("id = ?", categoryID).Find(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

// FindAll ...
func (c *categoryRepo) FindAll() ([]entity.Category, error) {
	var categories []entity.Category
	if err := c.db.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

// GetCategoryRepo ...
func GetCategoryRepo() CategoryRepo {
	if categoryRepoInstance == nil {
		categoryRepoInstance = &categoryRepo{
			config.GetDbConnection(),
		}
	}
	return categoryRepoInstance
}
