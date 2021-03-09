package main

import (
	"log"
	"net"
	"shopee-crawler/crawler"
	"shopee-crawler/proto"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	// utils.Mirgrate()
	godotenv.Load()
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
	crawlServer := crawler.CrawlerServer{}
	proto.RegisterCrawCategoryServer(grpcServer, &crawlServer)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
