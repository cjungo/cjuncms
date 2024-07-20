package controller

import "github.com/cjungo/cjungo"

type ScriptController struct {
}

func NewScriptController() *ScriptController {
	return &ScriptController{}
}

func (controller *ScriptController) Query(ctx cjungo.HttpContext) error {
	return nil
}
