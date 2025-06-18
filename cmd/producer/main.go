package main 

import (
	k "kafkagolang/internal/kafka"
	"github.com/sirupsen/logrus"
	"github.com/google/uuid"
	
	"time"
	"fmt"
)

const (
	topic = "my-topic"
	numOfKeys = 20
)

var address = []string{"localhost:9095", "localhost:9096", "localhost:9097"}

 func main(){
	p, err := k.NewProducer(address)
	if err != nil {
		logrus.Fatal(err)
	}

	keys := generateUUIDString()
	for i := 0; i < 100; i++ {
		msg := fmt.Sprintf("Kafka message %d", i)
		key := keys[i % numOfKeys]
		if err := p.Produce(msg, topic, key, time.Now()); err != nil {
			logrus.Error(err)
		}
	}
 }

func generateUUIDString() [numOfKeys]string {
	var uuids[numOfKeys]string
	for i := 0; i < numOfKeys; i++ {
		uuids[i] = uuid.NewString()
	}
	return uuids
} 