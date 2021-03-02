package pubsub

import (
	"log"
	"shopee-crawler/config"
	"shopee-crawler/dto"

	"github.com/Shopify/sarama"
)

var publisherInstance *publisher

// Publisher ...
type Publisher interface {
	Publish(objType, topic string, data interface{}) error
}

type publisher struct {
	producer sarama.SyncProducer
}

// GetPublisher ...
func GetPublisher() Publisher {
	if publisherInstance == nil {
		publisherInstance = &publisher{
			config.GetSyncProducer(),
		}
	}
	return publisherInstance
}

// Publish ...
func (p *publisher) Publish(objType, topic string, data interface{}) error {
	msg := dto.BuildMessageDTO(objType, topic, data)
	partition, offset, err := p.producer.SendMessage(&msg.Msg)
	if err != nil {
		log.Printf("FAILED to send message: %s\n", err)
	} else {
		log.Printf("> message sent to partition %d at offset %d\n", partition, offset)
	}
	return err
}
