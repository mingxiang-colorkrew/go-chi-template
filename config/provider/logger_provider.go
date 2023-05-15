package provider

import "go.uber.org/zap"

func NewLoggerProvider(env *EnvProvider) *zap.Logger {
	var logger *zap.Logger

	if env.logLevel == "debug" {
		logger, _ = zap.NewDevelopment()
	} else {
		logger, _ = zap.NewProduction()
	}

	defer logger.Sync()

	return logger
}
