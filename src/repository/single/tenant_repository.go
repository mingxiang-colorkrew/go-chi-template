package single

import (
	"go_chi_template/config"
	"go_chi_template/db/go_chi_template/public/model"
	tbl "go_chi_template/db/go_chi_template/public/table"

	. "github.com/go-jet/jet/v2/postgres"
)

func GetTenantByName(app *config.App, name string) (*model.Tenant, error) {
	tbl := tbl.Tenant
	stmt := SELECT(tbl.AllColumns).FROM(tbl).WHERE(tbl.Name.EQ(String(name))).LIMIT(1)

	rows := []model.Tenant{}
	err := stmt.Query(app.DB(), &rows)

	if len(rows) != 1 {
		return nil, err
	}

	row := rows[0]
	return &row, nil
}

func GetTenantByShortCode(app *config.App, shortCode string) (*model.Tenant, error) {
	tbl := tbl.Tenant
	stmt := SELECT(
		tbl.AllColumns,
	).FROM(tbl).
		WHERE(LOWER(tbl.ShortCode).EQ(LOWER(String(shortCode)))).
		LIMIT(1)

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
