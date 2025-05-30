package kafka

import (
	"fmt"

	"github.com/IBM/sarama"
)

type Producer struct {
	producer sarama.SyncProducer 
}

func NewProducer(adresses []string) (*Producer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true


	syncProducer, err := sarama.NewSyncProducer(adresses, config)  
	if err != nil {
		return nil, fmt.Errorf("error with NewProducer: %w", err)
	}
	
	return &Producer{
		producer: syncProducer,
	}, nil
}

func (p *Producer) SendMessage(topic, message string) error {
	msg := &sarama.ProducerMessage{
		Topic: 	topic,
		Key:	nil,
		Value:	sarama.StringEncoder(message),
	}

	_, _, err := p.producer.SendMessage(msg)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}

	return nil
}

func (p *Producer) Close() {
	p.producer.Close()
}