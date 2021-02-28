package dto

import "github.com/Shopify/sarama"

// MessageDTO ...
type MessageDTO *sarama.ProducerMessage

// SetValue ...
func (m *MessageDTO) SetValue(value []byte) {
	m.Value = sarama.ByteEncoder(value)
}

// AddObjectType ...
func (m *MessageDTO) AddObjectType(objType []byte) {
	if m.Headers == nil {
		m.Headers = []sarama.RecordHeader{
			{
				Key:   []byte("type"),
				Value: objType,
			},
		}
	}
	m.Headers = append(m.Headers, sarama.RecordHeader{
		Key:   []byte("type"),
		Value: objType,
	})
}
