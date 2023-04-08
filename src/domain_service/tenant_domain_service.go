package domainservice

import (
	"measure/db/measure/public/model"
	"time"
)

func NewTenant(name string, shortCode string) *model.Tenant {
	newTenant := model.Tenant{
		Name:      name,
		ShortCode: shortCode,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return &newTenant
}
