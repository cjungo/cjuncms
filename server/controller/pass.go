package controller

import (
	"fmt"

	"github.com/cjungo/cjuncms/model"
	"github.com/cjungo/cjungo"
	"github.com/cjungo/cjungo/db"
	"github.com/cjungo/cjungo/ext"
	"gorm.io/gorm"
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
	Skip  int     `json:"skip" validate:"optional" example:"0"`
	Take  int     `json:"take" validate:"optional" example:"100"`
	Plain *string `json:"plain" validate:"optional" example:"cjun"`
}

func (controller *PassController) Query(ctx cjungo.HttpContext) error {
	param := PassQueryParam{}
	if err := ctx.Bind(&param); err != nil {
		return err
	}
	if param.Take == 0 {
		param.Take = 100
	}

	query := controller.mysql.DB

	if param.Plain != nil {
		query = query.Where("host LIKE ?", fmt.Sprintf("%%%s%%", *param.Plain))
	}

	rows := []model.CjPass{}
	if err := query.
		Offset(param.Skip).
		Limit(param.Take).
		Find(&rows).Error; err != nil {
		return ctx.RespBad(err)
	}

	return ctx.Resp(rows)
}

type PassAddParam struct {
	Type    uint32 `json:"type"`    // 0.密码；1.密钥
	Host    string `json:"host"`    // 主机
	Port    uint32 `json:"port"`    // 端口
	Content string `json:"content"` // 内容
}

func (controller *PassController) Add(ctx cjungo.HttpContext) error {
	param := PassAddParam{}
	if err := ctx.Bind(&param); err != nil {
		return ctx.RespBad(err)
	}
	var pass model.CjPass
	ext.MoveField(&param, &pass)

	if err := controller.mysql.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&pass).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return ctx.RespBad(err)
	}

	return ctx.RespOk()
}

type PassEditParam struct {
	ID      uint32 `json:"id"`      // ID
	Type    uint32 `json:"type"`    // 0.密码；1.密钥
	Host    string `json:"host"`    // 主机
	Port    uint32 `json:"port"`    // 端口
	Content string `json:"content"` // 内容
}

func (controller *PassController) Edit(ctx cjungo.HttpContext) error {
	param := PassEditParam{}
	if err := ctx.Bind(&param); err != nil {
		return ctx.RespBad(err)
	}

	var pass model.CjPass
	ext.MoveField(&param, &pass)
	if err := controller.mysql.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(&pass).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return ctx.RespBad(err)
	}

	return ctx.Resp(&pass)
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
