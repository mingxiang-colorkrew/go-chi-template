package v1

import (
	"measure/config"
	"measure/oapi"
	"measure/src/repository/multi"
	"strconv"
)

func TenantListAppService(
	app *config.App,
	_ oapi.GetApiV1TenantRequestObject,
) (oapi.GetApiV1TenantResponseObject, error) {
	tenants, _ := multi.GetTenantsWithUserCount(app)

	tenantDtos := []oapi.Tenant{}
	for _, tenant := range tenants {
		tenantDto := oapi.Tenant{
			Id:        strconv.FormatInt(tenant.ID, 10),
			Name:      tenant.Name,
			ShortCode: tenant.ShortCode,
			UserCount: &tenant.UserCount,
		}
		tenantDtos = append(tenantDtos, tenantDto)
	}

	return &oapi.GetApiV1Tenant200JSONResponse{
		Tenants: tenantDtos,
	}, nil
}
