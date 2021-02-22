package models

// Product ...
type Product struct {
	Item     Item        `json:"item"`
	Version  string      `json:"version"`
	Data     interface{} `json:"data"`
	ErrorMsg interface{} `json:"error_msg"`
	Error    interface{} `json:"error"`
}

// TierVariation ...
type TierVariation struct {
	Images     []string      `json:"images"`
	Properties []interface{} `json:"properties"`
	Type       int           `json:"type"`
	Name       string        `json:"name"`
	Options    []string      `json:"options"`
}

// Extinfo ...
type Extinfo struct {
	SellerPromotionLimit       int         `json:"seller_promotion_limit"`
	HasShopeePromo             bool        `json:"has_shopee_promo"`
	GroupBuyInfo               interface{} `json:"group_buy_info"`
	HolidayModeOldStock        interface{} `json:"holiday_mode_old_stock"`
	TierIndex                  []int       `json:"tier_index"`
	SellerPromotionRefreshTime int         `json:"seller_promotion_refresh_time"`
}

// PriceStock ...
type PriceStock struct {
	ModelID                  int64         `json:"model_id"`
	StockoutTime             interface{}   `json:"stockout_time"`
	Region                   string        `json:"region"`
	Rebate                   interface{}   `json:"rebate"`
	Price                    int64         `json:"price"`
	PromotionType            int           `json:"promotion_type"`
	AllocatedStock           interface{}   `json:"allocated_stock"`
	ShopID                   int           `json:"shop_id"`
	EndTime                  int           `json:"end_time"`
	StockBreakdownByLocation []interface{} `json:"stock_breakdown_by_location"`
	ItemID                   int64         `json:"item_id"`
	PromotionID              int           `json:"promotion_id"`
	PurchaseLimit            interface{}   `json:"purchase_limit"`
	StartTime                int           `json:"start_time"`
	Stock                    interface{}   `json:"stock"`
}

// Model ...
type Model struct {
	Itemid                          int64        `json:"itemid"`
	Status                          int          `json:"status"`
	CurrentPromotionReservedStock   int          `json:"current_promotion_reserved_stock"`
	Name                            string       `json:"name"`
	Promotionid                     int          `json:"promotionid"`
	Price                           int64        `json:"price"`
	PriceStocks                     []PriceStock `json:"price_stocks"`
	CurrentPromotionHasReserveStock bool         `json:"current_promotion_has_reserve_stock"`
	Currency                        string       `json:"currency"`
	NormalStock                     int          `json:"normal_stock"`
	Extinfo                         Extinfo      `json:"extinfo"`
	PriceBeforeDiscount             int64        `json:"price_before_discount"`
	Modelid                         int64        `json:"modelid"`
	Sold                            int          `json:"sold"`
	Stock                           int          `json:"stock"`
}

// CoinInfo ...
type CoinInfo struct {
	SpendCashUnit int           `json:"spend_cash_unit"`
	CoinEarnItems []interface{} `json:"coin_earn_items"`
}

// Attribute ...
type Attribute struct {
	IsPendingQc bool   `json:"is_pending_qc"`
	Idx         int    `json:"idx"`
	Value       string `json:"value"`
	ID          int    `json:"id"`
	IsTimestamp bool   `json:"is_timestamp"`
	Name        string `json:"name"`
}

// ItemRating ...
type ItemRating struct {
	RatingStar        float64 `json:"rating_star"`
	RatingCount       []int   `json:"rating_count"`
	RcountWithImage   int     `json:"rcount_with_image"`
	RcountWithContext int     `json:"rcount_with_context"`
}

// CategoryItem ...
type CategoryItem struct {
	DisplayName        string      `json:"display_name"`
	Catid              int         `json:"catid"`
	Image              interface{} `json:"image"`
	NoSub              bool        `json:"no_sub"`
	IsDefaultSubcat    bool        `json:"is_default_subcat"`
	BlockBuyerPlatform interface{} `json:"block_buyer_platform"`
}

// Item ...
type Item struct {
	Itemid                                int64           `json:"itemid" gorm:"primarykey"`
	PriceMaxBeforeDiscount                int64           `json:"price_max_before_discount"`
	ItemStatus                            string          `json:"item_status"`
	CanUseWholesale                       bool            `json:"can_use_wholesale"`
	ShowFreeShipping                      bool            `json:"show_free_shipping"`
	EstimatedDays                         int             `json:"estimated_days"`
	IsHotSales                            bool            `json:"is_hot_sales"`
	IsSlashPriceItem                      bool            `json:"is_slash_price_item"`
	UpcomingFlashSale                     interface{}     `json:"upcoming_flash_sale"`
	SlashLowestPrice                      interface{}     `json:"slash_lowest_price"`
	IsPartialFulfilled                    bool            `json:"is_partial_fulfilled"`
	Condition                             int             `json:"condition"`
	ShowOriginalGuarantee                 bool            `json:"show_original_guarantee"`
	AddOnDealInfo                         interface{}     `json:"add_on_deal_info"`
	IsNonCcInstallmentPaymentEligible     bool            `json:"is_non_cc_installment_payment_eligible"`
	Categories                            []CategoryItem  `json:"categories"`
	Ctime                                 int             `json:"ctime"`
	Name                                  string          `json:"name"`
	ShowShopeeVerifiedLabel               bool            `json:"show_shopee_verified_label"`
	SizeChart                             interface{}     `json:"size_chart"`
	IsPreOrder                            bool            `json:"is_pre_order"`
	ServiceByShopeeFlag                   interface{}     `json:"service_by_shopee_flag"`
	HistoricalSold                        int             `json:"historical_sold"`
	ReferenceItemID                       string          `json:"reference_item_id"`
	RecommendationInfo                    interface{}     `json:"recommendation_info"`
	BundleDealInfo                        interface{}     `json:"bundle_deal_info"`
	PriceMax                              int64           `json:"price_max"`
	HasLowestPriceGuarantee               bool            `json:"has_lowest_price_guarantee"`
	ShippingIconType                      int             `json:"shipping_icon_type"`
	Images                                []string        `json:"images"`
	PriceBeforeDiscount                   int             `json:"price_before_discount"`
	CodFlag                               int             `json:"cod_flag"`
	Catid                                 int             `json:"catid"`
	IsOfficialShop                        bool            `json:"is_official_shop"`
	CoinEarnLabel                         interface{}     `json:"coin_earn_label"`
	HashtagList                           interface{}     `json:"hashtag_list"`
	Sold                                  int             `json:"sold"`
	Makeup                                interface{}     `json:"makeup"`
	ItemRating                            ItemRating      `json:"item_rating"`
	ShowOfficialShopLabelInTitle          bool            `json:"show_official_shop_label_in_title"`
	Discount                              string          `json:"discount"`
	Reason                                interface{}     `json:"reason"`
	LabelIds                              []int           `json:"label_ids"`
	HasGroupBuyStock                      bool            `json:"has_group_buy_stock"`
	OtherStock                            int             `json:"other_stock"`
	DeepDiscount                          interface{}     `json:"deep_discount"`
	Attributes                            []Attribute     `json:"attributes"`
	BadgeIconType                         int             `json:"badge_icon_type"`
	Liked                                 bool            `json:"liked"`
	CmtCount                              int             `json:"cmt_count"`
	Image                                 string          `json:"image"`
	RecommendationAlgorithm               interface{}     `json:"recommendation_algorithm"`
	IsCcInstallmentPaymentEligible        bool            `json:"is_cc_installment_payment_eligible"`
	Shopid                                int             `json:"shopid"`
	NormalStock                           int             `json:"normal_stock"`
	VideoInfoList                         []interface{}   `json:"video_info_list"`
	InstallmentPlans                      interface{}     `json:"installment_plans"`
	ViewCount                             int             `json:"view_count"`
	VoucherInfo                           interface{}     `json:"voucher_info"`
	CurrentPromotionHasReserveStock       bool            `json:"current_promotion_has_reserve_stock"`
	LikedCount                            int             `json:"liked_count"`
	ShowOfficialShopLabel                 bool            `json:"show_official_shop_label"`
	PriceMinBeforeDiscount                int             `json:"price_min_before_discount"`
	ShowDiscount                          int             `json:"show_discount"`
	PreviewInfo                           interface{}     `json:"preview_info"`
	Flag                                  int             `json:"flag"`
	ExclusivePriceInfo                    interface{}     `json:"exclusive_price_info"`
	CurrentPromotionReservedStock         int             `json:"current_promotion_reserved_stock"`
	WholesaleTierList                     []interface{}   `json:"wholesale_tier_list"`
	GroupBuyInfo                          interface{}     `json:"group_buy_info"`
	ShopeeVerified                        bool            `json:"shopee_verified"`
	ItemHasPost                           bool            `json:"item_has_post"`
	HiddenPriceDisplay                    interface{}     `json:"hidden_price_display"`
	TransparentBackgroundImage            string          `json:"transparent_background_image"`
	WelcomePackageInfo                    interface{}     `json:"welcome_package_info"`
	DiscountStock                         int             `json:"discount_stock"`
	CoinInfo                              CoinInfo        `json:"coin_info"`
	IsAdult                               bool            `json:"is_adult"`
	Currency                              string          `json:"currency"`
	RawDiscount                           int             `json:"raw_discount"`
	IsPreferredPlusSeller                 bool            `json:"is_preferred_plus_seller"`
	IsCategoryFailed                      bool            `json:"is_category_failed"`
	PriceMin                              int             `json:"price_min"`
	CanUseBundleDeal                      bool            `json:"can_use_bundle_deal"`
	CbOption                              int             `json:"cb_option"`
	Brand                                 string          `json:"brand"`
	Stock                                 int             `json:"stock"`
	Status                                int             `json:"status"`
	BundleDealID                          int             `json:"bundle_deal_id"`
	IsGroupBuyItem                        interface{}     `json:"is_group_buy_item"`
	Description                           string          `json:"description"`
	FlashSale                             interface{}     `json:"flash_sale"`
	Models                                []Model         `json:"models"`
	HasLowFulfillmentRate                 bool            `json:"has_low_fulfillment_rate"`
	Price                                 int             `json:"price"`
	ShopLocation                          string          `json:"shop_location"`
	TierVariations                        []TierVariation `json:"tier_variations"`
	Makeups                               interface{}     `json:"makeups"`
	WelcomePackageType                    int             `json:"welcome_package_type"`
	ShowOfficialShopLabelInNormalPosition interface{}     `json:"show_official_shop_label_in_normal_position"`
	ItemType                              int             `json:"item_type"`
}
