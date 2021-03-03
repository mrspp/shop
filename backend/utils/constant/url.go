package constant

var (
	// ListAllCategoryURL url to list all category
	ListAllCategoryURL = "https://shopee.vn/api/v4/official_shop/get_categories?tab_type=0"

	// ListAllShopURL url to list all shop
	ListAllShopURL = "https://shopee.vn/api/v4/official_shop/get_shops_by_category?need_zhuyin=0&category_id=-1"

	// ListAllItemByShopURL url to list all item by shop id
	// param: match_id : shop id,
	// 			order : desc
	// 			page_type : shop
	//			limit : 100
	//			newest: offset
	// exmalple : https://shopee.vn/api/v2/search_items/?limit=100&match_id=93936823&order=desc&page_type=shop&newest=30
	ListAllItemByShopURL = "https://shopee.vn/api/v2/search_items/?limit=100&match_id=%d&order=desc&page_type=shop&newest=%d"
)
