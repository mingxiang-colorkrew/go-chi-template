package handler

import (
	"context"
	"measure/oapi"
	v1 "measure/src/app_service/v1"
)

func (handler *Handler) GetApiV1Tenant(
	ctx context.Context,
	request oapi.GetApiV1TenantRequestObject,
) (oapi.GetApiV1TenantResponseObject, error) {
	resp, err := v1.ListTenantAppService(handler.app, request)
	return resp, err
}

func (handler *Handler) PostApiV1Tenant(
	ctx context.Context,
	request oapi.PostApiV1TenantRequestObject,
) (oapi.PostApiV1TenantResponseObject, error) {
	resp, err := v1.CreateTenantAppService(handler.app, request)
	return resp, err
}

func (handler *Handler) GetApiV1TenantTenantId(
	ctx context.Context,
	request oapi.GetApiV1TenantTenantIdRequestObject,
) (oapi.GetApiV1TenantTenantIdResponseObject, error) {
	resp, err := v1.DetailTenantAppService(handler.app, request)
	return resp, err
}
