package controller

import (
	"github.com/cjungo/cjuncms/misc"
	"github.com/cjungo/cjungo"
)

type MachineController struct {
	watcher *misc.MachineWatcher
}

func NewMachineController(
	watcher *misc.MachineWatcher,
) *MachineController {
	return &MachineController{
		watcher: watcher,
	}
}

func (controller *MachineController) PeekCpuInfo(ctx cjungo.HttpContext) error {
	return ctx.Resp(controller.watcher.CpuInfo())
}

func (controller *MachineController) PeekCpuTimes(ctx cjungo.HttpContext) error {
	return ctx.Resp(controller.watcher.CpuTimes())
}

func (controller *MachineController) PeekVirtualMemory(ctx cjungo.HttpContext) error {
	return ctx.Resp(controller.watcher.VirtualMemory())
}
