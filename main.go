package main

import (
	"log"
	"os"

	"proteng-bff/apis"
	"proteng-bff/configs"

	"github.com/gin-gonic/gin"
)

func main() {
	configs.AutomaticLoadEnv()

	r := gin.Default()

	// Health Check Endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "ok"})
	})

	// Init Router
	apis.InitRouter(r)

	// Start Server
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}
	err := r.Run(":" + httpPort)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
