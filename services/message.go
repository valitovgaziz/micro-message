package services

import (
	"github.com/valitovgaziz/micro-message/models"
	"github.com/valitovgaziz/micro-message/storage"
)

func CreateMessage(message *models.Message) (*models.Message, error) {
	message, err := storage.SaveMessage(message)
	// TODO save to satatistics
	// TODO send to kafka
	return message, err
}