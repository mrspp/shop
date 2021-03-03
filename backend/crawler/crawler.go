package crawler

// CrawlChain ...
func CrawlChain() {
	GetShopCrawler().Crawl()
	GetItemCrawler().Crawl()
}

// Crawler ...
type Crawler interface {
	Crawl() error
}
