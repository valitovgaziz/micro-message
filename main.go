package main

import (
	"github.com/valitovgaziz/micro-message/initializers"
	"github.com/valitovgaziz/micro-message/models"
	"github.com/valitovgaziz/micro-message/storage"
)

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	initializers.ConnectDB(&config)
	storage.DB.AutoMigrate(&models.Message{})
	initializers.InitServerAndRouting(&config)
}
