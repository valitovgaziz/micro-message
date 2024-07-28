package main

import (
	"log"
	"time"

	"encoding/json"

	"github.com/IBM/sarama"
	"gorm.io/gorm"
)

func main() {
	producer, err := sarama.NewSyncProducer([]string{"kafka:9092"}, nil)
	if err != nil {
		log.Fatalf("Failed to create producer: %v", err)
	}
	defer producer.Close()

	consumer, err := sarama.NewConsumer([]string{"kafka:9092"}, nil)
	if err != nil {
		log.Fatalf("Failed to create consumer: %v", err)
	}
	defer consumer.Close()

	partConsumer, err := consumer.ConsumePartition("testToPerf", 0, sarama.OffsetOldest)
	if err != nil {
		log.Fatalf("Failed to create partition consumer: %v", err)
	}
	defer partConsumer.Close()

	for {
		select {
		case msg, ok := <-partConsumer.Messages():
			if !ok {
				log.Println("Channel closed")
				return
			}

			var reseivedMessage Message
			err := json.Unmarshal(msg.Value, &reseivedMessage)
			if err != nil {
				log.Fatalf("Failed to unmarshal message: %v\n", err)
				continue
			}

			log.Printf("Received message: %v\n", &reseivedMessage)

			PerformMessage(&reseivedMessage)

			perfJsonMessage, err := json.Marshal(reseivedMessage)
			if err != nil {
				log.Fatalf("Failed to marshal message: %v\n", err)
				continue
			}

			resp := &sarama.ProducerMessage{
				Topic: "testFromPerf",
				Key:   sarama.StringEncoder(reseivedMessage.Id),
				Value: sarama.ByteEncoder(perfJsonMessage),
			}

			_, _, err = producer.SendMessage(resp)
			if err != nil {
				log.Fatalf("Failed to send message to Kafka: %v\n", err)
			}
		}
	}
}

type Message struct {
	Id          uint64         `json:"id"`           // message id
	UserId      uint64         `json:"user_id"`      // user id or sender id
	Destination string         `json:"destination"`  // message destination
	Content     string         `json:"content"`      // message content
	CreatedAt   time.Time      `json:"created_at"`   // message created time
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`   // message deleted time
	IsPerformed bool           `json:"is_performed"` // if message has been performed
	PerformedAt time.Time      `json:"performed_at"` // message performed time
}

func PerformMessage(message *Message) {
	// perform message
	message.IsPerformed = true
	message.PerformedAt = time.Now()

	// update message
	// TODO: update message in database

	// send message to destination
	// TODO: send message to destination

	// send message to kafka
	// TODO: send message to kafka

	// send message to user
	// TODO: send message to user

	// send message to admin
	// TODO: send message to admin

	// send message to log
	// TODO: send message to log

	// send message to history
	// TODO: send message to history
}
