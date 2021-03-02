package config

import (
	"log"
	"os"

	"github.com/Shopify/sarama"
)

var producer sarama.SyncProducer

// GetSyncProducer ...
func GetSyncProducer() sarama.SyncProducer {
	if producer == nil {
		broker := os.Getenv("KAFKA_BROKER")
		var err error
		producer, err = sarama.NewSyncProducer([]string{broker}, nil)
		if err != nil {
			log.Fatal("producer err: ", err)
		}
	}
	return producer
}
