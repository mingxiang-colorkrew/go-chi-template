package v1

import (
	"context"
	"go_chi_template/config"
	"go_chi_template/oapi"
	"go_chi_template/src/repository/single"
	"strconv"
)

func TenantDetailAppService(
	ctx context.Context,
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

	respDto := oapi.GetApiV1TenantTenantId200JSONResponse{
		Tenant: tenantDto,
	}

	return &respDto, nil
}
