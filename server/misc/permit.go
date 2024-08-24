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

type PermitTag = string

const (
	PERMIT_DEFAULT PermitTag = "default"

	PERMIT_EMPLOYEE                 PermitTag = "employee"
	PERMIT_EMPLOYEE_EDIT            PermitTag = "employee_edit"
	PERMIT_EMPLOYEE_DEPARTMENT_EDIT PermitTag = "employee_department_edit"

	PERMIT_PROJECT      PermitTag = "project"
	PERMIT_PROJECT_EDIT PermitTag = "project_edit"

	PERMIT_PASS      PermitTag = "pass"
	PERMIT_PASS_EDIT PermitTag = "pass_edit"

	PERMIT_SHELL      PermitTag = "shell"
	PERMIT_SHELL_EDIT PermitTag = "shell_edit"
	PERMIT_SHELL_RUN  PermitTag = "shell_run"
)

var Permissions []model.CjPermission

func init() {
	Permissions = []model.CjPermission{
		{ID: 10000, ParentID: 0, Tag: PERMIT_DEFAULT, Title: "基本信息"},

		//
		{ID: 11000, ParentID: 0, Tag: PERMIT_EMPLOYEE, Title: "员工信息", Level: 0},
		{ID: 11001, ParentID: 11000, Tag: PERMIT_EMPLOYEE_EDIT, Title: "员工管理", Level: 1},
		{ID: 11002, ParentID: 11000, Tag: PERMIT_EMPLOYEE_DEPARTMENT_EDIT, Title: "部门管理", Level: 1},

		//
		{ID: 12000, ParentID: 0, Tag: PERMIT_PROJECT, Title: "项目信息", Level: 0},
		{ID: 12001, ParentID: 12000, Tag: PERMIT_PROJECT_EDIT, Title: "项目管理", Level: 1},

		//
		{ID: 13000, ParentID: 0, Tag: PERMIT_PASS, Title: "密钥信息", Level: 0},
		{ID: 13001, ParentID: 13000, Tag: PERMIT_PASS_EDIT, Title: "密钥管理", Level: 1},

		//
		{ID: 14000, ParentID: 0, Tag: PERMIT_SHELL, Title: "脚本信息", Level: 0},
		{ID: 14001, ParentID: 14000, Tag: PERMIT_SHELL_EDIT, Title: "脚本管理", Level: 1},
		{ID: 14002, ParentID: 14000, Tag: PERMIT_SHELL_RUN, Title: "脚本运行", Level: 1},
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
