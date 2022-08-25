package main

import (
	"ecommerce/myJwt"
	"ecommerce/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

func main() {
	initMain()
}

func initMain() {
	myJwt.InitJWT()

	// Comment below to use env variables coming from a docker-compose.yml
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	backendPort := os.Getenv("BACKEND_PORT")
	if backendPort == "" {
		backendPort = "8000"
	}

	backendURL := os.Getenv("BACKEND_URL")
	if backendURL == "" {
		backendURL = "localhost"
	}

	frontendPort := os.Getenv("FRONTEND_PORT")
	if frontendPort == "" {
		frontendPort = "3000"
	}

	frontendURL := os.Getenv("FRONTEND_URL")
	if frontendURL == "" {
		frontendURL = "localhost"
	}

	//app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://" + frontendURL + ":" + frontendPort},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	// Setup unauthorized routes
	routes.UserRoutes(router)

	/*
		// Check for authorization
		router.Use(middleware.Authentication())

		// Setup authorized routes
		router.GET("/addtocart", app.AddToCart())
		router.GET("/removeitem", app.RemoveItem())
		router.GET("/listcart", controllers.GetItemsFromCart())
		router.POST("/addaddress", controllers.AddAddress())
		router.PUT("/edithomeaddress", controllers.EditHomeAddress())
		router.PUT("/editworkaddress", controllers.EditWorkAddress())
		router.GET("/deleteaddresses", controllers.DeleteAddress())
		router.GET("/cartcheckout", app.BuyFromCart())
		router.GET("/instantbuy", app.InstantBuy())
	*/

	// Init server API
	log.Fatal(router.Run(":" + backendPort))
}
