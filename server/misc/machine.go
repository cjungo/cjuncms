package misc

import (
	"sync"
	"time"

	"github.com/cjungo/cjuncms/model"
	"github.com/cjungo/cjungo/db"
	"github.com/cjungo/cjungo/ext"
	"github.com/rs/zerolog"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/v4/cpu"
	"gorm.io/gorm"
)

type MachineStat[T cpu.InfoStat | cpu.TimesStat | mem.VirtualMemoryStat] struct {
	stat  *T
	mutex sync.Mutex
}

func (machineStat *MachineStat[T]) Read() *T {
	machineStat.mutex.Lock()
	defer machineStat.mutex.Unlock()
	return machineStat.stat
}

func (machineStat *MachineStat[T]) Write(action func() (*T, error)) error {
	machineStat.mutex.Lock()
	defer machineStat.mutex.Unlock()
	stat, err := action()
	if err != nil {
		return err
	}
	machineStat.stat = stat
	return nil
}

type MachineWatcher struct {
	cpuInfo       MachineStat[cpu.InfoStat]
	cpuTimes      MachineStat[cpu.TimesStat]
	virtualMemory MachineStat[mem.VirtualMemoryStat]
}

func ProvideMachineWatcher() *MachineWatcher {
	return &MachineWatcher{
		cpuInfo: MachineStat[cpu.InfoStat]{
			stat:  &cpu.InfoStat{},
			mutex: sync.Mutex{},
		},
		cpuTimes: MachineStat[cpu.TimesStat]{
			stat:  &cpu.TimesStat{},
			mutex: sync.Mutex{},
		},
		virtualMemory: MachineStat[mem.VirtualMemoryStat]{
			stat:  &mem.VirtualMemoryStat{},
			mutex: sync.Mutex{},
		},
	}
}

func (machine *MachineWatcher) CpuInfo() *cpu.InfoStat {
	return machine.cpuInfo.Read()
}

func (machine *MachineWatcher) CpuTimes() *cpu.TimesStat {
	return machine.cpuTimes.Read()
}

func (machine *MachineWatcher) VirtualMemory() *mem.VirtualMemoryStat {
	return machine.virtualMemory.Read()
}

func (machine *MachineWatcher) Stat() error {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go machine.cpuInfo.Write(func() (*cpu.InfoStat, error) {
		cpuInfo, err := cpu.Info()
		if err != nil {
			return nil, err
		}
		wg.Done()
		return &cpuInfo[0], nil
	})

	wg.Add(1)
	go machine.cpuTimes.Write(func() (*cpu.TimesStat, error) {
		cpuTimes, err := cpu.Times(false)
		if err != nil {
			return nil, err
		}
		wg.Done()
		return &cpuTimes[0], nil
	})

	wg.Add(1)
	go machine.virtualMemory.Write(func() (*mem.VirtualMemoryStat, error) {
		virtualMemory, err := mem.VirtualMemory()
		if err != nil {
			return nil, err
		}
		wg.Done()
		return virtualMemory, nil
	})

	wg.Wait()
	return nil
}

func MachineTick(logger *zerolog.Logger, mysql *db.MySql, machine *MachineWatcher) {
	go func() {
		if err := ext.Tick(
			4*time.Second,
			func() error {
				if err := machine.Stat(); err != nil {
					return err
				}
				// cpuInfo := machine.CpuInfo()
				now := time.Now()
				cpuTimes := machine.CpuTimes()
				machineCpuTimes := &model.CjMachineCPUTime{
					CreateAt: now,
				}
				ext.MoveField(cpuTimes, machineCpuTimes)
	
				virtualMemory := machine.VirtualMemory()
				machineVirtualMemory := &model.CjMachineVirtualMemory{
					CreateAt: now,
				}
				ext.MoveField(virtualMemory, machineVirtualMemory)
	
				// logger.Info().
				// 	Any("cpuInfo", cpuInfo).
				// 	Any("cpuTimes", cpuTimes).
				// 	Any("virtualMemory", virtualMemory).
				// 	Msg("[MACHINE]")
	
				return mysql.Transaction(func(tx *gorm.DB) error {
					if err := tx.Create(machineCpuTimes).Error; err != nil {
						return err
					}
					return tx.Create(machineVirtualMemory).Error
				})
			},
		); err != nil {
			logger.Error().Err(err).Msg("[MACHINE]")
		}
	}()
}
