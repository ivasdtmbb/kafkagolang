package main

import (
	"fmt"

	k "kafkagolang/internal/kafka"

	"github.com/sirupsen/logrus"
)

const (
	topic = "my-topic"
)

var addresses = []string{"localhost:9095", "localhost:9096", "localhost:9097"}

func main() {
	p, err := k.NewProducer(addresses)
	if err != nil {
		logrus.Fatal(err)
	}

	for i := 0; i < 100; i++ {
		msg := fmt.Sprintf("Kafka message %d", i)
		if err := p.SendMessage(topic, msg); err != nil {
			logrus.Error(err)
		}
	}
}