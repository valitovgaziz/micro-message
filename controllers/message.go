package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/valitovgaziz/micro-message/models"
	"github.com/valitovgaziz/micro-message/services"
)

func SaveMessage(c *gin.Context) {
	message := &models.Message{}
	if err := c.BindJSON(&message); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	message, err := services.CreateMessage(message)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"message": "posted",
	})
}