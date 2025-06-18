package kafka

import (

)

type Consumer struct {
	consumer *kafka.Consumer

}

func NewConsumer(addresses []string) (*Consumer, error) {
	return &Consumer{
		consumer: c,
	}, nil
}