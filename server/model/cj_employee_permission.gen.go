// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameCjEmployeePermission = "cj_employee_permission"

// CjEmployeePermission mapped from table <cj_employee_permission>
type CjEmployeePermission struct {
	ID           uint32    `gorm:"column:id;type:int unsigned;primaryKey;autoIncrement:true" json:"id"`
	EmployeeID   uint32    `gorm:"column:employee_id;type:int unsigned;not null;uniqueIndex:EMPLOYEE_PERMISSION_UNIQUE,priority:1;comment:员工ID" json:"employee_id"`     // 员工ID
	PermissionID uint32    `gorm:"column:permission_id;type:int unsigned;not null;uniqueIndex:EMPLOYEE_PERMISSION_UNIQUE,priority:2;comment:权限ID" json:"permission_id"` // 权限ID
	CreateAt     time.Time `gorm:"column:create_at;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"create_at"`
}

// TableName CjEmployeePermission's table name
func (*CjEmployeePermission) TableName() string {
	return TableNameCjEmployeePermission
}
