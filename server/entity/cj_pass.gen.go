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

func newCjPass(db *gorm.DB, opts ...gen.DOOption) cjPass {
	_cjPass := cjPass{}

	_cjPass.cjPassDo.UseDB(db, opts...)
	_cjPass.cjPassDo.UseModel(&model.CjPass{})

	tableName := _cjPass.cjPassDo.TableName()
	_cjPass.ALL = field.NewAsterisk(tableName)
	_cjPass.ID = field.NewUint32(tableName, "id")
	_cjPass.Type = field.NewUint32(tableName, "type")
	_cjPass.Title = field.NewString(tableName, "title")
	_cjPass.Host = field.NewString(tableName, "host")
	_cjPass.Port = field.NewUint32(tableName, "port")
	_cjPass.Content = field.NewString(tableName, "content")

	_cjPass.fillFieldMap()

	return _cjPass
}

// cjPass 密钥
type cjPass struct {
	cjPassDo cjPassDo

	ALL     field.Asterisk
	ID      field.Uint32 // ID
	Type    field.Uint32 // 0.密码；1.密钥
	Title   field.String
	Host    field.String // 主机
	Port    field.Uint32 // 端口
	Content field.String // 内容

	fieldMap map[string]field.Expr
}

func (c cjPass) Table(newTableName string) *cjPass {
	c.cjPassDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c cjPass) As(alias string) *cjPass {
	c.cjPassDo.DO = *(c.cjPassDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *cjPass) updateTableName(table string) *cjPass {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewUint32(table, "id")
	c.Type = field.NewUint32(table, "type")
	c.Title = field.NewString(table, "title")
	c.Host = field.NewString(table, "host")
	c.Port = field.NewUint32(table, "port")
	c.Content = field.NewString(table, "content")

	c.fillFieldMap()

	return c
}

func (c *cjPass) WithContext(ctx context.Context) *cjPassDo { return c.cjPassDo.WithContext(ctx) }

func (c cjPass) TableName() string { return c.cjPassDo.TableName() }

func (c cjPass) Alias() string { return c.cjPassDo.Alias() }

func (c cjPass) Columns(cols ...field.Expr) gen.Columns { return c.cjPassDo.Columns(cols...) }

func (c *cjPass) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *cjPass) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 6)
	c.fieldMap["id"] = c.ID
	c.fieldMap["type"] = c.Type
	c.fieldMap["title"] = c.Title
	c.fieldMap["host"] = c.Host
	c.fieldMap["port"] = c.Port
	c.fieldMap["content"] = c.Content
}

func (c cjPass) clone(db *gorm.DB) cjPass {
	c.cjPassDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c cjPass) replaceDB(db *gorm.DB) cjPass {
	c.cjPassDo.ReplaceDB(db)
	return c
}

type cjPassDo struct{ gen.DO }

func (c cjPassDo) Debug() *cjPassDo {
	return c.withDO(c.DO.Debug())
}

func (c cjPassDo) WithContext(ctx context.Context) *cjPassDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c cjPassDo) ReadDB() *cjPassDo {
	return c.Clauses(dbresolver.Read)
}

func (c cjPassDo) WriteDB() *cjPassDo {
	return c.Clauses(dbresolver.Write)
}

func (c cjPassDo) Session(config *gorm.Session) *cjPassDo {
	return c.withDO(c.DO.Session(config))
}

func (c cjPassDo) Clauses(conds ...clause.Expression) *cjPassDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c cjPassDo) Returning(value interface{}, columns ...string) *cjPassDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c cjPassDo) Not(conds ...gen.Condition) *cjPassDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c cjPassDo) Or(conds ...gen.Condition) *cjPassDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c cjPassDo) Select(conds ...field.Expr) *cjPassDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c cjPassDo) Where(conds ...gen.Condition) *cjPassDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c cjPassDo) Order(conds ...field.Expr) *cjPassDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c cjPassDo) Distinct(cols ...field.Expr) *cjPassDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c cjPassDo) Omit(cols ...field.Expr) *cjPassDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c cjPassDo) Join(table schema.Tabler, on ...field.Expr) *cjPassDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c cjPassDo) LeftJoin(table schema.Tabler, on ...field.Expr) *cjPassDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c cjPassDo) RightJoin(table schema.Tabler, on ...field.Expr) *cjPassDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c cjPassDo) Group(cols ...field.Expr) *cjPassDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c cjPassDo) Having(conds ...gen.Condition) *cjPassDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c cjPassDo) Limit(limit int) *cjPassDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c cjPassDo) Offset(offset int) *cjPassDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c cjPassDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *cjPassDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c cjPassDo) Unscoped() *cjPassDo {
	return c.withDO(c.DO.Unscoped())
}

func (c cjPassDo) Create(values ...*model.CjPass) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c cjPassDo) CreateInBatches(values []*model.CjPass, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c cjPassDo) Save(values ...*model.CjPass) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c cjPassDo) First() (*model.CjPass, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.CjPass), nil
	}
}

func (c cjPassDo) Take() (*model.CjPass, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.CjPass), nil
	}
}

func (c cjPassDo) Last() (*model.CjPass, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.CjPass), nil
	}
}

func (c cjPassDo) Find() ([]*model.CjPass, error) {
	result, err := c.DO.Find()
	return result.([]*model.CjPass), err
}

func (c cjPassDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CjPass, err error) {
	buf := make([]*model.CjPass, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c cjPassDo) FindInBatches(result *[]*model.CjPass, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c cjPassDo) Attrs(attrs ...field.AssignExpr) *cjPassDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c cjPassDo) Assign(attrs ...field.AssignExpr) *cjPassDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c cjPassDo) Joins(fields ...field.RelationField) *cjPassDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c cjPassDo) Preload(fields ...field.RelationField) *cjPassDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c cjPassDo) FirstOrInit() (*model.CjPass, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.CjPass), nil
	}
}

func (c cjPassDo) FirstOrCreate() (*model.CjPass, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.CjPass), nil
	}
}

func (c cjPassDo) FindByPage(offset int, limit int) (result []*model.CjPass, count int64, err error) {
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

func (c cjPassDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c cjPassDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c cjPassDo) Delete(models ...*model.CjPass) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *cjPassDo) withDO(do gen.Dao) *cjPassDo {
	c.DO = *do.(*gen.DO)
	return c
}
