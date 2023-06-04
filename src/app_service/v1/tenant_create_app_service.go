package v1

import (
	"fmt"
	"go_chi_template/config"
	"go_chi_template/oapi"
	domainservice "go_chi_template/src/domain_service"
	"go_chi_template/src/mutation"

	"go.uber.org/zap"
)

func TenantCreateAppService(
	app *config.App,
	req oapi.PostApiV1TenantRequestObject,
) (oapi.PostApiV1TenantResponseObject, error) {
	newTenant := domainservice.NewTenant(req.Body.Name, req.Body.ShortCode)
	insertedTenant, insertErr := mutation.InsertTenant(app, newTenant)

	if insertErr != nil {
		app.Logger().With(
			zap.String("message", "insert tenant failed"),
			zap.String("name", req.Body.Name),
			zap.String("shortCode", req.Body.ShortCode),
		)

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
