package dto

import (
	"encoding/json"
	"log"

	"github.com/Shopify/sarama"
)

// MessageDTO ...
type MessageDTO struct {
	Msg sarama.ProducerMessage
}

// SetValue ...
func (m *MessageDTO) SetValue(value interface{}) error {
	byteVal, err := json.Marshal(value)
	if err != nil {
		return err
	}
	b := sarama.ByteEncoder(byteVal)
	m.Msg.Value = b

	return nil
}

// SetObjectType ...
func (m *MessageDTO) SetObjectType(objType []byte) {
	if m.Msg.Headers == nil {
		m.Msg.Headers = []sarama.RecordHeader{
			{
				Key:   []byte("type"),
				Value: objType,
			},
		}
		return
	}
	m.Msg.Headers = append(m.Msg.Headers, sarama.RecordHeader{
		Key:   []byte("type"),
		Value: objType,
	})
}

// BuildMessageDTO ...
func BuildMessageDTO(header, topic string, data interface{}) MessageDTO {
	msg := MessageDTO{
		Msg: sarama.ProducerMessage{
			Topic: topic,
		},
	}
	err := msg.SetValue(data)
	if err != nil {
		log.Fatal(err)
	}
	msg.SetObjectType([]byte("category"))
	return msg
}
