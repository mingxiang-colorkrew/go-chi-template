package provider

import "go.uber.org/zap"

func NewLoggerProvider() *zap.Logger {
	logger, _ := zap.NewProduction()

	defer logger.Sync()

	return logger
}
