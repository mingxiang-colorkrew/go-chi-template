package main

import (
	"measure/config"
	"measure/webserver/handler"
)

func main() {
	appEnv := config.GetAppEnv()
	app := config.NewApp(appEnv)
	h := handler.NewHandler(app)
	app.SetupRouter(h)
	app.PrintRoutes()
	app.Start()
}
