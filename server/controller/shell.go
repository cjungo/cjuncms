package controller

import (
	"github.com/cjungo/cjuncms/model"
	"github.com/cjungo/cjungo"
	"github.com/cjungo/cjungo/db"
)

type ShellController struct {
	mysql *db.MySql
}

func NewShellController(
	mysql *db.MySql,
) *ShellController {
	return &ShellController{
		mysql: mysql,
	}
}

func (controller *ShellController) Query (ctx cjungo.HttpContext) error {
	var rows []model.CjScript

	

	return ctx.Resp(&rows)
}