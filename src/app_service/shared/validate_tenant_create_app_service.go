package shared

import (
	"measure/config"
	"measure/oapi"

	"github.com/asaskevich/govalidator"
)

func ValidateTenantCreateAppService(
	app *config.App,
	p *oapi.PostApiV1TenantJSONRequestBody,
) (*oapi.PostApiV1Tenant400JSONResponse, error) {
	errData := oapi.TenantCreateValidationError{}

	var nameErrs []string
	var shortCodeErrs []string

	if !govalidator.StringLength(p.Name, "1", "255") {
		nameErrs = append(nameErrs, "validate.name:1,255")
	}

	if !govalidator.StringLength(p.ShortCode, "4", "5") {
		shortCodeErrs = append(shortCodeErrs, "validate.short_code:4,5")
	}

	errData.Name = &nameErrs
	errData.ShortCode = &nameErrs

	if hasNotNilField(&errData) {
		errResp := oapi.PostApiV1Tenant400JSONResponse{}
		errResp.Data = &errData

		return &errResp, nil
	}

	return nil, nil
}
