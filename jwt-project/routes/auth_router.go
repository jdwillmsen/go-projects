package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jdwillmsen/jwt-project/controllers"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/auth/sign-up", controllers.SignUp())
	incomingRoutes.POST("/auth/login", controllers.Login())
}
