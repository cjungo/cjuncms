package controller

import (
	"github.com/cjungo/cjungo"
	"github.com/cjungo/cjungo/db"
)

type DepartmentController struct {
	mysql *db.MySql
}

func NewDepartmentController(
	mysql *db.MySql,
) *DepartmentController {
	return &DepartmentController{
		mysql: mysql,
	}
}

func (controller *DepartmentController) Query(ctx cjungo.HttpContext) error {
	return nil
}

func (controller *DepartmentController) Add(ctx cjungo.HttpContext) error {
	return nil
}

func (controller *DepartmentController) Edit(ctx cjungo.HttpContext) error {
	return nil
}

func (controller *DepartmentController) Drop(ctx cjungo.HttpContext) error {
	return nil
}
