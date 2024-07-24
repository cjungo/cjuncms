// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/cjungo/cjuncms/model"
)

func newCjProjectEmployee(db *gorm.DB, opts ...gen.DOOption) cjProjectEmployee {
	_cjProjectEmployee := cjProjectEmployee{}

	_cjProjectEmployee.cjProjectEmployeeDo.UseDB(db, opts...)
	_cjProjectEmployee.cjProjectEmployeeDo.UseModel(&model.CjProjectEmployee{})

	tableName := _cjProjectEmployee.cjProjectEmployeeDo.TableName()
	_cjProjectEmployee.ALL = field.NewAsterisk(tableName)
	_cjProjectEmployee.ProjectID = field.NewUint32(tableName, "project_id")
	_cjProjectEmployee.EmployeeID = field.NewUint32(tableName, "employee_id")
	_cjProjectEmployee.CreateAt = field.NewTime(tableName, "create_at")

	_cjProjectEmployee.fillFieldMap()

	return _cjProjectEmployee
}

type cjProjectEmployee struct {
	cjProjectEmployeeDo cjProjectEmployeeDo

	ALL        field.Asterisk
	ProjectID  field.Uint32 // 项目ID
	EmployeeID field.Uint32 // 员工ID
	CreateAt   field.Time

	fieldMap map[string]field.Expr
}

func (c cjProjectEmployee) Table(newTableName string) *cjProjectEmployee {
	c.cjProjectEmployeeDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c cjProjectEmployee) As(alias string) *cjProjectEmployee {
	c.cjProjectEmployeeDo.DO = *(c.cjProjectEmployeeDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *cjProjectEmployee) updateTableName(table string) *cjProjectEmployee {
	c.ALL = field.NewAsterisk(table)
	c.ProjectID = field.NewUint32(table, "project_id")
	c.EmployeeID = field.NewUint32(table, "employee_id")
	c.CreateAt = field.NewTime(table, "create_at")

	c.fillFieldMap()

	return c
}

func (c *cjProjectEmployee) WithContext(ctx context.Context) *cjProjectEmployeeDo {
	return c.cjProjectEmployeeDo.WithContext(ctx)
}

func (c cjProjectEmployee) TableName() string { return c.cjProjectEmployeeDo.TableName() }

func (c cjProjectEmployee) Alias() string { return c.cjProjectEmployeeDo.Alias() }

func (c cjProjectEmployee) Columns(cols ...field.Expr) gen.Columns {
	return c.cjProjectEmployeeDo.Columns(cols...)
}

func (c *cjProjectEmployee) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *cjProjectEmployee) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 3)
	c.fieldMap["project_id"] = c.ProjectID
	c.fieldMap["employee_id"] = c.EmployeeID
	c.fieldMap["create_at"] = c.CreateAt
}

func (c cjProjectEmployee) clone(db *gorm.DB) cjProjectEmployee {
	c.cjProjectEmployeeDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c cjProjectEmployee) replaceDB(db *gorm.DB) cjProjectEmployee {
	c.cjProjectEmployeeDo.ReplaceDB(db)
	return c
}

type cjProjectEmployeeDo struct{ gen.DO }

func (c cjProjectEmployeeDo) Debug() *cjProjectEmployeeDo {
	return c.withDO(c.DO.Debug())
}

func (c cjProjectEmployeeDo) WithContext(ctx context.Context) *cjProjectEmployeeDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c cjProjectEmployeeDo) ReadDB() *cjProjectEmployeeDo {
	return c.Clauses(dbresolver.Read)
}

func (c cjProjectEmployeeDo) WriteDB() *cjProjectEmployeeDo {
	return c.Clauses(dbresolver.Write)
}

func (c cjProjectEmployeeDo) Session(config *gorm.Session) *cjProjectEmployeeDo {
	return c.withDO(c.DO.Session(config))
}

func (c cjProjectEmployeeDo) Clauses(conds ...clause.Expression) *cjProjectEmployeeDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c cjProjectEmployeeDo) Returning(value interface{}, columns ...string) *cjProjectEmployeeDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c cjProjectEmployeeDo) Not(conds ...gen.Condition) *cjProjectEmployeeDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c cjProjectEmployeeDo) Or(conds ...gen.Condition) *cjProjectEmployeeDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c cjProjectEmployeeDo) Select(conds ...field.Expr) *cjProjectEmployeeDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c cjProjectEmployeeDo) Where(conds ...gen.Condition) *cjProjectEmployeeDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c cjProjectEmployeeDo) Order(conds ...field.Expr) *cjProjectEmployeeDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c cjProjectEmployeeDo) Distinct(cols ...field.Expr) *cjProjectEmployeeDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c cjProjectEmployeeDo) Omit(cols ...field.Expr) *cjProjectEmployeeDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c cjProjectEmployeeDo) Join(table schema.Tabler, on ...field.Expr) *cjProjectEmployeeDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c cjProjectEmployeeDo) LeftJoin(table schema.Tabler, on ...field.Expr) *cjProjectEmployeeDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c cjProjectEmployeeDo) RightJoin(table schema.Tabler, on ...field.Expr) *cjProjectEmployeeDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c cjProjectEmployeeDo) Group(cols ...field.Expr) *cjProjectEmployeeDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c cjProjectEmployeeDo) Having(conds ...gen.Condition) *cjProjectEmployeeDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c cjProjectEmployeeDo) Limit(limit int) *cjProjectEmployeeDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c cjProjectEmployeeDo) Offset(offset int) *cjProjectEmployeeDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c cjProjectEmployeeDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *cjProjectEmployeeDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c cjProjectEmployeeDo) Unscoped() *cjProjectEmployeeDo {
	return c.withDO(c.DO.Unscoped())
}

func (c cjProjectEmployeeDo) Create(values ...*model.CjProjectEmployee) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c cjProjectEmployeeDo) CreateInBatches(values []*model.CjProjectEmployee, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c cjProjectEmployeeDo) Save(values ...*model.CjProjectEmployee) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c cjProjectEmployeeDo) First() (*model.CjProjectEmployee, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.CjProjectEmployee), nil
	}
}

func (c cjProjectEmployeeDo) Take() (*model.CjProjectEmployee, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.CjProjectEmployee), nil
	}
}

func (c cjProjectEmployeeDo) Last() (*model.CjProjectEmployee, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.CjProjectEmployee), nil
	}
}

func (c cjProjectEmployeeDo) Find() ([]*model.CjProjectEmployee, error) {
	result, err := c.DO.Find()
	return result.([]*model.CjProjectEmployee), err
}

func (c cjProjectEmployeeDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CjProjectEmployee, err error) {
	buf := make([]*model.CjProjectEmployee, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c cjProjectEmployeeDo) FindInBatches(result *[]*model.CjProjectEmployee, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c cjProjectEmployeeDo) Attrs(attrs ...field.AssignExpr) *cjProjectEmployeeDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c cjProjectEmployeeDo) Assign(attrs ...field.AssignExpr) *cjProjectEmployeeDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c cjProjectEmployeeDo) Joins(fields ...field.RelationField) *cjProjectEmployeeDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c cjProjectEmployeeDo) Preload(fields ...field.RelationField) *cjProjectEmployeeDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c cjProjectEmployeeDo) FirstOrInit() (*model.CjProjectEmployee, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.CjProjectEmployee), nil
	}
}

func (c cjProjectEmployeeDo) FirstOrCreate() (*model.CjProjectEmployee, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.CjProjectEmployee), nil
	}
}

func (c cjProjectEmployeeDo) FindByPage(offset int, limit int) (result []*model.CjProjectEmployee, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c cjProjectEmployeeDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c cjProjectEmployeeDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c cjProjectEmployeeDo) Delete(models ...*model.CjProjectEmployee) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *cjProjectEmployeeDo) withDO(do gen.Dao) *cjProjectEmployeeDo {
	c.DO = *do.(*gen.DO)
	return c
}
