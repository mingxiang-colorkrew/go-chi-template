package v1

import (
	"measure/config"
	"measure/oapi"
	"measure/src/repository/single"
	"strconv"
)

func ListUserAppService(
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

	return oapi.GetApiV1User200JSONResponse{
		Users: userDtos,
	}, nil

}
