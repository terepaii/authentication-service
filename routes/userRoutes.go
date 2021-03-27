package routes

import (
	controller "authentication-service/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/register", controller.Register())
	incomingRoutes.POST("/users/login", controller.Login())
}
