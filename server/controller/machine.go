package controller

import (
	"time"

	"github.com/cjungo/cjuncms/misc"
	"github.com/cjungo/cjuncms/model"
	"github.com/cjungo/cjungo"
	"github.com/cjungo/cjungo/db"
	"github.com/rs/zerolog"
)

type MachineController struct {
	mysql   *db.MySql
	watcher *misc.MachineWatcher
	logger  *zerolog.Logger
}

func NewMachineController(
	mysql *db.MySql,
	watcher *misc.MachineWatcher,
	logger *zerolog.Logger,
) *MachineController {
	return &MachineController{
		mysql:   mysql,
		watcher: watcher,
		logger:  logger,
	}
}

func (controller *MachineController) PeekCpuInfo(ctx cjungo.HttpContext) error {
	return ctx.Resp(controller.watcher.CpuInfo())
}

func (controller *MachineController) WatchCpuInfo(
	ctx cjungo.HttpContext,
	tx chan cjungo.SseEvent,
	rx chan error,
) {
	for {
		select {
		case <-ctx.Request().Context().Done():
			return
		case err := <-rx:
			tx <- cjungo.SseEvent{Data: err}
		default:
			tx <- cjungo.SseEvent{
				Data: controller.watcher.CpuInfo(),
			}
			time.Sleep(4 * time.Second)
		}
	}
}

func (controller *MachineController) PeekCpuTimes(ctx cjungo.HttpContext) error {
	return ctx.Resp(controller.watcher.CpuTimes())
}

func (controller *MachineController) WatchCpuTimes(
	ctx cjungo.HttpContext,
	tx chan cjungo.SseEvent,
	rx chan error,
) {
	for {
		select {
		case <-ctx.Request().Context().Done():
			return
		case err := <-rx:
			tx <- cjungo.SseEvent{Data: err}
		default:
			tx <- cjungo.SseEvent{
				Data: controller.watcher.CpuTimes(),
			}
			time.Sleep(4 * time.Second)
		}
	}
}

func (controller *MachineController) PeekVirtualMemory(ctx cjungo.HttpContext) error {
	return ctx.Resp(controller.watcher.VirtualMemory())
}

func (controller *MachineController) WatchVirtualMemory(
	ctx cjungo.HttpContext,
	tx chan cjungo.SseEvent,
	rx chan error,
) {
	for {
		select {
		case <-ctx.Request().Context().Done():
			return
		case err := <-rx:
			tx <- cjungo.SseEvent{Data: err}
		default:
			tx <- cjungo.SseEvent{
				Data: controller.watcher.VirtualMemory(),
			}
			time.Sleep(4 * time.Second)
		}
	}
}

func (controller *MachineController) PeekDiskUsage(ctx cjungo.HttpContext) error {
	return ctx.Resp(controller.watcher.DiskUsage())
}

func (controller *MachineController) WatchDiskUsage(
	ctx cjungo.HttpContext,
	tx chan cjungo.SseEvent,
	rx chan error,
) {
	for {
		select {
		case <-ctx.Request().Context().Done():
			return
		case err := <-rx:
			tx <- cjungo.SseEvent{Data: err}
		default:
			tx <- cjungo.SseEvent{
				Data: controller.watcher.DiskUsage(),
			}
			time.Sleep(4 * time.Second)
		}
	}
}

func (controller *MachineController) PeekProcesses(ctx cjungo.HttpContext) error {
	return ctx.Resp(controller.watcher.Processes())
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
	controller.logger.Info().
		Time("startAt", param.StartAt).
		Time("endAt", param.EndAt).
		Msg("[MACHINE]")

	rows := []model.CjMachineCPUTime{}
	if err := controller.mysql.
		Where("create_at BETWEEN ? AND ?", param.StartAt, param.EndAt).
		Find(&rows).Error; err != nil {
		return ctx.RespBad(err)
	}

	return ctx.Resp(rows)
}

func (controller *MachineController) WatchCpuTimesTimeline(
	ctx cjungo.HttpContext,
	tx chan cjungo.SseEvent,
	rx chan error,
) {
	for {
		select {
		case <-ctx.Request().Context().Done():
			return
		case err := <-rx:
			tx <- cjungo.SseEvent{Data: err}
		default:
			rows := []model.CjMachineCPUTime{}
			endAt := time.Now()
			startAt := endAt.Add(-24 * time.Hour)
			if err := controller.mysql.
				Where("create_at BETWEEN ? AND ?", startAt, endAt).
				Find(&rows).Error; err != nil {
				tx <- cjungo.SseEvent{Data: err}
				return
			}
			tx <- cjungo.SseEvent{
				Data: rows,
			}
			time.Sleep(4 * time.Second)
		}
	}
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
