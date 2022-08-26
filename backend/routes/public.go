package routes

import (
	"backend/controllers"
	"github.com/gin-gonic/gin"
)

func PublicRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", controllers.SignUp())
	incomingRoutes.POST("/users/login", controllers.Login())
	incomingRoutes.GET("/users/auth", controllers.IsAuth())

	//incomingRoutes.POST("/admin/addproduct", controllers.ProductViewerAdmin())
	//incomingRoutes.GET("/users/productview", controllers.SearchProduct())
	//incomingRoutes.GET("/users/search", controllers.SearchProductByQuery())
}
