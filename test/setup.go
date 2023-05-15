package test

import (
	"log"
	"measure/config"
	"measure/oapi"
	"measure/webserver"
	"measure/webserver/handler"
	"net/http"
	"net/http/httptest"
)

// mock HttpClient to work against generated openapi client
type fakeHttpServer interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type fakeHttpClient struct {
	server fakeHttpServer
}

func (c *fakeHttpClient) Do(r *http.Request) (*http.Response, error) {
	rr := httptest.NewRecorder()
	c.server.ServeHTTP(rr, r)
	return rr.Result(), nil
}

func SetupTestApp() *config.App {
	app := config.NewApp()
	app.UseTestDB()

	return app
}

func SetupClient(app *config.App) *oapi.ClientWithResponses {
	ws := webserver.NewWebserver(app)

	s := &http.Server{
		Handler: ws.router,
		Addr:    ws.serverAddr,
	}

	fakeClient := fakeHttpClient{
		server: ws,
	}

	client, err := oapi.NewClientWithResponses("", oapi.WithHTTPClient(&fakeClient))

	if err != nil {
		log.Fatal(err)
	}

	return client
}
