package dto

// ShopDTO ...
type ShopDTO struct {
	Username         string `json:"username"`
	Name             string `json:"brand_name"`
	Shopid           int    `json:"shopid"`
	Logo             string `json:"logo"`
	LogoPc           string `json:"logo_pc"`
	ShopCollectionID int    `json:"shop_collection_id"`
	Ctime            int    `json:"ctime"`
	BrandLabel       int    `json:"brand_label"`
}

// ShopIndex ...
type ShopIndex struct {
	Index string    `json:"index"`
	Total int       `json:"total"`
	Shops []ShopDTO `json:"brand_ids"`
}

// ShopData ...
type ShopData struct {
	ShopIndexes []ShopIndex `json:"brands"`
	Total       int         `json:"brand_count"`
}

// ShopResponse ...
type ShopResponse struct {
	Data     ShopData `json:"data"`
	Error    int      `json:"error"`
	ErrorMsg string   `json:"error_msg"`
}
