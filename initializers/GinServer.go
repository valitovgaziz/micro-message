package initializers

import "github.com/gin-gonic/gin"

func InitRouting(config *Config) {
	r := gin.Default()
	r.POST("/api/message", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "posted",
		})
	})
	r.GET("/api/statistic", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "getted",
		})
	})
	r.Run(":" + config.ServerPort)
}
