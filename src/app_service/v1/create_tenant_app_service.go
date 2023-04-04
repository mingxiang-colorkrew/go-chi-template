package v1

import (
	"fmt"
	"measure/config"
	"measure/oapi"
	domainservice "measure/src/domain_service"
	"measure/src/mutation"
	"measure/src/repository"
)

func CreateTenantAppService(app *config.App, req oapi.PostTenantRequestObject) (oapi.PostTenantResponseObject, error) {
	existingTenant, _ := repository.GetTenantByShortCode(app.DB(), req.Body.ShortCode)

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

		resp := oapi.PostTenant400JSONResponse{
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

	resp := oapi.PostTenant200JSONResponse{
		Tenant: &oapi.Tenant{
			Id:        fmt.Sprint(insertedTenant.ID),
			Name:      insertedTenant.Name,
			ShortCode: insertedTenant.ShortCode,
		},
	}

	return &resp, nil
}
