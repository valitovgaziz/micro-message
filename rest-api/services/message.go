package services

import (
	"github.com/valitovgaziz/rest-api/models"
	"github.com/valitovgaziz/rest-api/storage"
)

func CreateMessage(messageInput *models.MessageInput) (*models.Message, error) {
	message := models.CreateNewMessage(messageInput)
	message, err := storage.SaveMessage(message)
	if err!= nil {
		return nil, err
	}
	SaveSatatistics(message)
	err = SendKafkaTestTopic(message)
	if err!= nil {
		return message, err
	}
	return message, err
}
