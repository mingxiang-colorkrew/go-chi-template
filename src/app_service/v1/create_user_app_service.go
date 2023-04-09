package v1

import (
	"measure/config"
	"measure/oapi"
	domainservice "measure/src/domain_service"
	"measure/src/mutation"
	"measure/src/repository/single"
	"strconv"
)

func CreateUserAppService(
	app *config.App,
	req oapi.PostApiV1UserRequestObject,
) (oapi.PostApiV1UserResponseObject, error) {
	tenantId, _ := strconv.ParseInt(req.Body.TenantId, 10, 64)
	tenant, err := single.GetTenantById(app, tenantId)

	if tenant == nil {
		errDto := createUserErrResp()
		errDto.Errors.TenantId = &[]string{"validation.not_exist"}
		return errDto, err
	}

	existingUser, err := single.GetUserByEmail(app, req.Body.Email)

	if existingUser != nil {
		errDto := createUserErrResp()
		errDto.Errors.Email = &[]string{"validation.unique"}
		return errDto, err
	}

	user := domainservice.NewUser(tenant, req.Body.Name, req.Body.Email)
	insertedUser, err := mutation.InsertUser(app, user)

	role := "admin"

	userDto := oapi.User{
		Id:    strconv.FormatInt(insertedUser.ID, 10),
		Name:  insertedUser.Name,
		Email: &insertedUser.Email,
		Role:  &role,
	}

	return oapi.PostApiV1User200JSONResponse{User: userDto}, nil
}

func createUserErrResp() oapi.PostApiV1User400JSONResponse {
	// copied from oapi.generated.go
	errDto := struct {
		Email    *[]string `json:"email,omitempty"`
		Name     *[]string `json:"name,omitempty"`
		Role     *[]string `json:"role,omitempty"`
		TenantId *[]string `json:"tenantId,omitempty"`
	}{}

	return oapi.PostApiV1User400JSONResponse{Errors: &errDto}
}
