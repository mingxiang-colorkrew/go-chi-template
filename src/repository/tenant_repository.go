package repository

import (
	"database/sql"
	"measure/db/public/model"
	"measure/db/public/table"

	. "github.com/go-jet/jet/v2/postgres"
)

func GetTenantByShortCode(db *sql.DB, shortCode string) (*model.Tenant, error) {
	stmt := SELECT(table.Tenant.AllColumns).FROM(table.Tenant).WHERE(table.Tenant.ShortCode.EQ(String(shortCode))).LIMIT(1)

	rows := []model.Tenant{}
	err := stmt.Query(db, &rows)

	if len(rows) != 1 {
		return nil, err
	}

	row := rows[0]
	return &row, nil
}
