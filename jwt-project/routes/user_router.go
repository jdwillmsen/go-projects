package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jdwillmsen/jwt-project/controllers"
	"github.com/jdwillmsen/jwt-project/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.GET("/users/:user-id", controllers.GetUser())
}
