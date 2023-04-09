package v1

import (
	"measure/config"
	"measure/oapi"
	"measure/src/repository/single"
	"strconv"
)

func DetailTenantAppService(
	app *config.App,
	req oapi.GetApiV1TenantTenantIdRequestObject,
) (oapi.GetApiV1TenantTenantIdResponseObject, error) {
	tenantId, _ := strconv.ParseInt(req.TenantId, 10, 64)

	tenant, _ := single.GetTenantById(app, tenantId)

	tenantDto := oapi.Tenant{
		Id:        strconv.FormatInt(tenant.ID, 10),
		Name:      tenant.Name,
		ShortCode: tenant.ShortCode,
	}

	return oapi.GetApiV1TenantTenantId200JSONResponse{
		Tenant: tenantDto,
	}, nil
}
