package services

import (

	"github.com/valitovgaziz/rest-api/models"
	"github.com/valitovgaziz/rest-api/storage"
)

func CreateMessage(messageInput *models.MessageInput) (*models.Message, error) {
	message := models.CreateNewMessage(messageInput)
	message, err := storage.SaveMessage(message)
	SaveSatatistics(message)
	SendKafkaMessage(message)
	return message, err
}