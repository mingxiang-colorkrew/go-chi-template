package test

import (
	"log"
	"measure/config"
	"measure/oapi"
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
	a := config.NewApp()
	h := handler.NewHandler(a)
	ws := config.NewWebserver(a, h)

	return app
}

func SetupClient(app *config.App) *oapi.ClientWithResponses {
	fakeClient := fakeHttpClient{
		server: server,
	}

	client, err := oapi.NewClientWithResponses("", oapi.WithHTTPClient(&fakeClient))

	if err != nil {
		log.Fatal(err)
	}
	return client
}
