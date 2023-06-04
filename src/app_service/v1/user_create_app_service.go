package v1

import (
	"go_chi_template/config"
	"go_chi_template/oapi"
	domainservice "go_chi_template/src/domain_service"
	"go_chi_template/src/mutation"
	"go_chi_template/src/repository/single"
	"strconv"
)

func UserCreateAppService(
	app *config.App,
	req oapi.PostApiV1UserRequestObject,
) (oapi.PostApiV1UserResponseObject, error) {
	tenantId, _ := strconv.ParseInt(req.Body.TenantId, 10, 64)
	tenant, _ := single.GetTenantById(app, tenantId)

	user := domainservice.NewUser(tenant, req.Body.Name, req.Body.Email)
	insertedUser, insertErr := mutation.InsertUser(app, user)

	role := "admin"

	userDto := oapi.User{
		Id:    strconv.FormatInt(insertedUser.ID, 10),
		Name:  insertedUser.Name,
		Email: &insertedUser.Email,
		Role:  &role,
	}

	return &oapi.PostApiV1User200JSONResponse{User: userDto}, insertErr
}
