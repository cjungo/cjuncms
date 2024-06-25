package controller

import (
	"fmt"
	"time"

	"github.com/cjungo/cjuncms/misc"
	"github.com/cjungo/cjuncms/model"
	"github.com/cjungo/cjungo"
	"github.com/cjungo/cjungo/db"
	"github.com/cjungo/cjungo/ext"
	"gorm.io/gorm"
)

type EmployeeController struct {
	mysql   *db.MySql
	manager *misc.JwtClaimsManager
}

func NewEmployeeController(
	mysql *db.MySql,
	manager *misc.JwtClaimsManager,
) *EmployeeController {
	return &EmployeeController{
		mysql:   mysql,
		manager: manager,
	}
}

type EmployeeAddParam struct {
	Jobnumber string    `json:"jobnumber" form:"jobnumber" validate:"optional" example:"1"`
	Username  string    `json:"username" form:"username" validate:"optional" example:"admin"`
	Password  string    `json:"password" form:"password" validate:"optional" example:"admin"`
	Nickname  string    `json:"nickname" form:"nickname" validate:"optional" example:"管理员"`
	Fullname  string    `json:"fullname" form:"fullname" validate:"optional" example:"管理员"`
	Birthday  time.Time `json:"birthday" form:"birthday" validate:"optional" example:"2000-01-01"`
}

// EmployeeAdd godoc
// @Summary      添加员工
// @Description  添加员工
// @Tags         employee
// @Accept       json
// @Produce      json
// @Router       /employee/add [put]
// @Param request body EmployeeAddParam true "参数"
func (controller *EmployeeController) Add(ctx cjungo.HttpContext) error {
	param := &EmployeeAddParam{}
	if err := ctx.Bind(param); err != nil {
		return ctx.RespBad(err)
	}

	employee := &model.CjEmployee{}
	ext.MoveField(param, employee)

	if err := controller.mysql.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(employee).Error; err != nil {
			return err
		}

		if err := tx.Create(controller.manager.NewOperation(ctx, misc.OPT_ADD, "添加员工")).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return ctx.RespBad(err)
	}

	return ctx.RespOk()
}

type EmployeeEditParam struct {
	ID        uint32    `json:"id" form:"id" validate:"optional" example:"1"`
	Jobnumber string    `json:"jobnumber" form:"jobnumber" validate:"optional" example:"1"`
	Username  string    `json:"username" form:"username" validate:"optional" example:"admin"`
	Nickname  string    `json:"nickname" form:"nickname" validate:"optional" example:"管理员"`
	Fullname  string    `json:"fullname" form:"fullname" validate:"optional" example:"管理员"`
	Birthday  time.Time `json:"birthday" form:"birthday" validate:"optional" example:"2000-01-01"`
}

// EmployeeEdit godoc
// @Summary      修改员工
// @Description  修改员工
// @Tags         employee
// @Accept       json
// @Produce      json
// @Router       /employee/edit [put]
// @Param request body EmployeeAddParam true "参数"
func (controller *EmployeeController) Edit(ctx cjungo.HttpContext) error {
	param := &EmployeeEditParam{}
	if err := ctx.Bind(param); err != nil {
		return ctx.RespBad(err)
	}

	employee := &model.CjEmployee{}
	ext.MoveField(param, employee)

	if err := controller.mysql.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(employee).Error; err != nil {
			return err
		}

		if err := tx.Create(controller.manager.NewOperation(ctx, misc.OPT_EDIT, "修改员工")).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return ctx.RespBad(err)
	}

	return ctx.Resp(employee)
}

type EmployeeQueryParam struct {
	Skip  int     `json:"skip" validate:"optional" example:"0"`
	Take  int     `json:"take" validate:"optional" example:"100"`
	Plain *string `json:"plain" validate:"optional" example:"管理员"`
}

// EmployeeQuery godoc
// @Summary      查询员工
// @Description  查询员工
// @Tags         employee
// @Accept       json
// @Produce      json
// @Router       /employee/edit [get]
// @Param request body EmployeeQueryParam true "参数"
func (controller *EmployeeController) Query(ctx cjungo.HttpContext) error {
	param := &EmployeeQueryParam{}
	if err := ctx.Bind(param); err != nil {
		return ctx.RespBad(err)
	}
	if param.Take == 0 {
		param.Take = 100
	}

	query := controller.mysql.Where("is_removed=?", 0)
	if param.Plain != nil {
		query = query.Where(
			"username=? OR nickname=? OR fullname=? OR jobnumber=?",
			param.Plain,
			param.Plain,
			param.Plain,
			param.Plain,
		)
	}

	var rows []model.CjEmployee
	if err := query.
		Offset(param.Skip).
		Limit(param.Take).
		Find(&rows).
		Error; err != nil {
		return ctx.RespBad(err)
	}

	if err := controller.mysql.Create(
		controller.manager.NewOperation(ctx, misc.OPT_QUERY, "查找员工"),
	).Error; err != nil {
		return ctx.RespBad(err)
	}

	return ctx.Resp(rows)
}

type EmployeeDropParam struct {
	ID  *uint32   `json:"id" validate:"optional" example:"1"`
	IDs *[]uint32 `json:"ids" validate:"optional" example:"1,23,45"`
}

func (controller *EmployeeController) Drop(ctx cjungo.HttpContext) error {
	param := &EmployeeDropParam{}
	if err := ctx.Bind(param); err != nil {
		return ctx.RespBad(err)
	}
	ids := make([]uint32, 0)

	if param.ID != nil {
		ids = append(ids, *param.ID)
	}

	if param.IDs != nil {
		ids = append(ids, *param.IDs...)
	}

	if err := controller.mysql.Transaction(func(tx *gorm.DB) error {
		employees := []model.CjEmployee{}
		if err := tx.Where("id IN ?", ids).Find(&employees).Error; err != nil {
			return err
		}

		if err := tx.Model(&model.CjEmployee{}).Where("id IN ?", ids).Update("is_removed", 1).Error; err != nil {
			return err
		}

		employeeCount := len(employees)
		operations := make([]model.CjOperation, employeeCount)
		for i := 0; i < employeeCount; i++ {
			operations[i] = *controller.manager.NewOperation(ctx, misc.OPT_DROP, fmt.Sprintf("删除员工 %s", employees[i].Username))
		}

		if err := tx.CreateInBatches(operations, 100).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return ctx.RespBad(err)
	}

	return ctx.RespOk()
}
