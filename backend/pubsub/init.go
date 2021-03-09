package pubsub

import (
	"log"
	"shopee-crawler/proto"

	"google.golang.org/grpc"
)

var conn *grpc.ClientConn
var c proto.CrawCategoryClient

func init() {
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	c = proto.NewCrawCategoryClient(conn)
}
