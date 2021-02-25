package entity

// Category ...
type Category struct {
	ID    int    `json:"category_id" gorm:"column:id;primarykey"`
	Name  string `json:"name" gorm:"type:varchar(255)"`
	Shops []Shop `json:"shops" gorm:"many2many:shop_category"`
}

// TableName ...
func (c Category) TableName() string {
	return "category"
}
