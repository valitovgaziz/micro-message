package main

import (
	"github.com/valitovgaziz/micro-message/initializers"
	"github.com/valitovgaziz/micro-message/storage"
)

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	initializers.ConnectDB(&config)
	storage.AutoMigrate()
	initializers.InitServerAndRouting(&config)
}
