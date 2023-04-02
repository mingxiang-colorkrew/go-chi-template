package v1

import (
	"measure/config"
	"net/http"
)

func GetListUser(app *config.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	}
}
