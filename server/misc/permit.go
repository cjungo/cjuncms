package misc

import (
	"github.com/cjungo/cjuncms/model"
	"github.com/cjungo/cjungo/ext"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/elliotchance/pie/v2"
	"gorm.io/gorm"
)

var Permissions []model.CjPermission

func init() {
	Permissions = []model.CjPermission{
		{ID: 10000, ParentID: 0, Tag: "default", Title: "基本"},
		//
		{ID: 11000, ParentID: 0, Tag: "employee", Title: "员工管理", Level: 0},
		{ID: 11001, ParentID: 11000, Tag: "employee_find", Title: "员工查看", Level: 1},
		{ID: 11002, ParentID: 11000, Tag: "employee_edit", Title: "员工修改", Level: 1},
		//
		{ID: 12000, ParentID: 0, Tag: "project", Title: "项目管理", Level: 0},
		{ID: 12001, ParentID: 12000, Tag: "project_find", Title: "项目查看", Level: 1},
		{ID: 12002, ParentID: 12000, Tag: "project_edit", Title: "项目修改", Level: 1},
		{ID: 12003, ParentID: 12000, Tag: "project_script", Title: "执行脚本", Level: 1},
		//
	}
}

func EnsurePermissions(tx *gorm.DB) error {
	permissions := []model.CjPermission{}
	if err := tx.Find(&permissions).Error; err != nil {
		return err
	}
	added, _ := pie.Diff(permissions, Permissions)
	if len(added) > 0 {
		if err := tx.CreateInBatches(added, 100).Error; err != nil {
			return err
		}
	}
	return nil
}

func EnsureAdmin(tx *gorm.DB) error {
	admin := &model.CjEmployee{}
	if err := tx.Find(admin, 1).Error; err != nil {
		return err
	}
	if admin.ID != 1 {
		name := "管理员"
		admin := &model.CjEmployee{
			ID:        1,
			Jobnumber: "1",
			Username:  "admin",
			Password:  ext.Sha256("admin"),
			Nickname:  &name,
			Fullname:  &name,
		}
		if err := tx.Save(admin).Error; err != nil {
			return err
		}
	}
	return EnsureEmployeePermissions(tx, admin)
}

func EnsureEmployeePermissions(tx *gorm.DB, employee *model.CjEmployee) error {
	epids := []int32{}
	if err := tx.Table("cj_employee_permission").Select("permission_id").Where("employee_id=?", employee.ID).Find(&epids).Error; err != nil {
		return err
	}
	epidset := mapset.NewSet(epids...)
	needs := pie.Filter(Permissions, func(p model.CjPermission) bool {
		return !epidset.ContainsOne(int32(p.ID))
	})
	added := make([]model.CjEmployeePermission, len(needs))
	for i, p := range needs {
		added[i].EmployeeID = employee.ID
		added[i].PermissionID = p.ID
	}
	if err := tx.CreateInBatches(added, 100).Error; err != nil {
		return err
	}
	return nil
}
