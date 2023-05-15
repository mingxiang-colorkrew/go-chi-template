package handler

import (
	"context"
	"encoding/json"
	"measure/test"
	"measure/test/factory"
	"testing"
)

func TestGetTenant(t *testing.T) {
	app := test.SetupTestApp(t)
	oapiClient := test.SetupTestClient(app)

	factory.SeedTenant(app, "Tenant A", "test0")

	authHeaders := test.SetupAuthHeaders(app)
	res, _ := oapiClient.GetApiV1Tenant(context.Background(), authHeaders)

	app.Logger().Sugar().Infow("test", "body", json.NewDecoder(res.Body))
	res.Body.Close()
}
