package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type AppEnv int64

const (
	AppEnvLocal       AppEnv = iota
	AppEnvTest        AppEnv = iota
	AppEnvDevelopment AppEnv = iota
	AppEnvProduction  AppEnv = iota
)

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

func readEnv(env AppEnv) map[string]string {
	basepath, _ := os.Getwd()
	// basepath := filepath.Base(b)

	var envFileName string

	switch env {
	case AppEnvTest:
		envFileName = ".env.test"
	case AppEnvLocal:
		envFileName = ".env.local"
	default:
		return map[string]string{}
	}

	envFilePath := filepath.Join(basepath, envFileName)

	envVars, error := godotenv.Read(envFilePath)

	if error != nil {
		errMsg := fmt.Sprintf("Unable to load env file %s", envFilePath)
		log.Fatal(errMsg)
	}

	return envVars
}

func GetAppEnv() AppEnv {
	appEnv, appEnvIsSet := os.LookupEnv("APP_ENV")

	if appEnvIsSet == false {
		appEnv = "local"
	}

	switch appEnv {
	case "local":
		return AppEnvLocal
	case "test":
		return AppEnvTest
	case "development":
		return AppEnvDevelopment
	case "production":
		return AppEnvProduction
	default:
		log.Fatal("Invalid APP_ENV provided")
	}

	return AppEnvLocal
}

func NewEnvProvider(appEnv AppEnv) *EnvProvider {
	envVars := readEnv(appEnv)

	appServer, exists := envVars["APP_ENV"]
	if exists == false {
		appServer = "local"
	}

	serverPort, exists := envVars["SERVER_PORT"]
	if exists == false {
		serverPort = "3000"
	}

	databaseUrl, exists := envVars["DATABASE_URL"]

	envProvider := EnvProvider{
		appEnv:      appServer,
		serverPort:  serverPort,
		databaseUrl: databaseUrl,
	}

	return &envProvider
}
