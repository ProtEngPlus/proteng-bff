package main

import (
	"fmt"
	"log"
	"net/http"
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

	// CORS
	r.Use(CORSMiddleware())

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

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Writer.Header().Set("Access-Control-Allow-Origin", fmt.Sprintf("้http://localhost:5173, http://localhost:3000, %s", os.Getenv("FRONTEND_URL")))
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
