package misc

import (
	"sync"
	"time"

	"github.com/cjungo/cjuncms/model"
	"github.com/cjungo/cjungo"
	"github.com/cjungo/cjungo/db"
	"github.com/cjungo/cjungo/ext"
	"github.com/rs/zerolog"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/mem"
	"github.com/shirou/gopsutil/v4/process"
	"gorm.io/gorm"
)

type MachineStat[T cpu.InfoStat | cpu.TimesStat | mem.VirtualMemoryStat | []disk.UsageStat | []model.CjMachineProcess] struct {
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
	diskUsage     MachineStat[[]disk.UsageStat]
	processes     MachineStat[[]model.CjMachineProcess]
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
		diskUsage: MachineStat[[]disk.UsageStat]{
			stat:  &[]disk.UsageStat{},
			mutex: sync.Mutex{},
		},
		processes: MachineStat[[]model.CjMachineProcess]{
			stat:  &[]model.CjMachineProcess{},
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

func (machine *MachineWatcher) DiskUsage() *[]disk.UsageStat {
	return machine.diskUsage.Read()
}

func (machine *MachineWatcher) Processes() *[]model.CjMachineProcess {
	return machine.processes.Read()
}

func (machine *MachineWatcher) Stat() error {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go machine.cpuInfo.Write(func() (*cpu.InfoStat, error) {
		defer wg.Done()
		cpuInfo, err := cpu.Info()
		if err != nil {
			return nil, err
		}
		return &cpuInfo[0], nil
	})

	wg.Add(1)
	go machine.cpuTimes.Write(func() (*cpu.TimesStat, error) {
		defer wg.Done()
		cpuTimes, err := cpu.Times(false)
		if err != nil {
			return nil, err
		}
		return &cpuTimes[0], nil
	})

	wg.Add(1)
	go machine.virtualMemory.Write(func() (*mem.VirtualMemoryStat, error) {
		defer wg.Done()
		virtualMemory, err := mem.VirtualMemory()
		if err != nil {
			return nil, err
		}
		return virtualMemory, nil
	})

	wg.Add(1)
	go machine.diskUsage.Write(func() (*[]disk.UsageStat, error) {
		defer wg.Done()
		partitionStats, err := disk.Partitions(true)
		if err != nil {
			return nil, err
		}
		result := []disk.UsageStat{}
		for _, partitionStat := range partitionStats {
			usedStat, err := disk.Usage(partitionStat.Mountpoint)
			if err != nil {
				return nil, err
			}
			result = append(result, *usedStat)
		}
		return &result, nil
	})

	wg.Add(1)
	go machine.processes.Write(func() (*[]model.CjMachineProcess, error) {
		defer wg.Done()

		processes, err := process.Processes()
		if err != nil {
			return nil, err
		}

		now := time.Now()
		result := make([]model.CjMachineProcess, len(processes))
		for i, p := range processes {
			result[i].Pid = p.Pid
			result[i].CreateAt = now
			if name, err := p.Name(); err == nil {
				result[i].Name = name
			}
			if username, err := p.Username(); err == nil {
				result[i].Username = username
			}
			if cmdline, err := p.Cmdline(); err == nil {
				result[i].Cmdline = cjungo.LimitStr(cmdline, 120)
			}
			if workdir, err := p.Cwd(); err == nil {
				result[i].Workdir = cjungo.LimitStr(workdir, 120)
			}
			if cpuPercent, err := p.CPUPercent(); err == nil {
				result[i].CPUPercent = cpuPercent
			}
			if memPercent, err := p.MemoryPercent(); err == nil {
				result[i].MemPercent = memPercent
			}
		}

		return &result, nil
	})

	wg.Wait()
	return nil
}

func MachineTick(logger *zerolog.Logger, mysql *db.MySql, machine *MachineWatcher) {
	// 每 4 秒采样到内存，时间尽量短（确保数据实时性）
	go func() {
		if err := ext.Tick(
			4*time.Second,
			func() error {
				return machine.Stat()
			},
		); err != nil {
			logger.Error().Err(err).Msg("[MACHINE]")
		}
	}()

	// 每分钟采样 CPU 内存 磁盘 数据落盘（时间尽量长，不然数据太多）
	go func() {
		if err := ext.Tick(1*time.Minute, func() error {
			now := time.Now()

			// CPU
			// cpuInfo := machine.CpuInfo()
			cpuTimes := machine.CpuTimes()
			machineCpuTimes := &model.CjMachineCPUTime{
				CreateAt: now,
			}
			ext.MoveField(cpuTimes, machineCpuTimes)

			// 内存
			virtualMemory := machine.VirtualMemory()
			machineVirtualMemory := &model.CjMachineVirtualMemory{
				CreateAt: now,
			}
			ext.MoveField(virtualMemory, machineVirtualMemory)

			// 硬盘
			diskUsage := machine.DiskUsage()
			machineDiskUsage := make([]model.CjMachineDiskUsage, len(*diskUsage))
			for i, d := range *diskUsage {
				machineDiskUsage[i].CreateAt = now
				ext.MoveField(&d, &machineDiskUsage[i])
			}

			// logger.Info().
			// 	// Any("cpuInfo", cpuInfo).
			// 	Any("cpuTimes", cpuTimes).
			// 	Any("virtualMemory", virtualMemory).
			// 	Any("diskUsage", diskUsage).
			// 	Any("processes", processes).
			// 	Msg("[MACHINE]")

			return mysql.TransactionSilent(func(tx *gorm.DB) error {
				if err := tx.Create(machineCpuTimes).Error; err != nil {
					return err
				}
				if err := tx.Create(machineVirtualMemory).Error; err != nil {
					return err
				}
				if err := tx.CreateInBatches(&machineDiskUsage, 100).Error; err != nil {
					return err
				}

				return nil
			})
		}); err != nil {
			logger.Error().Err(err).Msg("[MACHINE]")
		}
	}()

	// 每 10 分钟采样 进程 数据落盘（时间尽量长，不然数据太多）
	go func() {
		if err := ext.Tick(10*time.Minute, func() error {
			// 进程
			processes := machine.Processes()
			return mysql.TransactionSilent(func(tx *gorm.DB) error {
				return tx.CreateInBatches(processes, 1000).Error
			})
		}); err != nil {
			logger.Error().Err(err).Msg("[MACHINE]")
		}
	}()
}
