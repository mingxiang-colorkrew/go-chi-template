package v1

import (
	"fmt"
	"measure/config"
	"measure/oapi"
	domainservice "measure/src/domain_service"
	"measure/src/mutation"
	"measure/src/repository/single"
)

func CreateTenantAppService(
	app *config.App,
	req oapi.PostApiV1TenantRequestObject,
) (oapi.PostApiV1TenantResponseObject, error) {
	existingTenant, _ := single.GetTenantByShortCode(app, req.Body.ShortCode)

	if existingTenant != nil {
		errCode := "validation_07x4g2"
		errMsg := "Validation failed"
		errData := struct {
			Name      *[]string `json:"name"`
			ShortCode *[]string `json:"shortCode"`
		}{
			Name:      &[]string{},
			ShortCode: &[]string{"validation.unique"},
		}

		resp := oapi.PostApiV1Tenant400JSONResponse{
			ErrorCode:    &errCode,
			ErrorMessage: &errMsg,
			Data:         &errData,
		}
		return &resp, nil
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
