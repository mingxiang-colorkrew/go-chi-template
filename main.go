package main

import (
	"measure/config"
	"measure/webserver/handler"
	"os"
)

func main() {
	command := ""
	if len(os.Args[1:]) > 0 {
		command = os.Args[1]
	}

	appEnv := config.GetAppEnv()
	app := config.NewApp(appEnv)
	h := handler.NewHandler(app)
	app.SetupRouter(h)

	if command == "routes:list" {
		app.PrintRoutes()
	} else {
		app.Start()
	}
}
