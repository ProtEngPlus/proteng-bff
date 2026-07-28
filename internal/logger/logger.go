package logger

import (
	"os"

	"go.uber.org/zap"
)

var Zap *zap.Logger

func InitZap() {
	// Note: use os.Getenv("ENV") because we will initialize the logger before loading the env
	if os.Getenv("ENV") == "prod" {
		_logger, _ := zap.NewProduction()
		Zap = _logger
	} else {
		_logger, _ := zap.NewDevelopment()
		Zap = _logger
	}

	Zap.Info("Logger initialized")
}

func CloseZap() {
	_ = Zap.Sync()
}
