package controller

import (
	"time"

	"github.com/cjungo/cjuncms/model"
	"github.com/cjungo/cjungo"
	"github.com/cjungo/cjungo/db"
	"github.com/cjungo/cjungo/ext"
	"gorm.io/gorm"
)

type SignController struct {
	mysql             *db.MySql
	captchaController *ext.CaptchaController
}

func NewLoginController(
	mysql *db.MySql,
	captchaController *ext.CaptchaController,
) *SignController {
	return &SignController{
		mysql:             mysql,
		captchaController: captchaController,
	}
}

type SignInParam struct {
	Username      string `json:"username" form:"username"`
	Password      string `json:"password" form:"password"`
	CaptchaID     string `json:"captchaId" form:"captchaId" query:"captchaId"`
	CaptchaAnswer string `json:"captchaAnswer" form:"captchaAnswer" query:"captchaAnswer"`
}

// SignIn godoc
// @Summary      登录接口
// @Description  登录接口
// @Tags         sign
// @Accept       json
// @Produce      json
// @Router       /sign/in [post]
// @Param        username   body      string  true  "账号"
// @Param        password   body      string  true  "密码"
// @Param        captchaId   body      string  true  "验证码ID"
// @Param        captchaAnswer   body      string  true  "验证码答案"
func (controller *SignController) SignIn(ctx cjungo.HttpContext) error {
	param := &SignInParam{}
	if err := ctx.Bind(param); err != nil {
		return ctx.RespBad(err)
	}

	if err := controller.captchaController.Verify(param.CaptchaID, param.CaptchaAnswer, true); err != nil {
		return ctx.RespBad(err)
	}

	if err := controller.mysql.Transaction(func(tx *gorm.DB) error {
		employee := &model.CjEmployee{}
		if err := tx.Select("*").
			Where("username=? AND password=UNHEX(?)", param.Username, param.Password).
			Find(&employee).Error; err != nil {
			return err
		}

		if employee.ID == 0 {
			return ctx.RespBad("账号或密码有误")
		}

		if err := tx.Save(&model.CjOperation{
			OperatorID:     employee.ID,
			OperatorType:   0,
			OperateAt:      time.Now(),
			OperateType:    0,
			OperateSummary: "登录",
		}).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return ctx.RespBad(err)
	}

	return ctx.RespOk()
}

func (controller *SignController) SignOut(ctx cjungo.HttpContext) error {
	return ctx.RespOk()
}
