package config

import (
	"log"
	"shopee-crawler/crawler"

	"google.golang.org/grpc"
)

var conn *grpc.ClientConn
var c crawler.CrawCategoryClient

func init() {
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	c = crawler.NewCrawCategoryClient(conn)
}
