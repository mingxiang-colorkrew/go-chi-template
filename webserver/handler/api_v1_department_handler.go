package handler

import (
	"context"
	"go_chi_template/oapi"
	v1 "go_chi_template/src/app_service/v1"
)

func (h *Handler) GetApiV1Department(
	ctx context.Context,
	request oapi.GetApiV1DepartmentRequestObject,
) (oapi.GetApiV1DepartmentResponseObject, error) {
	resp, err := v1.DepartmentListAppService(ctx, h.app, request)
	return resp, err
}
