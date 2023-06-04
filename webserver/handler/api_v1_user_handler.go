package handler

import (
	"context"
	"go_chi_template/oapi"
	"go_chi_template/src/app_service/shared"
	v1 "go_chi_template/src/app_service/v1"
)

func (h *Handler) PostApiV1User(
	ctx context.Context,
	request oapi.PostApiV1UserRequestObject,
) (oapi.PostApiV1UserResponseObject, error) {
	errResp, err := shared.ValidateUserCreateAppService(ctx, h.app, request.Body)

	if errResp != nil {
		return errResp, err
	}

	resp, err := v1.UserCreateAppService(ctx, h.app, request)

	return resp, err
}

func (h *Handler) GetApiV1User(
	ctx context.Context,
	request oapi.GetApiV1UserRequestObject,
) (oapi.GetApiV1UserResponseObject, error) {
	resp, err := v1.UserListAppService(ctx, h.app, request)
	return resp, err
}

func (h *Handler) GetApiV1UserUserId(
	ctx context.Context,
	request oapi.GetApiV1UserUserIdRequestObject,
) (oapi.GetApiV1UserUserIdResponseObject, error) {
	resp, err := v1.UserDetailAppService(ctx, h.app, request)
	return resp, err
}
