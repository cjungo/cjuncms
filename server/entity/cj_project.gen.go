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

func newCjProject(db *gorm.DB, opts ...gen.DOOption) cjProject {
	_cjProject := cjProject{}

	_cjProject.cjProjectDo.UseDB(db, opts...)
	_cjProject.cjProjectDo.UseModel(&model.CjProject{})

	tableName := _cjProject.cjProjectDo.TableName()
	_cjProject.ALL = field.NewAsterisk(tableName)
	_cjProject.ID = field.NewUint32(tableName, "id")
	_cjProject.Name = field.NewString(tableName, "name")

	_cjProject.fillFieldMap()

	return _cjProject
}

// cjProject 项目
type cjProject struct {
	cjProjectDo cjProjectDo

	ALL  field.Asterisk
	ID   field.Uint32
	Name field.String // 项目名

	fieldMap map[string]field.Expr
}

func (c cjProject) Table(newTableName string) *cjProject {
	c.cjProjectDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c cjProject) As(alias string) *cjProject {
	c.cjProjectDo.DO = *(c.cjProjectDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *cjProject) updateTableName(table string) *cjProject {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewUint32(table, "id")
	c.Name = field.NewString(table, "name")

	c.fillFieldMap()

	return c
}

func (c *cjProject) WithContext(ctx context.Context) *cjProjectDo {
	return c.cjProjectDo.WithContext(ctx)
}

func (c cjProject) TableName() string { return c.cjProjectDo.TableName() }

func (c cjProject) Alias() string { return c.cjProjectDo.Alias() }

func (c cjProject) Columns(cols ...field.Expr) gen.Columns { return c.cjProjectDo.Columns(cols...) }

func (c *cjProject) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *cjProject) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 2)
	c.fieldMap["id"] = c.ID
	c.fieldMap["name"] = c.Name
}

func (c cjProject) clone(db *gorm.DB) cjProject {
	c.cjProjectDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c cjProject) replaceDB(db *gorm.DB) cjProject {
	c.cjProjectDo.ReplaceDB(db)
	return c
}

type cjProjectDo struct{ gen.DO }

func (c cjProjectDo) Debug() *cjProjectDo {
	return c.withDO(c.DO.Debug())
}

func (c cjProjectDo) WithContext(ctx context.Context) *cjProjectDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c cjProjectDo) ReadDB() *cjProjectDo {
	return c.Clauses(dbresolver.Read)
}

func (c cjProjectDo) WriteDB() *cjProjectDo {
	return c.Clauses(dbresolver.Write)
}

func (c cjProjectDo) Session(config *gorm.Session) *cjProjectDo {
	return c.withDO(c.DO.Session(config))
}

func (c cjProjectDo) Clauses(conds ...clause.Expression) *cjProjectDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c cjProjectDo) Returning(value interface{}, columns ...string) *cjProjectDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c cjProjectDo) Not(conds ...gen.Condition) *cjProjectDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c cjProjectDo) Or(conds ...gen.Condition) *cjProjectDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c cjProjectDo) Select(conds ...field.Expr) *cjProjectDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c cjProjectDo) Where(conds ...gen.Condition) *cjProjectDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c cjProjectDo) Order(conds ...field.Expr) *cjProjectDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c cjProjectDo) Distinct(cols ...field.Expr) *cjProjectDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c cjProjectDo) Omit(cols ...field.Expr) *cjProjectDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c cjProjectDo) Join(table schema.Tabler, on ...field.Expr) *cjProjectDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c cjProjectDo) LeftJoin(table schema.Tabler, on ...field.Expr) *cjProjectDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c cjProjectDo) RightJoin(table schema.Tabler, on ...field.Expr) *cjProjectDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c cjProjectDo) Group(cols ...field.Expr) *cjProjectDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c cjProjectDo) Having(conds ...gen.Condition) *cjProjectDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c cjProjectDo) Limit(limit int) *cjProjectDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c cjProjectDo) Offset(offset int) *cjProjectDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c cjProjectDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *cjProjectDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c cjProjectDo) Unscoped() *cjProjectDo {
	return c.withDO(c.DO.Unscoped())
}

func (c cjProjectDo) Create(values ...*model.CjProject) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c cjProjectDo) CreateInBatches(values []*model.CjProject, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c cjProjectDo) Save(values ...*model.CjProject) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c cjProjectDo) First() (*model.CjProject, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.CjProject), nil
	}
}

func (c cjProjectDo) Take() (*model.CjProject, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.CjProject), nil
	}
}

func (c cjProjectDo) Last() (*model.CjProject, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.CjProject), nil
	}
}

func (c cjProjectDo) Find() ([]*model.CjProject, error) {
	result, err := c.DO.Find()
	return result.([]*model.CjProject), err
}

func (c cjProjectDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CjProject, err error) {
	buf := make([]*model.CjProject, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c cjProjectDo) FindInBatches(result *[]*model.CjProject, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c cjProjectDo) Attrs(attrs ...field.AssignExpr) *cjProjectDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c cjProjectDo) Assign(attrs ...field.AssignExpr) *cjProjectDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c cjProjectDo) Joins(fields ...field.RelationField) *cjProjectDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c cjProjectDo) Preload(fields ...field.RelationField) *cjProjectDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c cjProjectDo) FirstOrInit() (*model.CjProject, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.CjProject), nil
	}
}

func (c cjProjectDo) FirstOrCreate() (*model.CjProject, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.CjProject), nil
	}
}

func (c cjProjectDo) FindByPage(offset int, limit int) (result []*model.CjProject, count int64, err error) {
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

func (c cjProjectDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c cjProjectDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c cjProjectDo) Delete(models ...*model.CjProject) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *cjProjectDo) withDO(do gen.Dao) *cjProjectDo {
	c.DO = *do.(*gen.DO)
	return c
}