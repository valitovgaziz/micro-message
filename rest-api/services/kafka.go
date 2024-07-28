package services

import (
	"encoding/json"
	"errors"
	"log"
	"sync"
	"time"

	"github.com/IBM/sarama"
	"github.com/google/uuid"
	"github.com/valitovgaziz/rest-api/models"
)

var ResponseChannels map[string]chan *sarama.ConsumerMessage
var Mu sync.Mutex
var Producer sarama.SyncProducer
var Consumer sarama.Consumer
var PartConsumer sarama.PartitionConsumer
var mu sync.Mutex

func InitKafka() {
	ResponseChannels = make(map[string]chan *sarama.ConsumerMessage)
	Producer, err := sarama.NewSyncProducer([]string{"kafka:9092}"}, nil)
	if err != nil {
		log.Fatalf("Failed create producer: %v", err)
	}
	log.Printf("Producer created: %v", Producer)

	Consumer, err := sarama.NewConsumer([]string{"kafka:9092"}, nil)
	if err != nil {
		log.Fatalf("Failed create consumer: %v", err)
	}
	log.Printf("Consumer created: %v", Consumer)

	partConsumer, err := Consumer.ConsumePartition("test", 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("Failed create partition consumer: %v", err)
	}
	log.Printf("Partition consumer created: %v", partConsumer)

	go func() {
		for {
			select {
			case msg, ok := <-partConsumer.Messages():
				if !ok {
					log.Println("Channel closed, exiting goroutine")
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

func SendKafkaTestTopic(message *models.Message) error {
	requestID := uuid.New().String()
	jsonMessage, err := json.Marshal(message)
	if err!= nil {
		log.Printf("Failed to marshal message: %v", err)
		return err
	}
	msg := &sarama.ProducerMessage{
		Topic: "test",
		Key:   sarama.StringEncoder(requestID),
		Value: sarama.ByteEncoder(jsonMessage),
	}
	
	_, _, err = Producer.SendMessage(msg)
	if err!= nil {
		log.Printf("Failed to send message to Kafka: %v", err)
		return err
	}

	responseCh := make(chan *sarama.ConsumerMessage)
	mu.Lock()
	ResponseChannels[requestID] = responseCh
	mu.Unlock()

	select {
	case responseMsg := <- responseCh:
		log.Printf("Received response: %s", string(responseMsg.Value))
		return nil
	case <-time.After(10 * time.Second):
		mu.Lock()
		delete(ResponseChannels, requestID)
		mu.Unlock()
		return errors.New("Timeout")
	}
}

