package controller

import "github.com/cjungo/cjungo"

type DemandController struct {
}

func NewDemandController() *DemandController {
	return &DemandController{}
}

func (controller *DemandController) Query(ctx cjungo.HttpContext) error {
	return ctx.RespOk()
}
