package handler

import (
	"context"
	"encoding/json"
	"measure/oapi"
	"measure/test"
	tf "measure/test/factory/tenant_factory"
	"testing"
)

func TestGetTenant(t *testing.T) {
	app := test.SetupTestApp(t)
	oapiClient := test.SetupTestClient(app)

	tf.SeedTenant(app, tf.WithName("Tenant A"), tf.WithShortCode("aaay2"))

	authHeaders := test.SetupAuthHeaders(app)
	res, _ := oapiClient.GetApiV1Tenant(context.Background(), authHeaders)

	if res.StatusCode != 200 {
		t.FailNow()
	}

	var respDto oapi.GetApiV1Tenant200JSONResponse
	json.NewDecoder(res.Body).Decode(&respDto)
	res.Body.Close()

	if len(respDto.Tenants) != 1 {
		t.FailNow()
	}

	if respDto.Tenants[0].ShortCode != "aaay2" {
		t.FailNow()
	}
}
