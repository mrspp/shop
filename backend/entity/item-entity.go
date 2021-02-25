package entity

// Item ...
type Item struct {
	ID                     int64      `json:"itemid"`
	PriceMaxBeforeDiscount int64      `json:"price_max_before_discount"`
	Image                  string     `json:"image"`
	ShopID                 int        `json:"shopid"`
	ReferenceItemID        string     `json:"reference_item_id"`
	Currency               string     `json:"currency"`
	RawDiscount            int        `json:"raw_discount"`
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

// ItemRating ...
type ItemRating struct {
	ID                int     `json:"ID" gorm:"primarykey;column:ID"`
	RatingStar        float64 `json:"rating_star"`
	RatingCount       int     `json:"rating_count"` // total star -> 1-5 star
	OneStarRate       int     `json:"one_star"`     // total star -> 1-5 star
	TwoStarRate       int     `json:"two_star"`     // total star -> 1-5 star
	ThreeStarRate     int     `json:"three_star"`   // total star -> 1-5 star
	FourStarRate      int     `json:"four_star"`    // total star -> 1-5 star
	FiveStarRate      int     `json:"five_star"`    // total star -> 1-5 star
	RcountWithImage   int     `json:"rcount_with_image"`
	RcountWithContext int     `json:"rcount_with_context"`
	ItemID            int
}

// TableName of item
func (i Item) TableName() string {
	return "item"
}

// TableName of item rating
func (ir ItemRating) TableName() string {
	return "item_rating"
}
