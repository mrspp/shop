package dto

// ShopDTO ...
type ShopDTO struct {
	Userid           int    `json:"userid"`
	Username         string `json:"username"`
	Shopid           int    `json:"shopid"`
	ShopName         string `json:"shop_name"`
	Logo             string `json:"logo"`
	LogoPc           string `json:"logo_pc"`
	ShopCollectionID int    `json:"shop_collection_id"`
	Ctime            int    `json:"ctime"`
	BrandLabel       int    `json:"brand_label"`
}

// ShopData ...
type ShopData struct {
	Total         int       `json:"total"`
	OfficialShops []ShopDTO `json:"official_shops"`
}

// ShopResponse ...
type ShopResponse struct {
	Data     ShopData `json:"data"`
	Error    int      `json:"error"`
	ErrorMsg string   `json:"error_msg"`
}
