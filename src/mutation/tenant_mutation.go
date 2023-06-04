package mutation

import (
	"go_chi_template/config"
	"go_chi_template/db/go_chi_template/public/model"
	"go_chi_template/db/go_chi_template/public/table"
)

func InsertTenant(app *config.App, tenant *model.Tenant) (*model.Tenant, error) {
	insertStmt := table.Tenant.INSERT(table.Tenant.MutableColumns).
		MODEL(tenant).
		RETURNING(table.Tenant.AllColumns)
	dest := []model.Tenant{}

	err := insertStmt.Query(app.DB(), &dest)

	if err != nil || len(dest) != 1 {
		return nil, err
	}

	insertedTenant := dest[0]

	return &insertedTenant, nil
}
