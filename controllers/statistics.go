package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/valitovgaziz/micro-message/services"
)

// GetStatisticsByUserId godoc: Get statistics by user ID
func GetStatisticsByUserId(c *gin.Context) {
	// TODO: Implement this method
	UserId := c.Param("id")
	UserStatistics, err := services.GetSatatisticsByUserId(UserId)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{"userStatistics": UserStatistics})
}
