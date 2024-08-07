package misc

import (
	"github.com/cjungo/cjuncms/model"
	"github.com/cjungo/cjungo"
	"github.com/cjungo/cjungo/ext"
	"github.com/cjungo/cjungo/mid"
	mapSet "github.com/deckarep/golang-set/v2"
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
		//
		{ID: 13000, ParentID: 0, Tag: "pass", Title: "密钥管理", Level: 0},
		{ID: 13001, ParentID: 13000, Tag: "pass_find", Title: "密钥查看", Level: 1},
		{ID: 13002, ParentID: 13000, Tag: "pass_edit", Title: "密钥修改", Level: 1},
		//
		{ID: 14000, ParentID: 0, Tag: "shell", Title: "脚本管理", Level: 0},
		{ID: 14001, ParentID: 14000, Tag: "shell_find", Title: "脚本查看", Level: 1},
		{ID: 14002, ParentID: 14000, Tag: "shell_edit", Title: "脚本修改", Level: 1},
		{ID: 14003, ParentID: 14000, Tag: "shell_run", Title: "脚本运行", Level: 1},
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
	eIds := []int32{}
	if err := tx.Table("cj_employee_permission").
		Select("permission_id").
		Where("employee_id=?", employee.ID).
		Find(&eIds).Error; err != nil {
		return err
	}
	eIdSet := mapSet.NewSet(eIds...)
	needs := pie.Filter(Permissions, func(p model.CjPermission) bool {
		return !eIdSet.ContainsOne(int32(p.ID))
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

func ProvidePermitManager() mid.PermitManagerProvide[string, EmployeeToken] {
	return mid.NewPermitManager(func(ctx cjungo.HttpContext) (mid.PermitProof[string, EmployeeToken], error) {
		claims := &JwtClaims{}
		if _, err := ext.ParseJwtToken(ctx, claims); err != nil {
			return nil, &cjungo.ApiError{
				Code:     API_ERR_TOKEN_INVALID,
				Message:  "TOKEN 无效",
				HttpCode: 400,
				Reason:   err,
			}
		}
		return claims, nil
	})
}
