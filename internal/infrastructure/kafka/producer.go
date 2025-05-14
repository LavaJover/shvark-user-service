package kafka

import (
	"context"
	"encoding/json"

	"github.com/segmentio/kafka-go"
)

type EventProducer interface {
	Produce(ctx context.Context, topic string, message interface{}) error
}

type KafkaProducer struct {
	writer *kafka.Writer
}

func NewKafkaProducer(brokerAddr string) *KafkaProducer {
	return &KafkaProducer{
		writer: &kafka.Writer{
			Addr: kafka.TCP(brokerAddr),
			Balancer: &kafka.LeastBytes{},
		},
	}
}

func (kp *KafkaProducer) Produce(ctx context.Context, topic string, message interface{}) error {
	bytes, err := json.Marshal(message)
	if err != nil {
		return err
	}

	return kp.writer.WriteMessages(ctx, kafka.Message{
		Topic: topic,
		Value: bytes,
	})
}