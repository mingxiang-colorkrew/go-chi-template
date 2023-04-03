package handler

import (
	"measure/config"
	"measure/oapi"
)

type Handler struct {
	app *config.App
}

func NewHandler(app *config.App) oapi.StrictServerInterface {
	handler := Handler{
		app: app,
	}
	return &handler
}
