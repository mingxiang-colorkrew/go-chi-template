package v1

import (
	"measure/config"
	"measure/src/domain_service"
	"measure/src/mutation"
)

func CreateTenantAppService(app *config.App) string {
	newTenant := domainservice.NewTenant("Tenant C", "ck06")
	insertedTenant, err := mutation.InsertTenant(app, newTenant)

	if err != nil {
		return err.Error()
	}

	return insertedTenant.Name
}
