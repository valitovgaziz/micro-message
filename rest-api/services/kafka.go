package services

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/IBM/sarama"
	"github.com/google/uuid"
	"github.com/valitovgaziz/rest-api/models"
)

var ResponseChannels map[string]chan *sarama.ConsumerMessage
var mu sync.Mutex

func InitKafka() {
	ResponseChannels = make(map[string]chan *sarama.ConsumerMessage)

	Producer, err := sarama.NewSyncProducer([]string{"kafka:9092"}, nil)
	if err != nil {
		log.Fatal("Failed to create producer: %v", err)
	}
	defer Producer.Close()

	Consumer, err := sarama.NewConsumer([]string{"kafka:9092"}, nil)
	if err != nil {
		log.Fatal("Failed to create consumer: %v", err)
	}
	defer Consumer.Close()

	PartConsumer, err := Consumer.ConsumePartition("performed", 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatal("Failed to create partition consumer: %v", err)
	}
	defer PartConsumer.Close()

	go func() {
		for {
			select {
			case msg, ok := <-PartConsumer.Messages():
				if !ok {
					log.Fatal("Channel closed, exiting goroutine")
					return
				}
				responseID := string(msg.Key)
				mu.Lock()
				ch, exists := ResponseChannels[responseID]
				if exists {
					ch <- msg
					delete(ResponseChannels, responseID)
				}
				mu.Unlock()
			}
		}
	}()
}


func SendKafkaMessage(message *models.Message) {

	requestID := uuid.New().String()

	bytes, err := json.Marshal(message)
	if err!= nil {
		log.Fatal("Error while marshalling message: %v", err)
		return
	}

	msg := &sarama.ProducerMessage{
		Topic: "perform",
		Key: sarama.StringEncoder(requestID),
		Value: sarama.ByteEncoder(bytes),
	}

	_, _, err = Producer.SendMessage(msg)

}