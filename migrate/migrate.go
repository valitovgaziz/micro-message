package main

import (
	"fmt"
	"log"

	"github.com/valitovgaziz/micro-auth/initializers"
	"github.com/valitovgaziz/micro-auth/models"
)

// load the environment variables and create a database connection pool to the Postgres server.
func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
	fmt.Println("? Migration complete")
}

