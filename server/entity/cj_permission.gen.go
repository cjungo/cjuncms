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

func newCjPermission(db *gorm.DB, opts ...gen.DOOption) cjPermission {
	_cjPermission := cjPermission{}

	_cjPermission.cjPermissionDo.UseDB(db, opts...)
	_cjPermission.cjPermissionDo.UseModel(&model.CjPermission{})

	tableName := _cjPermission.cjPermissionDo.TableName()
	_cjPermission.ALL = field.NewAsterisk(tableName)
	_cjPermission.ID = field.NewUint32(tableName, "id")
	_cjPermission.ParentID = field.NewUint32(tableName, "parent_id")
	_cjPermission.Tag = field.NewString(tableName, "tag")
	_cjPermission.Title = field.NewString(tableName, "title")
	_cjPermission.Level = field.NewUint32(tableName, "level")
	_cjPermission.Description = field.NewString(tableName, "description")

	_cjPermission.fillFieldMap()

	return _cjPermission
}

// cjPermission 权限
type cjPermission struct {
	cjPermissionDo cjPermissionDo

	ALL         field.Asterisk
	ID          field.Uint32
	ParentID    field.Uint32
	Tag         field.String
	Title       field.String
	Level       field.Uint32
	Description field.String

	fieldMap map[string]field.Expr
}

func (c cjPermission) Table(newTableName string) *cjPermission {
	c.cjPermissionDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c cjPermission) As(alias string) *cjPermission {
	c.cjPermissionDo.DO = *(c.cjPermissionDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *cjPermission) updateTableName(table string) *cjPermission {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewUint32(table, "id")
	c.ParentID = field.NewUint32(table, "parent_id")
	c.Tag = field.NewString(table, "tag")
	c.Title = field.NewString(table, "title")
	c.Level = field.NewUint32(table, "level")
	c.Description = field.NewString(table, "description")

	c.fillFieldMap()

	return c
}

func (c *cjPermission) WithContext(ctx context.Context) *cjPermissionDo {
	return c.cjPermissionDo.WithContext(ctx)
}

func (c cjPermission) TableName() string { return c.cjPermissionDo.TableName() }

func (c cjPermission) Alias() string { return c.cjPermissionDo.Alias() }

func (c cjPermission) Columns(cols ...field.Expr) gen.Columns {
	return c.cjPermissionDo.Columns(cols...)
}

func (c *cjPermission) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *cjPermission) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 6)
	c.fieldMap["id"] = c.ID
	c.fieldMap["parent_id"] = c.ParentID
	c.fieldMap["tag"] = c.Tag
	c.fieldMap["title"] = c.Title
	c.fieldMap["level"] = c.Level
	c.fieldMap["description"] = c.Description
}

func (c cjPermission) clone(db *gorm.DB) cjPermission {
	c.cjPermissionDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c cjPermission) replaceDB(db *gorm.DB) cjPermission {
	c.cjPermissionDo.ReplaceDB(db)
	return c
}

type cjPermissionDo struct{ gen.DO }

func (c cjPermissionDo) Debug() *cjPermissionDo {
	return c.withDO(c.DO.Debug())
}

func (c cjPermissionDo) WithContext(ctx context.Context) *cjPermissionDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c cjPermissionDo) ReadDB() *cjPermissionDo {
	return c.Clauses(dbresolver.Read)
}

func (c cjPermissionDo) WriteDB() *cjPermissionDo {
	return c.Clauses(dbresolver.Write)
}

func (c cjPermissionDo) Session(config *gorm.Session) *cjPermissionDo {
	return c.withDO(c.DO.Session(config))
}

func (c cjPermissionDo) Clauses(conds ...clause.Expression) *cjPermissionDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c cjPermissionDo) Returning(value interface{}, columns ...string) *cjPermissionDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c cjPermissionDo) Not(conds ...gen.Condition) *cjPermissionDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c cjPermissionDo) Or(conds ...gen.Condition) *cjPermissionDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c cjPermissionDo) Select(conds ...field.Expr) *cjPermissionDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c cjPermissionDo) Where(conds ...gen.Condition) *cjPermissionDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c cjPermissionDo) Order(conds ...field.Expr) *cjPermissionDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c cjPermissionDo) Distinct(cols ...field.Expr) *cjPermissionDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c cjPermissionDo) Omit(cols ...field.Expr) *cjPermissionDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c cjPermissionDo) Join(table schema.Tabler, on ...field.Expr) *cjPermissionDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c cjPermissionDo) LeftJoin(table schema.Tabler, on ...field.Expr) *cjPermissionDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c cjPermissionDo) RightJoin(table schema.Tabler, on ...field.Expr) *cjPermissionDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c cjPermissionDo) Group(cols ...field.Expr) *cjPermissionDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c cjPermissionDo) Having(conds ...gen.Condition) *cjPermissionDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c cjPermissionDo) Limit(limit int) *cjPermissionDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c cjPermissionDo) Offset(offset int) *cjPermissionDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c cjPermissionDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *cjPermissionDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c cjPermissionDo) Unscoped() *cjPermissionDo {
	return c.withDO(c.DO.Unscoped())
}

func (c cjPermissionDo) Create(values ...*model.CjPermission) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c cjPermissionDo) CreateInBatches(values []*model.CjPermission, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c cjPermissionDo) Save(values ...*model.CjPermission) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c cjPermissionDo) First() (*model.CjPermission, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.CjPermission), nil
	}
}

func (c cjPermissionDo) Take() (*model.CjPermission, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.CjPermission), nil
	}
}

func (c cjPermissionDo) Last() (*model.CjPermission, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.CjPermission), nil
	}
}

func (c cjPermissionDo) Find() ([]*model.CjPermission, error) {
	result, err := c.DO.Find()
	return result.([]*model.CjPermission), err
}

func (c cjPermissionDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CjPermission, err error) {
	buf := make([]*model.CjPermission, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c cjPermissionDo) FindInBatches(result *[]*model.CjPermission, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c cjPermissionDo) Attrs(attrs ...field.AssignExpr) *cjPermissionDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c cjPermissionDo) Assign(attrs ...field.AssignExpr) *cjPermissionDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c cjPermissionDo) Joins(fields ...field.RelationField) *cjPermissionDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c cjPermissionDo) Preload(fields ...field.RelationField) *cjPermissionDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c cjPermissionDo) FirstOrInit() (*model.CjPermission, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.CjPermission), nil
	}
}

func (c cjPermissionDo) FirstOrCreate() (*model.CjPermission, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.CjPermission), nil
	}
}

func (c cjPermissionDo) FindByPage(offset int, limit int) (result []*model.CjPermission, count int64, err error) {
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

func (c cjPermissionDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c cjPermissionDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c cjPermissionDo) Delete(models ...*model.CjPermission) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *cjPermissionDo) withDO(do gen.Dao) *cjPermissionDo {
	c.DO = *do.(*gen.DO)
	return c
}
