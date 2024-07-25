package main

import (
	"sync"

	"github.com/IBM/sarama"
)

var responseChannels map[string]chan *sarama.ConsumerMessage
var mu sync.Mutex

func main() {
	
}
