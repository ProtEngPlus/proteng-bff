package configs

import (
	"fmt"
	"os"

	"github.com/protengplus/proteng-bff/internal/logger"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

var Config config

type config struct {
	Env      string `envconfig:"ENV" default:"dev"`
	HttpPort string `envconfig:"HTTP_PORT" default:"8080"`

	UserMgmtUrl  string `envconfig:"USER_MGMT_URL" default:"http://localhost:8082"`
	ConductorUrl string `envconfig:"CONDUCTOR_URL" default:"http://localhost:8081"`
	FrontendUrl  string `envconfig:"FRONTEND_URL" default:"http://localhost:5173"`

	AccessTokenPublicKey string `envconfig:"ACCESS_TOKEN_PUBLIC_KEY" default:""`
}

func AutomaticLoadEnv() {
	if env, ok := os.LookupEnv("ENV"); ok {
		if err := LoadEnvFromPath(".env." + env); err != nil {
			logger.Zap.Fatal(fmt.Sprintf("Error loading .env file from file: %v", err))
		} else {
			logger.Zap.Info(fmt.Sprintf("Running in environment: %s", env))
		}
	}

	err := envconfig.Process("", &Config)
	if err != nil {
		logger.Zap.Fatal(fmt.Sprintf("Error unmarshalling env vars: %v", err))
	}
}

func LoadEnvFromPath(path string) error {
	return godotenv.Load(path)
}
