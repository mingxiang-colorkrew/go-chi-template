package handler

import (
	"context"
	"measure/oapi"
	"measure/src/app_service/shared"
	v1 "measure/src/app_service/v1"
)

func (h *Handler) PostApiV1User(
	ctx context.Context,
	request oapi.PostApiV1UserRequestObject,
) (oapi.PostApiV1UserResponseObject, error) {
	errResp, err := shared.ValidateUserCreateAppService(h.app, request.Body)

	if errResp != nil {
		return errResp, err
	}

  resp, err := v1.UserCreateAppService(h.app, request)

	return resp, err
}

func (h *Handler) GetApiV1User(
	ctx context.Context,
	request oapi.GetApiV1UserRequestObject,
) (oapi.GetApiV1UserResponseObject, error) {
	resp, err := v1.UserListAppService(h.app, request)
	return resp, err
}

func (h *Handler) GetApiV1UserUserId(
	ctx context.Context,
	request oapi.GetApiV1UserUserIdRequestObject,
) (oapi.GetApiV1UserUserIdResponseObject, error) {
	resp, err := v1.UserDetailAppService(h.app, request)
	return resp, err
}
