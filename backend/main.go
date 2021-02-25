package main

import (
	"shopee-crawler/crawler"
	"shopee-crawler/utils"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	utils.Mirgrate()
	crawler.CrawlChain()
}
