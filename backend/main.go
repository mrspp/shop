package main

import (
	"shopee-crawler/pubsub"
)

func main() {
	// godotenv.Load(".env")
	// utils.Mirgrate()
	// crawler.CrawlChain()
	pubsub.Publish("alo alo 123")
}
