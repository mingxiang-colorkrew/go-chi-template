package provider

import "os"

type EnvProvider struct {
	appEnv        string
	serverPort    string
	databaseUrl   string
	redisUrl      string
	redisPassword string
	jwtSecret     string
	logLevel      string
}

func (e *EnvProvider) AppEnv() string {
	return e.appEnv
}

func (e *EnvProvider) ServerPort() string {
	return e.serverPort
}

func NewEnvProvider(rootDir string) *EnvProvider {
	appServer, exists := os.LookupEnv("APP_ENV")
	if exists == false {
		appServer = "local"
	}

	serverPort, exists := os.LookupEnv("SERVER_PORT")
	if exists == false {
		serverPort = "3000"
	}

	databaseUrl, exists := os.LookupEnv("DATABASE_URL")
	redisUrl, exists := os.LookupEnv("REDIS_URL")
	redisPassword, exists := os.LookupEnv("REDIS_PASSWORD")
	jwtSecret, exists := os.LookupEnv("JWT_SECRET")
	logLevel, exists := os.LookupEnv("LOG_LEVEL")

	envProvider := EnvProvider{
		appEnv:        appServer,
		serverPort:    serverPort,
		databaseUrl:   databaseUrl,
		redisUrl:      redisUrl,
		redisPassword: redisPassword,
		jwtSecret:     jwtSecret,
		logLevel:      logLevel,
	}

	return &envProvider
}
