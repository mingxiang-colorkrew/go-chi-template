package v1

import (
	"measure/config"
	"measure/oapi"
	"measure/src/repository/single"
	"strconv"
)

func ListTenantAppService(
	app *config.App,
	_ oapi.GetApiV1TenantRequestObject,
) (oapi.GetApiV1TenantResponseObject, error) {
	tenants, _ := single.GetAllTenants(app)

	tenantDtos := []oapi.Tenant{}
	for _, tenant := range tenants {
		tenantDto := oapi.Tenant{
			Id:        strconv.FormatInt(tenant.ID, 10),
			Name:      tenant.Name,
			ShortCode: tenant.ShortCode,
		}
		tenantDtos = append(tenantDtos, tenantDto)
	}

	return oapi.GetApiV1Tenant200JSONResponse{
		Tenants: tenantDtos,
	}, nil
}
