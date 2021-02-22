package main

import (
	"fmt"
	httpclient "shopee-crawler/http-client"
)

func main() {
	client := httpclient.NewShopeeClient()
	body, err := client.Get("http://shopee.vn/api/v2/item/get?itemid=5610905722&shopid=65589552")
	fmt.Println(string(body), err)
}
