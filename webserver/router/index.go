package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Setup(router *chi.Mux) *chi.Mux {
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	router.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})
	return router
}
