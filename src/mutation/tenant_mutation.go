package mutation

import (
	"measure/config"
	"measure/db/public/model"
	"measure/db/public/table"
)

func InsertTenant(app *config.App, tenant *model.Tenant) (*model.Tenant, error) {
	insertStmt := table.Tenant.INSERT(table.Tenant.MutableColumns).MODEL(tenant)
	dest := []model.Tenant{}

	err := insertStmt.Query(app.DB(), &dest)

	if err != nil {
		return nil, err
	}

	insertedTenant := dest[0]

	return &insertedTenant, nil
}
