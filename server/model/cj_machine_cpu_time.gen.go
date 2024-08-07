// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameCjMachineCPUTime = "cj_machine_cpu_time"

// CjMachineCPUTime CPU 状态
type CjMachineCPUTime struct {
	ID        uint32    `gorm:"column:id;type:int unsigned;primaryKey" json:"id"`
	CPU       string    `gorm:"column:cpu;type:varchar(45);not null;index:CPU_INDEX,priority:1" json:"cpu"`
	User      float64   `gorm:"column:user;type:double;not null;comment:用户时间" json:"user"`     // 用户时间
	System    float64   `gorm:"column:system;type:double;not null;comment:系统时间" json:"system"` // 系统时间
	Idle      float64   `gorm:"column:idle;type:double;not null;comment:空闲时间" json:"idle"`     // 空闲时间
	Nice      float64   `gorm:"column:nice;type:double;not null" json:"nice"`
	Iowait    float64   `gorm:"column:iowait;type:double;not null;comment:IO等待" json:"iowait"`  // IO等待
	Irq       float64   `gorm:"column:irq;type:double;not null;comment:硬中断" json:"irq"`         // 硬中断
	Softirq   float64   `gorm:"column:softirq;type:double;not null;comment:软中断" json:"softirq"` // 软中断
	Steal     float64   `gorm:"column:steal;type:double;not null" json:"steal"`
	Guest     float64   `gorm:"column:guest;type:double;not null" json:"guest"`
	GuestNice float64   `gorm:"column:guest_nice;type:double;not null" json:"guest_nice"`
	CreateAt  time.Time `gorm:"column:create_at;type:datetime;not null;index:CREATE_AT_INDEX,priority:1;comment:记录时间" json:"create_at"` // 记录时间
}

// TableName CjMachineCPUTime's table name
func (*CjMachineCPUTime) TableName() string {
	return TableNameCjMachineCPUTime
}
