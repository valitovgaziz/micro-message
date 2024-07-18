package services

import (
	"github.com/valitovgaziz/micro-message/models"
	"github.com/valitovgaziz/micro-message/storage"
)

func CreateMessage(messageInput *models.MessageInput) (*models.Message, error) {
	message := models.CreateNewMessage(messageInput)
	message, err := storage.SaveMessage(message)
	SaveSatatistics(message)
	// TODO send to kafka
	return message, err
}
