package config

import "go.uber.org/zap"

func (app *App) setupLogger() *zap.Logger {
	logger, _ := zap.NewProduction()

	defer logger.Sync()

	return logger
}
