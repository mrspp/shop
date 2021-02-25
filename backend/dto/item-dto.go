package dto

// ItemRating ...
type ItemRating struct {
	RatingStar        float64 `json:"rating_star"`
	RatingCount       []int   `json:"rating_count"` // total star -> 1-5 star
	RcountWithImage   int     `json:"rcount_with_image"`
	RcountWithContext int     `json:"rcount_with_context"`
}

// ItemDTO ...
type ItemDTO struct {
	Itemid                 int64      `json:"itemid"`
	PriceMaxBeforeDiscount int64      `json:"price_max_before_discount"`
	Image                  string     `json:"image"`
	Shopid                 int        `json:"shopid"`
	ReferenceItemID        string     `json:"reference_item_id"`
	Currency               string     `json:"currency"`
	RawDiscount            int        `json:"raw_discount"`
	Images                 []string   `json:"images"`
	PriceBeforeDiscount    int64      `json:"price_before_discount"`
	CmtCount               int        `json:"cmt_count"`
	ViewCount              int        `json:"view_count"`
	Catid                  int        `json:"catid"`
	IsOfficialShop         bool       `json:"is_official_shop"`
	Brand                  string     `json:"brand"`
	PriceMin               int64      `json:"price_min"`
	PriceMinBeforeDiscount int64      `json:"price_min_before_discount"`
	Sold                   int        `json:"sold"`
	Stock                  int        `json:"stock"`
	PriceMax               int64      `json:"price_max"`
	Price                  int64      `json:"price"`
	ItemRating             ItemRating `json:"item_rating"`
	Name                   string     `json:"name"`
	Ctime                  int        `json:"ctime"`
	ItemStatus             string     `json:"item_status"`
	HistoricalSold         int        `json:"historical_sold"`
}

// ItemResponse ...
type ItemResponse struct {
	TotalCount      int           `json:"total_count"`
	Nomore          bool          `json:"nomore"`
	Items           []ItemDTO     `json:"items"`
	ReservedKeyword string        `json:"reserved_keyword"`
	HintKeywords    []interface{} `json:"hint_keywords"`
}
