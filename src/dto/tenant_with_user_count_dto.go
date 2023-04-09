package dto

import (
	"measure/db/measure/public/model"
)

type TenantWithUserCountDto struct {
	model.Tenant
	UserCount int `alias:"tenant.user_count"`
}
