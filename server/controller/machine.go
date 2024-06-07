package controller

import (
	"github.com/cjungo/cjungo"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/mem"
)

type MachineController struct {
}

func NewMachineController() *MachineController {
	return &MachineController{}
}

func (controller *MachineController) PeekCpuInfo(ctx cjungo.HttpContext) error {
	result, err := cpu.Info()
	if err != nil {
		return ctx.RespBad(err)
	}
	return ctx.Resp(result)
}

func (controller *MachineController) PeekCpuTimes(ctx cjungo.HttpContext) error {
	result, err := cpu.Times(false)
	if err != nil {
		return ctx.RespBad(err)
	}
	return ctx.Resp(result)
}

func (controller *MachineController) PeekVirtualMemory(ctx cjungo.HttpContext) error {
	result, err := mem.VirtualMemory()
	if err != nil {
		return ctx.RespBad(err)
	}
	return ctx.Resp(result)
}
