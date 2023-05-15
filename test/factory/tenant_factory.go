package factory

import (
	"measure/config"
	"measure/db/measure/public/model"
	"measure/src/mutation"
)

func SeedTenant(app *config.App, name string, shortCode string) (*model.Tenant, error) {
	tenant := model.Tenant{
		Name:      name,
		ShortCode: shortCode,
	}

	inserted, err := mutation.InsertTenant(app, &tenant)

	return inserted, err
}
