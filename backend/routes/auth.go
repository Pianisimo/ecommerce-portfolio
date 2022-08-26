package routes

import (
	"backend/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/users/logout", controllers.Logout())
}
