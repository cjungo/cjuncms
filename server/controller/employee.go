package controller

import (
	"time"

	"github.com/cjungo/cjuncms/model"
	"github.com/cjungo/cjungo"
	"github.com/cjungo/cjungo/db"
	"github.com/cjungo/cjungo/ext"
	"gorm.io/gorm"
)

type EmployeeController struct {
	mysql *db.MySql
}

func NewEmployeeController(
	mysql *db.MySql,
) *EmployeeController {
	return &EmployeeController{
		mysql: mysql,
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

		return nil
	}); err != nil {
		return ctx.RespBad(err)
	}

	return ctx.Resp(employee)
}
