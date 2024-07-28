package main

import (
	"github.com/valitovgaziz/rest-api/initializers"
	"github.com/valitovgaziz/rest-api/models"
	"github.com/valitovgaziz/rest-api/services"
	"github.com/valitovgaziz/rest-api/storage"
)

var Kafka  models.Kafka

func init() {

}

func main() {
	initializers.ConnectDB()
	storage.AutoMigrate()
	initializers.InitServerAndRouting()
	services.InitKafka()
}
