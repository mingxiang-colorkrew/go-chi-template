package mutation

import (
	"measure/config"
	"measure/db/measure/public/model"
	"measure/db/measure/public/table"
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
