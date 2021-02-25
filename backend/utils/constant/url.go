package constant

var (
	// ListAllShopeeCategoryURL url to list all shopee category
	ListAllShopeeCategoryURL = "https://shopee.vn/api/v4/official_shop/get_categories?tab_type=0"

	// ListAllShopByCategoryURL url to list all shop has category
	// request param : category_id
	// example : https://shopee.vn/api/v4/official_shop/get_shops?category_id=78
	ListAllShopByCategoryURL = "https://shopee.vn/api/v4/official_shop/get_shops?category_id=%d"

	// ListAllItemByShopURL url to list all item by shop id
	// param: match_id : shop id,
	// 			order : desc
	// 			page_type : shop
	//			limit : 100
	//			newest: offset
	// exmalple : https://shopee.vn/api/v2/search_items/?limit=100&match_id=93936823&order=desc&page_type=shop&newest=30
	ListAllItemByShopURL = "https://shopee.vn/api/v2/search_items/?limit=100&match_id=%d&order=desc&page_type=shop&newest=%d"
)
