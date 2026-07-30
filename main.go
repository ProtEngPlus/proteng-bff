package main

import (
	"fmt"
	"time"

	"github.com/protengplus/proteng-bff/apis"
	"github.com/protengplus/proteng-bff/configs"
	"github.com/protengplus/proteng-bff/internal/logger"

	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"

	docs "github.com/protengplus/proteng-bff/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Proteng API
// @version 1.0
// @description

// @contact.name Proteng API Support
// @contact.email proteng.plus@gmail.com

// @SecurityDefinitions.apikey ApiKeyAuth
// @In Header
// @Name Authorization
// @Type apiKey
// @BearerFormat

func main() {
	logger.InitZap()
	configs.AutomaticLoadEnv()

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	docs.SwaggerInfo.BasePath = "/"

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Health Check Endpoint
	r.GET("/health", HealthCheck)

	// CORS
	r.Use(CORSMiddleware())

	// Logger
	r.Use(ginzap.GinzapWithConfig(logger.Zap, &ginzap.Config{
		TimeFormat: time.RFC3339,
		UTC:        true,
		SkipPaths:  []string{"/metrics", "/health"},
	}))

	// Init Router
	apis.InitRouter(r)

	// Panic Recovery
	r.Use(ginzap.RecoveryWithZap(logger.Zap, true))

	// Start Server
	httpPort := configs.Config.HttpPort
	if httpPort == "" {
		httpPort = "8080"
	}
	logger.Zap.Info("proteng-bff is running on :" + httpPort)
	err := r.Run(":" + httpPort)
	if err != nil {
		logger.Zap.Fatal(fmt.Sprintf("Error starting server: %v", err))
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     append([]string{"http://localhost:5173", "http://localhost:3000"}, configs.Config.FrontendUrls...),
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "Cache-Control", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}

// HealthCheck godoc
// @Summary Health Check Endpoint
// @Description Check the health of the service
// @Tags Healthchack
// @Produce json
// @Success 200 {object} models.HttpResponseOK "Successful operation"
// @Router /health [get]
func HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{"message": "ok"})
}
