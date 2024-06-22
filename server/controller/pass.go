package controller

import "github.com/cjungo/cjungo"

type PassController struct {
}

func NewPassController() *PassController {
	return &PassController{}
}

type PassQueryParam struct {
	Plain *string `json:"plain" validate:"optional" example:"cjun"`
}

func (controller *PassController) Query(ctx cjungo.HttpContext) error {
	param := PassQueryParam{}
	if err := ctx.Bind(&param); err != nil {
		return err
	}

	return ctx.RespOk()
}
