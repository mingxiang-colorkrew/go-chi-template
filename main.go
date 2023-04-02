package main

import (
	"measure/config"
	"measure/webserver/route"
)

func main() {
	appEnv := config.GetAppEnv()
	app := config.NewApp(appEnv)
	route.SetupRoute(app)
	app.Start()
}
