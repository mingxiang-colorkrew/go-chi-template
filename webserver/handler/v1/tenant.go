package v1

import (
	"measure/config"
	v1 "measure/src/app_service/v1"
	"net/http"
)

func GetListTenant(app *config.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	}
}

func PostCreateTenant(app *config.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := v1.CreateTenantAppService(app)

		w.Write([]byte(name))
	}
}
