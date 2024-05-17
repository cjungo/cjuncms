package controller

import (
	"time"

	"github.com/cjungo/cjuncms/model"
	"github.com/cjungo/cjungo"
	"github.com/cjungo/cjungo/ext"
)

type EmployeeController struct {
}

func NewEmployeeController() *EmployeeController {
	return &EmployeeController{}
}

type EmployeeAddParam struct {
	Jobnumber string    `json:"jobnumber"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Nickname  string    `json:"nickname"`
	Fullname  string    `json:"fullname"`
	Birthday  time.Time `json:"birthday"`
}

// SignIn godoc
// @Summary      添加员工
// @Description  添加员工
// @Tags         sign
// @Accept       json
// @Produce      json
// @Router       /employee/add [put]
// @Param jobnumber string true "工号"
// @Param username  string true "账号"
// @Param password  []byte  true "密码"
// @Param nickname  *string false "昵称"
// @Param fullname  *string false "全称"
// @Param birthday  *time.Time false "生日"
func (controller *EmployeeController) Add(ctx cjungo.HttpContext) error {
	param := &EmployeeAddParam{}
	if err := ctx.Bind(param); err != nil {

	}

	employee := &model.CjEmployee{}
	ext.MoveField(param, employee)


	return ctx.RespOk()
}
