package config

import (
	"fmt"
	"log"
	"measure/oapi"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type Webserver struct {
	router     *chi.Mux
	serverAddr string
	app        *App
}

func setupLoggerMiddleware(l *zap.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
			t1 := time.Now()
			defer func() {
				reqLogger := l.With(
					zap.String("proto", r.Proto),
					zap.String("path", r.URL.Path),
					zap.String("reqId", middleware.GetReqID(r.Context())),
					zap.Duration("lat", time.Since(t1)),
					zap.Int("status", ww.Status()),
					zap.Int("size", ww.BytesWritten()),
				)
				reqLogger.Info("Served")
			}()
			next.ServeHTTP(ww, r)
		}
		return http.HandlerFunc(fn)
	}
}

func NewWebserver(app *App, handler oapi.StrictServerInterface) *Webserver {
	serverAddr := ":" + app.EnvVars().ServerPort()

	r := chi.NewRouter()
	r.Use(setupLoggerMiddleware(app.Logger()))

	baseUrl := ""
	strictHandler := oapi.NewStrictHandler(handler, []oapi.StrictMiddlewareFunc{})
	oapi.HandlerFromMuxWithBaseURL(strictHandler, r, baseUrl)

	return &Webserver{
		router:     r,
		serverAddr: serverAddr,
	}
}

func (ws *Webserver) Start() {

	log.Print("WebServer listening on " + ws.serverAddr)

	s := &http.Server{
		Handler: ws.router,
		Addr:    ws.serverAddr,
	}

	log.Fatal(s.ListenAndServe())
}

func (ws *Webserver) PrintRoutes() {
	chi.Walk(
		ws.router,
		func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
			fmt.Printf("[%s]: '%s' has %d middlewares\n", method, route, len(middlewares))
			return nil
		},
	)
}
