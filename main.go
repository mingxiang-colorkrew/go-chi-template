package main

import (
	"measure/config"
	"measure/webserver/handler"
)

func main() {
	appEnv := config.GetAppEnv()
	app := config.NewApp(appEnv)
	handler := handler.NewHandler(app)
	app.SetupRouter(handler)
  app.PrintRoutes()
	app.Start()
}
