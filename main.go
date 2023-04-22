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

	a := config.NewApp()
	h := handler.NewHandler(a)
	webserver := config.NewWebserver(a, h)

	if command == "routes:list" {
		webserver.PrintRoutes()
	} else {
		webserver.Start()
	}
}
