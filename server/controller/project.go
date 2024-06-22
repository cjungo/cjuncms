package controller

import (
	"github.com/cjungo/cjungo"
	"github.com/cjungo/cjungo/db"
)

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

type ProjectQueryParam struct {
	Plain *string `json:"plain" validate:"optional" example:"cjun"`
}

func (controller *ProjectController) Query(ctx cjungo.HttpContext) error {
	return ctx.RespOk()
}
