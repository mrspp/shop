package crawler

// CrawlChain ...
func CrawlChain() {
	GetCategoryCrawler().Crawl()
	GetShopCrawler().Crawl()
	GetItemCrawler().Crawl()
}

// Crawler ...
type Crawler interface {
	Crawl() error
}
