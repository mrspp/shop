package main

import (
	"shopee-crawler/crawler"

	"github.com/joho/godotenv"
)

func main() {
	// utils.Mirgrate()
	godotenv.Load()
	crawler.CrawlChain()
}
