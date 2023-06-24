package v1

import (
	"context"
	"go_chi_template/config"
	"go_chi_template/oapi"
	"go_chi_template/src/repository/single"
	"strconv"
)

func DepartmentListAppService(
	ctx context.Context,
	app *config.App,
	_ oapi.GetApiV1DepartmentRequestObject,
) (oapi.GetApiV1DepartmentResponseObject, error) {
	depts, error := single.GetDepartments(app)

	if error != nil {
		app.Logger().Sugar().Error(error)
	}

	deptDtos := []oapi.Department{}

	for _, dept := range depts {
		hierarchy := []oapi.Department{}

		for _, ancestor := range dept.Ancestors {
			ancestorDto := oapi.Department{
				Id:        strconv.FormatInt(ancestor.ID, 10),
				Name:      *ancestor.Name,
				CustomId:  ancestor.CustomID,
				Hierarchy: []oapi.Department{},
			}
			hierarchy = append(hierarchy, ancestorDto)
		}

		deptDto := oapi.Department{
			Id:        strconv.FormatInt(dept.ID, 10),
			Name:      *dept.Name,
			CustomId:  dept.CustomID,
			Hierarchy: hierarchy,
		}

		deptDtos = append(deptDtos, deptDto)
	}

	return &oapi.GetApiV1Department200JSONResponse{
		Departments: deptDtos,
	}, nil
}
