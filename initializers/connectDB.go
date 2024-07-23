package initializers

import (
	"fmt"
	"log"
	"os"

	"github.com/valitovgaziz/micro-message/storage"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// create a new connection pool with the Postgres database.
func ConnectDB() {
	var err error
	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	fmt.Println(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	storage.DB, err = gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		},
	)
	if err != nil {
		log.Fatal("Failed to connect to the Database")
		os.Exit(2)
	}
	fmt.Println("? Connected Successfully to the Database")
	storage.DB.Logger = logger.Default.LogMode(logger.Info)
}
