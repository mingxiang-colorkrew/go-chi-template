package domainservice

import (
	"measure/db/measure/public/model"
	"time"
)

func NewUser(tenant *model.Tenant, name *string, email string) *model.User {
	newUser := model.User{
		Name:      name,
		Email:     email,
		Role:      10,
		TenantID:  tenant.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return &newUser
}
