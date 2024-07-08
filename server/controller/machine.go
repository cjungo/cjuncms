package controller

import (
	"time"

	"github.com/cjungo/cjuncms/misc"
	"github.com/cjungo/cjuncms/model"
	"github.com/cjungo/cjungo"
	"github.com/cjungo/cjungo/db"
)

type MachineController struct {
	mysql   *db.MySql
	watcher *misc.MachineWatcher
}

func NewMachineController(
	mysql *db.MySql,
	watcher *misc.MachineWatcher,
) *MachineController {
	return &MachineController{
		mysql:   mysql,
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

type ListCpuTimesParam struct {
	StartAt time.Time `json:"startAt"`
	EndAt   time.Time `json:"endAt"`
}

func (controller *MachineController) ListCpuTimes(ctx cjungo.HttpContext) error {
	param := ListCpuTimesParam{}
	if err := ctx.Bind(&param); err != nil {
		return ctx.RespBad(err)
	}

	rows := []model.CjMachineCPUTime{}
	if err := controller.mysql.
		Where("create_at BETWEEN ? AND ?", param.StartAt, param.EndAt).
		Find(&rows).Error; err != nil {
		return ctx.RespBad(err)
	}

	return ctx.Resp(rows)
}

type ListVirtualMemoryParam struct {
	StartAt time.Time `json:"startAt"`
	EndAt   time.Time `json:"endAt"`
}

func (controller *MachineController) ListVirtualMemory(ctx cjungo.HttpContext) error {
	param := ListVirtualMemoryParam{}
	if err := ctx.Bind(&param); err != nil {
		return ctx.RespBad(err)
	}

	rows := []model.CjMachineVirtualMemory{}
	if err := controller.mysql.
		Where("create_at BETWEEN ? AND ?", param.StartAt, param.EndAt).
		Find(&rows).Error; err != nil {
		return ctx.RespBad(err)
	}

	return ctx.Resp(rows)
}

type ListProcessesParam struct {
	StartAt time.Time `json:"startAt"`
	EndAt   time.Time `json:"endAt"`
}

func (controller *MachineController) ListProcesses(ctx cjungo.HttpContext) error {
	param := ListProcessesParam{}
	if err := ctx.Bind(&param); err != nil {
		return ctx.RespBad(err)
	}

	rows := []model.CjMachineProcess{}
	if err := controller.mysql.
		Where("create_at BETWEEN ? AND ?", param.StartAt, param.EndAt).
		Find(&rows).Error; err != nil {
		return ctx.RespBad(err)
	}

	return ctx.Resp(rows)
}
