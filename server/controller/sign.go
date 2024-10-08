package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/cjungo/cjuncms/misc"
	"github.com/cjungo/cjuncms/model"
	"github.com/cjungo/cjungo"
	"github.com/cjungo/cjungo/db"
	"github.com/cjungo/cjungo/ext"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type SignController struct {
	mysql             *db.MySql
	manager           *misc.JwtClaimsManager
	captchaController *ext.CaptchaController
}

func NewSignController(
	mysql *db.MySql,
	manager *misc.JwtClaimsManager,
	captchaController *ext.CaptchaController,
) *SignController {
	return &SignController{
		mysql:             mysql,
		manager:           manager,
		captchaController: captchaController,
	}
}

type SignInParam struct {
	Username      string `json:"username" form:"username" validate:"optional" example:"admin"`
	Password      string `json:"password" form:"password" validate:"optional" example:"admin"`
	CaptchaID     string `json:"captchaId" form:"captchaId" query:"captchaId"`
	CaptchaAnswer string `json:"captchaAnswer" form:"captchaAnswer" query:"captchaAnswer"`
}

type SignInResult struct {
	Token       string            `json:"token" example:"xxx"`
	Permissions []string          `json:"permissions" example:""`
	User        *model.CjEmployee `json:"user" example:""`
}

// SignIn godoc
// @Summary      登录接口
// @Description  登录接口
// @Tags         sign
// @Accept       json
// @Produce      json
// @Router       /sign/in [post]
// @Param        request   body      SignInParam  true  "请求参数"
// @Success      200  {object}   SignInResult
func (controller *SignController) SignIn(ctx cjungo.HttpContext) error {
	param := &SignInParam{}
	if err := ctx.Bind(param); err != nil {
		return ctx.RespBad(err)
	}

	if err := controller.captchaController.Verify(param.CaptchaID, param.CaptchaAnswer, true); err != nil {
		return ctx.RespBad(err)
	}

	result := &SignInResult{}
	if err := controller.mysql.Transaction(func(tx *gorm.DB) error {
		employee := &model.CjEmployee{}
		password := ext.Sha256(param.Password).Hex()
		if err := tx.Select("*").
			Where("username=? AND password=UNHEX(?) AND is_removed=0", param.Username, password).
			Find(&employee).Error; err != nil {
			return err
		}

		if employee.ID == 0 {
			return fmt.Errorf("账号或密码有误")
		} else {
			employee.Password = []byte{}
			result.User = employee
		}

		var permissions []string
		if err := tx.Select("P.tag").
			Table("cj_employee_permission AS EP").
			Joins(
				"JOIN cj_permission AS P ON P.id=EP.permission_id",
			).
			Where("EP.employee_id=?", employee.ID).
			Find(&permissions).Error; err != nil {
			return err
		}

		if err := tx.Create(&model.CjOperation{
			OperatorID:     employee.ID,
			OperatorType:   0,
			OperateAt:      time.Now(),
			OperateType:    0,
			OperateSummary: "登录",
		}).Error; err != nil {
			return err
		}

		claims := &misc.JwtClaims{
			EmployeeToken: misc.EmployeeToken{
				EmployeeId:          employee.ID,
				EmployeeNickname:    cjungo.GetOrDefault(employee.Nickname, ""),
				EmployeePermissions: permissions,
			},
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
				NotBefore: jwt.NewNumericDate(time.Now()),
				Issuer:    "test",
				Subject:   "somebody",
				ID:        "1",
				Audience:  []string{"somebody_else"},
			},
		}

		token, err := ext.MakeJwtToken(claims)
		if err != nil {
			return err
		}

		result.Permissions = permissions
		result.Token = token
		ctx.SetCookie(&http.Cookie{
			Name:     "jwt",
			Value:    token,
			Path:     "/",
			HttpOnly: true,
		})
		return nil
	}); err != nil {
		return ctx.RespBad(err)
	}

	return ctx.Resp(result)
}

func (controller *SignController) SignRenewal(ctx cjungo.HttpContext) error {
	token, err := controller.manager.Renewal(ctx)
	if err != nil {
		return ctx.RespBad(err)
	}
	ctx.SetCookie(&http.Cookie{
		Name:     "jwt",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
	})
	return ctx.Resp(token)
}

func (controller *SignController) SignOut(ctx cjungo.HttpContext) error {
	return ctx.RespOk()
}

func (controller *SignController) GetProfile(ctx cjungo.HttpContext) error {
	profile, err := controller.manager.Profile(ctx)
	if err != nil {
		return err
	}
	return ctx.Resp(profile)
}

func (controller *SignController) SetProfile(ctx cjungo.HttpContext) error {
	profile := misc.EmployeeProfile{}
	if err := ctx.Bind(&profile); err != nil {
		return err
	}
	if self, err := controller.manager.Profile(ctx); err != nil {
		profile.ID = self.ID
	}

	if err := controller.mysql.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(&profile.CjEmployee).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}

	return ctx.RespOk()
}
