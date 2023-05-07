package handler

import (
	"context"
	"measure/oapi"
	"measure/src/app_service/shared"
	v1 "measure/src/app_service/v1"
)

func (h *Handler) GetApiV1Tenant(
	ctx context.Context,
	request oapi.GetApiV1TenantRequestObject,
) (oapi.GetApiV1TenantResponseObject, error) {
	resp, err := v1.TenantListAppService(h.app, request)
	return resp, err
}

func (h *Handler) PostApiV1Tenant(
	ctx context.Context,
	request oapi.PostApiV1TenantRequestObject,
) (oapi.PostApiV1TenantResponseObject, error) {
	errResp, err := shared.ValidateTenantCreateAppService(h.app, request.Body)

	if errResp != nil {
		return errResp, err
	}

	resp, err := v1.TenantCreateAppService(h.app, request)

	return resp, err
}

func (h *Handler) GetApiV1TenantTenantId(
	ctx context.Context,
	request oapi.GetApiV1TenantTenantIdRequestObject,
) (oapi.GetApiV1TenantTenantIdResponseObject, error) {
	resp, err := v1.TenantDetailAppService(h.app, request)
	return resp, err
}
