package integration

import (
	"log"
	"measure/oapi"
	"measure/test"
	"testing"
)

func TestPostCreateTenant(t *testing.T) {
	app := test.SetupTestApp()
	client := test.SetupClient(app)

	reqBody := oapi.PostTenantJSONRequestBody{
		Name:      "Tenant A",
		ShortCode: "123",
	}

	resp, err := client.PostTenant(nil, reqBody)
	log.Fatal(resp, err)

	t.Fatal("Not yet implemented.")
}
