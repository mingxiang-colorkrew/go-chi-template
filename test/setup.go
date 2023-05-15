package test

import (
	"context"
	"log"
	"measure/config"
	"measure/oapi"
	"measure/webserver"
	"net/http"
	"net/http/httptest"
	"testing"
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

func SetupTestApp(t *testing.T) *config.App {
	app := config.NewApp()
	app.UseTestDB()

	t.Cleanup(func() {
		app.DB().Close()
	})

	return app
}

func SetupTestClient(app *config.App) *oapi.ClientWithResponses {
	ws := webserver.NewWebserver(app)

	fakeClient := fakeHttpClient{
		server: ws.Router(),
	}

	client, err := oapi.NewClientWithResponses("", oapi.WithHTTPClient(&fakeClient))

	if err != nil {
		log.Fatal(err)
	}

	return client
}

func SetupAuthHeaders(app *config.App) oapi.RequestEditorFn {
	return func (ctx context.Context, req *http.Request) error {
		// TODO: hardcoded user ID, can setup seed user ID beforehand
		_, tokenStr, _ := app.JWTAuth().Encode(map[string]interface{}{"user_id": 123})
		req.Header.Add("Authorization", "Bearer " + tokenStr)

		return nil
	}
}
