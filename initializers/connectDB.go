package initializers

import (
	"fmt"
	"log"

	"github.com/valitovgaziz/micro-message/storage"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// create a new connection pool with the Postgres database.
func ConnectDB(config *Config) {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", config.DBHost, config.DBUserName, config.DBUserPassword, config.DBName, config.DBPort)

	storage.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}
	fmt.Println("? Connected Successfully to the Database")
}

