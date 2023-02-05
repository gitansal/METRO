package routes

import (
	controllers "github.com/gitansal/metro/controllers"

	"github.com/gin-gonic/gin"
)

func appointmentRoutes(incomingRoutes *gin.Engine) {
	//incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.POST("appointments/schedule", controllers.Schedule())
	//incomingRoutes.GET("/download/:image_name",controller.DownloadFile())
}
