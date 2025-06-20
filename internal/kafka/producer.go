
package kafka

import (
	// "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	// github.com/IBM/sarama
	// github.com/segmentio/kafka-go
	// github.com/lovoo/goka

	"errors"
	"fmt"
	"strings"
	"time"
)

const(
	flushTimeout = 500 // ms
)

var errUnknownType = errors.New("unknown event type")

type Producer struct {
	producer *kafka.Producer
}

func NewProducer(address []string) (*Producer, error){
	conf := &kafka.ConfigMap{
		"bootstrap.servers": strings.Join(address, ","),

	}
	p, err := kafka.NewProducer(conf)
	if err != nil {
		return nil, fmt.Errorf("error with new producer: %w", err)
	}
	
	return &Producer{producer: p}, nil
}

func (p *Producer) Produce(message, topic, key string, t time.Time) error{
	kafkaMsg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic: &topic,
			Partition: kafka.PartitionAny,
		},
		Value:			[]byte(message),
		Key:			[]byte(key),
		Timestamp: 		t,
	}
	
	kafkaChan := make(chan kafka.Event)
	if err := p.producer.Produce(kafkaMsg, kafkaChan); err != nil {
		return err
	}
	
	e := <- kafkaChan
	switch ev := e.(type) {
	case *kafka.Message:
		return nil
	case kafka.Error:
		return ev
	default:
		return errUnknownType
	}
}

func (p *Producer) Close(){
	p.producer.Flush(flushTimeout)
	p.producer.Close()
}
