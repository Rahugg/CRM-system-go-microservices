package kafka

import (
	"crm_system/config/auth"
	"fmt"
	"github.com/IBM/sarama"
)

type Producer struct {
	asyncProducer sarama.AsyncProducer
	topic         string
}

func NewProducer(cfg *auth.Configuration) (*Producer, error) {
	samaraConfig := sarama.NewConfig()

	asyncProducer, err := sarama.NewAsyncProducer(cfg.Kafka.Brokers, samaraConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to NewAsyncProducer err: %w", err)
	}

	return &Producer{
		asyncProducer: asyncProducer,
		topic:         cfg.Kafka.Producer.Topic,
	}, nil
}

func (p *Producer) ProduceMessage(message []byte) {
	msg := &sarama.ProducerMessage{
		Topic: p.topic,
		Value: sarama.ByteEncoder(message),
	}

	p.asyncProducer.Input() <- msg
}
