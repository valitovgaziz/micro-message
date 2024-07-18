package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/valitovgaziz/micro-message/models"
	"github.com/valitovgaziz/micro-message/services"
)

func SaveMessage(c *gin.Context) {
	messageInput := &models.MessageInput{}
	if err := c.BindJSON(&messageInput); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	message, err := services.CreateMessage(messageInput)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"message": message.Id,
	})
}