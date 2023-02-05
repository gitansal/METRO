package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.appointmentRoutes(router)
	router.GET("/api-3", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-3"})
	})

	router.Run(":" + port)
}
