package single

import (
	"measure/config"
	"measure/db/measure/public/model"
	"measure/db/measure/public/table"

	. "github.com/go-jet/jet/v2/postgres"
)

func GetUserByEmail(app *config.App, email string) (*model.User, error) {
	tbl := table.User
	stmt := SELECT(tbl.AllColumns).FROM(tbl).WHERE(tbl.Email.EQ(String(email))).LIMIT(1)

	rows := []model.User{}
	err := stmt.Query(app.DB(), &rows)

	if len(rows) != 1 {
		return nil, err
	}

	row := rows[0]
	return &row, err
}

func GetUsersByTenantId(app *config.App, tenantId int64) ([]model.User, error) {
	tbl := table.User
	stmt := SELECT(tbl.AllColumns).FROM(tbl).WHERE(tbl.TenantID.EQ(Int(tenantId)))

	rows := []model.User{}
	err := stmt.Query(app.DB(), &rows)

	return rows, err
}
