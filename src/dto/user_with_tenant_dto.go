package dto

import (
	"go_chi_template/db/go_chi_template/public/model"
)

type UserWithTenantDto struct {
	Tenant model.Tenant
	User   model.User
}
