package main

import (
	"github.com/valitovgaziz/micro-message/initializers"
)

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	initializers.InitRouting(&config)
	initializers.ConnectDB(&config)
}
