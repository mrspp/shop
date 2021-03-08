package crawler

import "context"

type CrawlerServer struct {
}

// CrawCategory ...
func (c *CrawlerServer) CrawCategory(ctx context.Context, blank *Blank) (*SuccessResponse, error) {
	return &SuccessResponse{Success: true}, nil
}
