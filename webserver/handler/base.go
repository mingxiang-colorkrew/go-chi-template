package handler

import (
	"go_chi_template/config"
	"go_chi_template/oapi"
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
