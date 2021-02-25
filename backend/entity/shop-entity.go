package entity

// Shop ...
type Shop struct {
	Userid           int        `json:"userid"`
	Username         string     `json:"username"`
	ID               int        `json:"shopid" gorm:"column:id;primarykey"`
	ShopName         string     `json:"shop_name"`
	Logo             string     `json:"logo"`
	LogoPc           string     `json:"logo_pc"`
	ShopCollectionID int        `json:"shop_collection_id"`
	Ctime            int        `json:"ctime"`
	BrandLabel       int        `json:"brand_label"`
	Categories       []Category `json:"categories" gorm:"many2many:shop_category"`
	Items            []Item
}

// TableName ...
func (s Shop) TableName() string {
	return "shop"
}
