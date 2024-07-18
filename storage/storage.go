package storage

import (
	"github.com/valitovgaziz/micro-message/models"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SaveMessage(message *models.Message) (*models.Message, error) {
	err := DB.Create(&message).Error
	if err!= nil {
		return message, err
	}
	return message, nil
}