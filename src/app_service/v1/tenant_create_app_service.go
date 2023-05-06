package v1

import (
	"fmt"
	"measure/config"
	"measure/oapi"
	domainservice "measure/src/domain_service"
	"measure/src/mutation"
	"measure/src/repository/single"
)

func TenantCreateAppService(
	app *config.App,
	req oapi.PostApiV1TenantRequestObject,
) (oapi.PostApiV1TenantResponseObject, error) {
	existingTenant, _ := single.GetTenantByShortCode(app, req.Body.ShortCode)

	if existingTenant != nil {
		test := new(oapi.PostApiV1Tenant400JSONResponse)
		test.Data.ShortCode = &[]string{"validation.unique"}
		return test, nil
	}

	newTenant := domainservice.NewTenant(req.Body.Name, req.Body.ShortCode)
	insertedTenant, insertErr := mutation.InsertTenant(app, newTenant)

	if insertErr != nil {
		return nil, insertErr
	}

	resp := oapi.PostApiV1Tenant200JSONResponse{
		Tenant: oapi.Tenant{
			Id:        fmt.Sprint(insertedTenant.ID),
			Name:      insertedTenant.Name,
			ShortCode: insertedTenant.ShortCode,
		},
	}

	return &resp, nil
}
