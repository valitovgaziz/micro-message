package initializers

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/valitovgaziz/rest-api/controllers"
)

func InitServerAndRouting() {

	// Initialize
	r := gin.Default()

	// Routes
	r.POST("/api/message", controllers.SaveMessage)
	r.GET("/api/statistics/byuserid", controllers.GetStatisticsByUserId)

	// StartServer
	r.Run(":" + os.Getenv("SERVER_PORT"))
}
