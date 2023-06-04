package v1

import (
	"go_chi_template/config"
	"go_chi_template/oapi"
	"go_chi_template/src/repository/multi"
	"strconv"
)

func UserDetailAppService(
	app *config.App,
	req oapi.GetApiV1UserUserIdRequestObject,
) (oapi.GetApiV1UserUserIdResponseObject, error) {
	userId, _ := strconv.ParseInt(req.UserId, 10, 64)

	row, _ := multi.GetUserWithTenantById(app, int(userId))

	userDto := oapi.User{
		Id:    strconv.FormatInt(row.User.ID, 10),
		Name:  row.User.Name,
		Email: &row.User.Email,
		Tenant: &oapi.Tenant{
			Id:        strconv.FormatInt(row.Tenant.ID, 10),
			Name:      row.Tenant.Name,
			ShortCode: row.Tenant.ShortCode,
		},
	}

	return &oapi.GetApiV1UserUserId200JSONResponse{User: userDto}, nil
}
