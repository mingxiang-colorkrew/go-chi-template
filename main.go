package main

import (
	"measure/config"
)

func main() {
	appEnv := config.GetAppEnv()
	app := config.NewApp(appEnv)
	app.Start()
}
