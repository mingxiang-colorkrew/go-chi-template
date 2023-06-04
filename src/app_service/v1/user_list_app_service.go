package v1

import (
	"go_chi_template/config"
	"go_chi_template/oapi"
	"go_chi_template/src/repository/single"
	"strconv"
)

func UserListAppService(
	app *config.App,
	req oapi.GetApiV1UserRequestObject,
) (oapi.GetApiV1UserResponseObject, error) {
	tenantId, _ := strconv.ParseInt(req.Body.TenantId, 10, 64)

	users, _ := single.GetUsersByTenantId(app, tenantId)

	userDtos := []oapi.User{}

	for _, user := range users {
		userDto := oapi.User{
			Id:    strconv.FormatInt(user.ID, 10),
			Name:  user.Name,
			Email: user.Name,
		}
		userDtos = append(userDtos, userDto)
	}

	return &oapi.GetApiV1User200JSONResponse{
		Users: userDtos,
	}, nil
}
