package initializers

import (
	"github.com/gin-gonic/gin"
	"github.com/valitovgaziz/micro-message/controllers"
)

func InitServerAndRouting(config *Config) {
	
	// Initialize
	r := gin.Default()

	// Routes
	r.POST("/api/message", controllers.SaveMessage)
	r.GET("/api/statistic", controllers.GetStatistics)

	// StartServer
	r.Run(":" + config.ServerPort)
}

