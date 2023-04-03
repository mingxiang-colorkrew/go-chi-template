package v1

import (
	"fmt"
	"measure/config"
	"measure/oapi"
	domainservice "measure/src/domain_service"
	"measure/src/mutation"
)

func CreateTenantAppService(app *config.App, req oapi.PostTenantRequestObject) (*oapi.PostTenant200JSONResponse, error) {
	newTenant := domainservice.NewTenant(req.Body.Name, req.Body.ShortCode)
	insertedTenant, err := mutation.InsertTenant(app, newTenant)

	if err != nil {
		return nil, err
	}

	resp := oapi.PostTenant200JSONResponse{
		Tenant: &oapi.Tenant{
			Id:        fmt.Sprint(insertedTenant.ID),
			Name:      insertedTenant.Name,
			ShortCode: insertedTenant.ShortCode,
		},
	}

	return &resp, nil
}
