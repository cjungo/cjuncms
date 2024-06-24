package controller

import (
	"fmt"

	"github.com/cjungo/cjuncms/model"
	"github.com/cjungo/cjungo"
	"github.com/cjungo/cjungo/db"
)

type PassController struct {
	mysql *db.MySql
}

func NewPassController(
	mysql *db.MySql,
) *PassController {
	return &PassController{
		mysql: mysql,
	}
}

type PassQueryParam struct {
	Plain *string `json:"plain" validate:"optional" example:"cjun"`
}

func (controller *PassController) Query(ctx cjungo.HttpContext) error {
	param := PassQueryParam{}
	if err := ctx.Bind(&param); err != nil {
		return err
	}

	query := controller.mysql.DB

	if param.Plain != nil {
		query = query.Where("host LIKE ?", fmt.Sprintf("%%%s%%", *param.Plain))
	}

	rows := []model.CjPass{}
	if err := query.Find(&rows).Error; err != nil {
		return ctx.RespBad(err)
	}

	return ctx.Resp(rows)
}

type PassDropParam struct {
	ID  *uint32   `json:"id" validate:"optional" example:"1"`
	IDs *[]uint32 `json:"ids" validate:"optional" example:"1,2,3,4,5"`
}

func (controller *PassController) Drop(ctx cjungo.HttpContext) error {
	param := PassDropParam{}
	if err := ctx.Bind(&param); err != nil {
		return ctx.RespBad(err)
	}
	return ctx.RespOk()
}
