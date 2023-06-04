package multi

import (
	"go_chi_template/config"
	"go_chi_template/db/go_chi_template/public/table"
	"go_chi_template/src/dto"

	. "github.com/go-jet/jet/v2/postgres"
)

func GetUserWithTenantById(app *config.App, id int) (*dto.UserWithTenantDto, error) {
	tenantTbl := table.Tenant
	userTbl := table.User

	stmt := SELECT(
		userTbl.AllColumns,
		tenantTbl.AllColumns,
	).FROM(tenantTbl.INNER_JOIN(userTbl, userTbl.TenantID.EQ(tenantTbl.ID)))

	rows := []dto.UserWithTenantDto{}
	err := stmt.Query(app.DB(), &rows)

	if len(rows) != 1 {
		return nil, err
	}

	return &rows[0], err
}
