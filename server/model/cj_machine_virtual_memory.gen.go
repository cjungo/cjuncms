// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameCjMachineVirtualMemory = "cj_machine_virtual_memory"

// CjMachineVirtualMemory 内存
type CjMachineVirtualMemory struct {
	ID             uint32    `gorm:"column:id;type:int unsigned;primaryKey" json:"id"`
	Total          uint64    `gorm:"column:total;type:bigint unsigned;not null;comment:全部" json:"total"` // 全部
	Available      uint64    `gorm:"column:available;type:bigint unsigned;not null" json:"available"`
	Used           uint64    `gorm:"column:used;type:bigint unsigned;not null;comment:已用" json:"used"` // 已用
	UsedPercent    float64   `gorm:"column:used_percent;type:double;not null" json:"used_percent"`
	Free           uint64    `gorm:"column:free;type:bigint unsigned;not null;comment:空闲" json:"free"` // 空闲
	Active         uint64    `gorm:"column:active;type:bigint unsigned;not null" json:"active"`
	Inactive       uint64    `gorm:"column:inactive;type:bigint unsigned;not null" json:"inactive"`
	Wired          uint64    `gorm:"column:wired;type:bigint unsigned;not null" json:"wired"`
	Laundry        uint64    `gorm:"column:laundry;type:bigint unsigned;not null" json:"laundry"`
	Buffers        uint64    `gorm:"column:buffers;type:bigint unsigned;not null" json:"buffers"`
	Cached         uint64    `gorm:"column:cached;type:bigint unsigned;not null" json:"cached"`
	WriteBack      uint64    `gorm:"column:write_back;type:bigint unsigned;not null" json:"write_back"`
	Dirty          uint64    `gorm:"column:dirty;type:bigint unsigned;not null" json:"dirty"`
	WriteBackTmp   uint64    `gorm:"column:write_back_tmp;type:bigint unsigned;not null" json:"write_back_tmp"`
	Shared         uint64    `gorm:"column:shared;type:bigint unsigned;not null" json:"shared"`
	Slab           uint64    `gorm:"column:slab;type:bigint unsigned;not null" json:"slab"`
	Sreclaimable   uint64    `gorm:"column:sreclaimable;type:bigint unsigned;not null" json:"sreclaimable"`
	Sunreclaim     uint64    `gorm:"column:sunreclaim;type:bigint unsigned;not null" json:"sunreclaim"`
	PageTables     uint64    `gorm:"column:page_tables;type:bigint unsigned;not null" json:"page_tables"`
	SwapCached     uint64    `gorm:"column:swap_cached;type:bigint unsigned;not null" json:"swap_cached"`
	CommitLimit    uint64    `gorm:"column:commit_limit;type:bigint unsigned;not null" json:"commit_limit"`
	CommittedAs    uint64    `gorm:"column:committed_as;type:bigint unsigned;not null" json:"committed_as"`
	HighTotal      uint64    `gorm:"column:high_total;type:bigint unsigned;not null" json:"high_total"`
	HighFree       uint64    `gorm:"column:high_free;type:bigint unsigned;not null" json:"high_free"`
	LowTotal       uint64    `gorm:"column:low_total;type:bigint unsigned;not null" json:"low_total"`
	LowFree        uint64    `gorm:"column:low_free;type:bigint unsigned;not null" json:"low_free"`
	SwapTotal      uint64    `gorm:"column:swap_total;type:bigint unsigned;not null" json:"swap_total"`
	SwapFree       uint64    `gorm:"column:swap_free;type:bigint unsigned;not null" json:"swap_free"`
	Mapped         uint64    `gorm:"column:mapped;type:bigint unsigned;not null" json:"mapped"`
	VmallocTotal   uint64    `gorm:"column:vmalloc_total;type:bigint unsigned;not null" json:"vmalloc_total"`
	VmallocUsed    uint64    `gorm:"column:vmalloc_used;type:bigint unsigned;not null" json:"vmalloc_used"`
	VmallocChunk   uint64    `gorm:"column:vmalloc_chunk;type:bigint unsigned;not null" json:"vmalloc_chunk"`
	HugePagesTotal uint64    `gorm:"column:huge_pages_total;type:bigint unsigned;not null" json:"huge_pages_total"`
	HugePagesFree  uint64    `gorm:"column:huge_pages_free;type:bigint unsigned;not null" json:"huge_pages_free"`
	HugePagesRsvd  uint64    `gorm:"column:huge_pages_rsvd;type:bigint unsigned;not null" json:"huge_pages_rsvd"`
	HugePagesSurp  uint64    `gorm:"column:huge_pages_surp;type:bigint unsigned;not null" json:"huge_pages_surp"`
	HugePageSize   uint64    `gorm:"column:huge_page_size;type:bigint unsigned;not null" json:"huge_page_size"`
	AnonHugePages  uint64    `gorm:"column:anon_huge_pages;type:bigint unsigned;not null" json:"anon_huge_pages"`
	CreateAt       time.Time `gorm:"column:create_at;type:datetime;not null;index:CREATE_AT_INDEX,priority:1;comment:记录时间" json:"create_at"` // 记录时间
}

// TableName CjMachineVirtualMemory's table name
func (*CjMachineVirtualMemory) TableName() string {
	return TableNameCjMachineVirtualMemory
}
