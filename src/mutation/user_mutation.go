package mutation

import (
	"go_chi_template/config"
	"go_chi_template/db/go_chi_template/public/model"
	"go_chi_template/db/go_chi_template/public/table"
)

func InsertUser(app *config.App, user *model.User) (*model.User, error) {
	insertStmt := table.User.INSERT(table.User.MutableColumns).
		MODEL(user).
		RETURNING(table.User.AllColumns)

	dest := []model.User{}

	err := insertStmt.Query(app.DB(), &dest)

	if err != nil || len(dest) != 1 {
		return nil, err
	}

	insertedUser := dest[0]

	return &insertedUser, nil
}
