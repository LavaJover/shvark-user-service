package kafka

import (
	"context"
	"log"
	"os"
	"os/signal"

	"gopkg.in/Shopify/sarama.v1"
)

type ConsumerHandler func(message *sarama.ConsumerMessage)

type KafkaConsumer struct {
	Group sarama.ConsumerGroup
}

func NewKafkaConsumer(brokers []string, groupID string) (*KafkaConsumer, error) {
	config := sarama.NewConfig()
	config.Version = sarama.V2_1_0_0
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	group, err := sarama.NewConsumerGroup(brokers, groupID, config)
	if err != nil {
		return nil, err
	}

	return &KafkaConsumer{Group: group}, nil
}

type consumerGroupHandler struct {
	handler ConsumerHandler
}

func (c *consumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (c *consumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }

func (c *consumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		c.handler(message)
		session.MarkMessage(message, "")
	}
	return nil
}

func (kc *KafkaConsumer) StartConsuming(topics []string, handler ConsumerHandler) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	consumer := &consumerGroupHandler{handler: handler}

	go func() {
		sigchan := make(chan os.Signal, 1)
		signal.Notify(sigchan, os.Interrupt)
		<-sigchan
		log.Println("Got signal after stop")
		cancel()
	}()

	for {
		if err := kc.Group.Consume(ctx, topics, consumer); err != nil {
			log.Printf("Failed to read: %v", err)
			return err
		}

		if ctx.Err() != nil {
			return ctx.Err()
		}
	}
}