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
	ws := config.NewWebserver(a, h)

	if command == "routes:list" {
		ws.PrintRoutes()
	} else {
		ws.Start()
	}
}
