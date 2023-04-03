package handler

import (
	"context"
	"measure/oapi"
	v1 "measure/src/app_service/v1"
)

func (handler *Handler) PostTenant(ctx context.Context, req oapi.PostTenantRequestObject) (oapi.PostTenantResponseObject, error) {

	return v1.CreateTenantAppService(handler.app, req)
}
