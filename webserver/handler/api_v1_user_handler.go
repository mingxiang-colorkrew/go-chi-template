package handler

import (
	"context"
	"measure/oapi"
	v1 "measure/src/app_service/v1"
)

func (handler *Handler) PostApiV1User(
	ctx context.Context,
	request oapi.PostApiV1UserRequestObject,
) (oapi.PostApiV1UserResponseObject, error) {
	resp, err := v1.CreateUserAppService(handler.app, request)
	return resp, err
}

func (handler *Handler) GetApiV1User(
	ctx context.Context,
	request oapi.GetApiV1UserRequestObject,
) (oapi.GetApiV1UserResponseObject, error) {
	resp, err := v1.ListUserAppService(handler.app, request)
	return resp, err
}

func (handler *Handler) GetApiV1UserUserId(
	ctx context.Context,
	request oapi.GetApiV1UserUserIdRequestObject,
) (oapi.GetApiV1UserUserIdResponseObject, error) {
	resp, err := v1.DetailUserAppService(handler.app, request)
	return resp, err
}
