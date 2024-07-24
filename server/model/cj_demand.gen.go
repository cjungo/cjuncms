// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameCjDemand = "cj_demand"

// CjDemand 需求
type CjDemand struct {
	ID       uint32    `gorm:"column:id;type:int unsigned;primaryKey;autoIncrement:true" json:"id"`
	Title    string    `gorm:"column:title;type:varchar(60);not null;comment:标题" json:"title"`  // 标题
	Content  string    `gorm:"column:content;type:longtext;not null;comment:内容" json:"content"` // 内容
	CreateAt time.Time `gorm:"column:create_at;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"create_at"`
}

// TableName CjDemand's table name
func (*CjDemand) TableName() string {
	return TableNameCjDemand
}
