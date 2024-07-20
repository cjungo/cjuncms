package controller

import (
	"fmt"

	"github.com/cjungo/cjuncms/model"
	"github.com/cjungo/cjungo"
	"github.com/cjungo/cjungo/db"
	"gorm.io/gorm"
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
	Skip  int     `json:"skip" example:"0"`
	Take  int     `json:"take" example:"100"`
	Plain *string `json:"plain" validate:"optional" example:"cjun"`
}

func (controller *ProjectController) Query(ctx cjungo.HttpContext) error {
	param := ProjectQueryParam{}
	if err := ctx.Bind(&param); err != nil {
		return ctx.RespBad(err)
	}
	query := controller.mysql.DB
	if param.Plain != nil {
		plain := fmt.Sprintf("%%%s%%", *param.Plain)
		query = query.Where("shortname LIKE ? OR fullname LIKE ? OR number LIKE ?", plain, plain, plain)
	}

	take := cjungo.Max(param.Take, 40)
	rows := []model.CjProject{}
	if err := query.
		Offset(param.Skip).
		Limit(take).
		Find(&rows).Error; err != nil {
		return ctx.RespBad(err)
	}

	return ctx.Resp(rows)
}

type ProjectAddParam struct {
	Info        *model.CjProject `json:"info"`
	EmployeeIds []uint32         `json:"employeeIds"`
}

func (controller *ProjectController) Add(ctx cjungo.HttpContext) error {
	row := ProjectAddParam{}
	if err := ctx.Bind(&row); err != nil {
		return ctx.RespBad(err)
	}

	if err := controller.mysql.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&row.Info).Error; err != nil {
			return err
		}
		pes := []model.CjProjectEmployee{}
		if err := tx.CreateInBatches(&pes, 100).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}

	return ctx.Resp(row)
}

func (controller *ProjectController) Edit(ctx cjungo.HttpContext) error {
	if err := controller.mysql.Transaction(func(tx *gorm.DB) error {
		return nil
	}); err != nil {
		return ctx.RespBad(err)
	}
	return ctx.RespOk()
}

type ProjectDropParam struct {
	ID  *uint32   `json:"id" validate:"optional" example:"1"`
	IDs *[]uint32 `json:"ids" validate:"optional" example:"1,23,45"`
}

func (controller *ProjectController) Drop(ctx cjungo.HttpContext) error {
	param := ProjectDropParam{}
	if err := ctx.Bind(&param); err != nil {
		return ctx.RespBad(err)
	}
	if err := controller.mysql.Transaction(func(tx *gorm.DB) error {
		ids := []uint32{}
		if param.ID != nil {
			ids = append(ids, *param.ID)
		}
		if param.IDs != nil {
			ids = append(ids, *param.IDs...)
		}
		if err := tx.Delete(&model.CjProject{}, ids).Error; err != nil {
			return err
		}

		if err := tx.Where("project_id IN ?", ids).
			Delete(&model.CjProjectEmployee{}).
			Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return ctx.RespBad(err)
	}
	return ctx.RespOk()
}
