package controller

import "github.com/cjungo/cjungo/db"

type ProjectController struct {
	mysql *db.MySql
}

func NewProjectController(
	mysql *db.MySql,
) *ProjectController {
	return &ProjectController{
		mysql: mysql,
	}
}
