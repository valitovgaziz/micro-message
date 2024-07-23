package storage

import "github.com/valitovgaziz/micro-message/models"

func AutoMigrate() {
	DB.AutoMigrate(
		&models.Message{},
		&models.Statistics{},
	)
}