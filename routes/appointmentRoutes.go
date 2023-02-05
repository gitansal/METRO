package routes

import (
	controllers "github.com/gitansal/METRO/controllers"

	"github.com/gin-gonic/gin"
)

func AppointmentRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("appointment/schedule", controllers.Schedule())
}
