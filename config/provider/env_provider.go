package provider

import "os"

type EnvProvider struct {
	appEnv      string
	serverPort  string
	databaseUrl string
}

func (e *EnvProvider) AppEnv() string {
	return e.appEnv
}

func (e *EnvProvider) ServerPort() string {
	return e.serverPort
}

func (e *EnvProvider) DatabaseUrl() string {
	return e.databaseUrl
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

	envProvider := EnvProvider{
		appEnv:      appServer,
		serverPort:  serverPort,
		databaseUrl: databaseUrl,
	}

	return &envProvider
}
