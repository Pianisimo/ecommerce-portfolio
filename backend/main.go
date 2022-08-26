package main

import (
	"backend/middleware"
	"backend/myJwt"
	"backend/routes"
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
		AllowOrigins: []string{
			"http://" + frontendURL + ":" + frontendPort,
			"http://localhost:3000", // Always allow origin for development
		},
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
	routes.PublicRoutes(router)

	// Check for authorization
	router.Use(middleware.Authentication())

	// Setup authorized routes
	routes.AuthRoutes(router)

	// Init server API
	log.Fatal(router.Run(":" + backendPort))
}
