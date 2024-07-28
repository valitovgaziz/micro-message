package models

import (
	"sync"

	"github.com/IBM/sarama"
)

type Kafka struct {
	ResponseChannels map[string]chan *sarama.ConsumerMessage
	Mu sync.Mutex
	Producer *sarama.SyncProducer
	Consumer *sarama.Consumer
	PartConsumer *sarama.PartitionConsumer
}