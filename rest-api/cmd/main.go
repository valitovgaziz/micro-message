package main

import (
	"github.com/valitovgaziz/micro-message/initializers"
	"github.com/valitovgaziz/micro-message/storage"
)

func main() {
	initializers.ConnectDB()
	storage.AutoMigrate()
	initializers.InitServerAndRouting()
}
