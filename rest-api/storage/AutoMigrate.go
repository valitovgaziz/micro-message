package storage

import "github.com/valitovgaziz/rest-api/models"

func AutoMigrate() {
	DB.AutoMigrate(
		&models.Message{},
		&models.Statistics{},
	)
}