package controller

import (
	"github.com/cjungo/cjuncms/model"
	"github.com/cjungo/cjungo"
	"github.com/cjungo/cjungo/db"
)

type DepartmentController struct {
	mysql *db.MySql
}

func NewDepartmentController(
	mysql *db.MySql,
) *DepartmentController {
	return &DepartmentController{
		mysql: mysql,
	}
}

type DepartmentNode struct {
	model.CjDepartment
	Children []*DepartmentNode `gorm:"-" json:"children"`
}

func (controller *DepartmentController) Query(ctx cjungo.HttpContext) error {
	departments := []DepartmentNode{}
	if err := controller.mysql.
		Find(&departments).
		Order("level ASC").
		Error; err != nil {
		return err
	}

	tree := []*DepartmentNode{}

	nodeMap := map[uint32]*DepartmentNode{}

	for _, d := range departments {
		nodeMap[d.ID] = &d
		if d.ParentID == 0 {
			tree = append(tree, &d)
		} else {
			parent := nodeMap[d.ParentID]
			parent.Children = append(parent.Children, &d)
		}
	}

	return ctx.Resp(tree)
}

func (controller *DepartmentController) Add(ctx cjungo.HttpContext) error {
	return nil
}

func (controller *DepartmentController) Edit(ctx cjungo.HttpContext) error {
	return nil
}

func (controller *DepartmentController) Drop(ctx cjungo.HttpContext) error {
	return nil
}
