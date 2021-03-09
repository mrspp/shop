package crawler

import (
	"context"
	"shopee-crawler/proto"
)

type CrawlerServer struct {
}

// CrawCategory ...
func (c *CrawlerServer) CrawCategory(ctx context.Context, blank *proto.Blank) (*proto.SuccessResponse, error) {
	return &proto.SuccessResponse{Success: true}, nil
}
