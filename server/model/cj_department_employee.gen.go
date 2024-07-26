// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameCjDepartmentEmployee = "cj_department_employee"

// CjDepartmentEmployee mapped from table <cj_department_employee>
type CjDepartmentEmployee struct {
	DepartmentID uint32    `gorm:"column:department_id;type:int unsigned;primaryKey" json:"department_id"`
	EmployeeID   uint32    `gorm:"column:employee_id;type:int unsigned;primaryKey" json:"employee_id"`
	PositionID   uint32    `gorm:"column:position_id;type:int unsigned;primaryKey" json:"position_id"`
	CreateAt     time.Time `gorm:"column:create_at;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"create_at"`
}

// TableName CjDepartmentEmployee's table name
func (*CjDepartmentEmployee) TableName() string {
	return TableNameCjDepartmentEmployee
}
