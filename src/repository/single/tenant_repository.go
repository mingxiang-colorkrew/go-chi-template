package single

import (
	"measure/config"
	"measure/db/measure/public/model"
	tbl "measure/db/measure/public/table"

	. "github.com/go-jet/jet/v2/postgres"
)

func GetTenantByShortCode(app *config.App, shortCode string) (*model.Tenant, error) {
	tbl := tbl.Tenant
	stmt := SELECT(tbl.AllColumns).FROM(tbl).WHERE(tbl.ShortCode.EQ(String(shortCode))).LIMIT(1)

	rows := []model.Tenant{}
	err := stmt.Query(app.DB(), &rows)

	if len(rows) != 1 {
		return nil, err
	}

	row := rows[0]
	return &row, nil
}

func GetAllTenants(app *config.App) ([]model.Tenant, error) {
	tbl := tbl.Tenant
	stmt := SELECT(tbl.AllColumns).FROM(tbl)

	rows := []model.Tenant{}
	err := stmt.Query(app.DB(), &rows)

	return rows, err
}

func GetTenantById(app *config.App, id int64) (*model.Tenant, error) {
	tbl := tbl.Tenant
	stmt := SELECT(tbl.AllColumns).FROM(tbl).WHERE(tbl.ID.EQ(Int(id))).LIMIT(1)

	rows := []model.Tenant{}
	err := stmt.Query(app.DB(), &rows)

	if len(rows) != 1 {
		return nil, err
	}

	row := rows[0]
	return &row, err
}
