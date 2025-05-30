package kafka

import (
	"log"

	"gopkg.in/Shopify/sarama.v1"
)

type KafkaProducer struct {
	SyncProducer sarama.SyncProducer
}

func NewKafkaProducer(brokers []string) (*KafkaProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}

	return &KafkaProducer{SyncProducer: producer}, nil
}

func (p *KafkaProducer) SendMessage(topic, message string) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	partition, offset, err := p.SyncProducer.SendMessage(msg)
	if err != nil {
		return err
	}

	log.Printf("Сообщение отправлено: topic=%s partition=%d offset=%d", topic, partition, offset)
	return nil
}

func (p *KafkaProducer) Close() error {
	return p.SyncProducer.Close()
}