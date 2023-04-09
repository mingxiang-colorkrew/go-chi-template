package dto

import (
	"measure/db/measure/public/model"
)

type UserWithTenantDto struct {
	Tenant model.Tenant
	User   model.User
}
