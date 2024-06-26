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

func newCjEmployee(db *gorm.DB, opts ...gen.DOOption) cjEmployee {
	_cjEmployee := cjEmployee{}

	_cjEmployee.cjEmployeeDo.UseDB(db, opts...)
	_cjEmployee.cjEmployeeDo.UseModel(&model.CjEmployee{})

	tableName := _cjEmployee.cjEmployeeDo.TableName()
	_cjEmployee.ALL = field.NewAsterisk(tableName)
	_cjEmployee.ID = field.NewUint32(tableName, "id")
	_cjEmployee.Jobnumber = field.NewString(tableName, "jobnumber")
	_cjEmployee.Username = field.NewString(tableName, "username")
	_cjEmployee.Password = field.NewBytes(tableName, "password")
	_cjEmployee.Nickname = field.NewString(tableName, "nickname")
	_cjEmployee.Fullname = field.NewString(tableName, "fullname")
	_cjEmployee.Birthday = field.NewTime(tableName, "birthday")
	_cjEmployee.AvatarPath = field.NewString(tableName, "avatar_path")
	_cjEmployee.IsRemoved = field.NewUint32(tableName, "is_removed")

	_cjEmployee.fillFieldMap()

	return _cjEmployee
}

// cjEmployee 员工
type cjEmployee struct {
	cjEmployeeDo cjEmployeeDo

	ALL        field.Asterisk
	ID         field.Uint32 // ID
	Jobnumber  field.String // 工号
	Username   field.String // 用户名
	Password   field.Bytes  // 密码
	Nickname   field.String // 昵称
	Fullname   field.String // 全称
	Birthday   field.Time   // 生日
	AvatarPath field.String // 头像路径
	IsRemoved  field.Uint32 // 是否删除

	fieldMap map[string]field.Expr
}

func (c cjEmployee) Table(newTableName string) *cjEmployee {
	c.cjEmployeeDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c cjEmployee) As(alias string) *cjEmployee {
	c.cjEmployeeDo.DO = *(c.cjEmployeeDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *cjEmployee) updateTableName(table string) *cjEmployee {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewUint32(table, "id")
	c.Jobnumber = field.NewString(table, "jobnumber")
	c.Username = field.NewString(table, "username")
	c.Password = field.NewBytes(table, "password")
	c.Nickname = field.NewString(table, "nickname")
	c.Fullname = field.NewString(table, "fullname")
	c.Birthday = field.NewTime(table, "birthday")
	c.AvatarPath = field.NewString(table, "avatar_path")
	c.IsRemoved = field.NewUint32(table, "is_removed")

	c.fillFieldMap()

	return c
}

func (c *cjEmployee) WithContext(ctx context.Context) *cjEmployeeDo {
	return c.cjEmployeeDo.WithContext(ctx)
}

func (c cjEmployee) TableName() string { return c.cjEmployeeDo.TableName() }

func (c cjEmployee) Alias() string { return c.cjEmployeeDo.Alias() }

func (c cjEmployee) Columns(cols ...field.Expr) gen.Columns { return c.cjEmployeeDo.Columns(cols...) }

func (c *cjEmployee) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *cjEmployee) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 9)
	c.fieldMap["id"] = c.ID
	c.fieldMap["jobnumber"] = c.Jobnumber
	c.fieldMap["username"] = c.Username
	c.fieldMap["password"] = c.Password
	c.fieldMap["nickname"] = c.Nickname
	c.fieldMap["fullname"] = c.Fullname
	c.fieldMap["birthday"] = c.Birthday
	c.fieldMap["avatar_path"] = c.AvatarPath
	c.fieldMap["is_removed"] = c.IsRemoved
}

func (c cjEmployee) clone(db *gorm.DB) cjEmployee {
	c.cjEmployeeDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c cjEmployee) replaceDB(db *gorm.DB) cjEmployee {
	c.cjEmployeeDo.ReplaceDB(db)
	return c
}

type cjEmployeeDo struct{ gen.DO }

func (c cjEmployeeDo) Debug() *cjEmployeeDo {
	return c.withDO(c.DO.Debug())
}

func (c cjEmployeeDo) WithContext(ctx context.Context) *cjEmployeeDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c cjEmployeeDo) ReadDB() *cjEmployeeDo {
	return c.Clauses(dbresolver.Read)
}

func (c cjEmployeeDo) WriteDB() *cjEmployeeDo {
	return c.Clauses(dbresolver.Write)
}

func (c cjEmployeeDo) Session(config *gorm.Session) *cjEmployeeDo {
	return c.withDO(c.DO.Session(config))
}

func (c cjEmployeeDo) Clauses(conds ...clause.Expression) *cjEmployeeDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c cjEmployeeDo) Returning(value interface{}, columns ...string) *cjEmployeeDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c cjEmployeeDo) Not(conds ...gen.Condition) *cjEmployeeDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c cjEmployeeDo) Or(conds ...gen.Condition) *cjEmployeeDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c cjEmployeeDo) Select(conds ...field.Expr) *cjEmployeeDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c cjEmployeeDo) Where(conds ...gen.Condition) *cjEmployeeDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c cjEmployeeDo) Order(conds ...field.Expr) *cjEmployeeDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c cjEmployeeDo) Distinct(cols ...field.Expr) *cjEmployeeDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c cjEmployeeDo) Omit(cols ...field.Expr) *cjEmployeeDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c cjEmployeeDo) Join(table schema.Tabler, on ...field.Expr) *cjEmployeeDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c cjEmployeeDo) LeftJoin(table schema.Tabler, on ...field.Expr) *cjEmployeeDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c cjEmployeeDo) RightJoin(table schema.Tabler, on ...field.Expr) *cjEmployeeDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c cjEmployeeDo) Group(cols ...field.Expr) *cjEmployeeDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c cjEmployeeDo) Having(conds ...gen.Condition) *cjEmployeeDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c cjEmployeeDo) Limit(limit int) *cjEmployeeDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c cjEmployeeDo) Offset(offset int) *cjEmployeeDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c cjEmployeeDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *cjEmployeeDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c cjEmployeeDo) Unscoped() *cjEmployeeDo {
	return c.withDO(c.DO.Unscoped())
}

func (c cjEmployeeDo) Create(values ...*model.CjEmployee) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c cjEmployeeDo) CreateInBatches(values []*model.CjEmployee, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c cjEmployeeDo) Save(values ...*model.CjEmployee) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c cjEmployeeDo) First() (*model.CjEmployee, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.CjEmployee), nil
	}
}

func (c cjEmployeeDo) Take() (*model.CjEmployee, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.CjEmployee), nil
	}
}

func (c cjEmployeeDo) Last() (*model.CjEmployee, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.CjEmployee), nil
	}
}

func (c cjEmployeeDo) Find() ([]*model.CjEmployee, error) {
	result, err := c.DO.Find()
	return result.([]*model.CjEmployee), err
}

func (c cjEmployeeDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CjEmployee, err error) {
	buf := make([]*model.CjEmployee, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c cjEmployeeDo) FindInBatches(result *[]*model.CjEmployee, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c cjEmployeeDo) Attrs(attrs ...field.AssignExpr) *cjEmployeeDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c cjEmployeeDo) Assign(attrs ...field.AssignExpr) *cjEmployeeDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c cjEmployeeDo) Joins(fields ...field.RelationField) *cjEmployeeDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c cjEmployeeDo) Preload(fields ...field.RelationField) *cjEmployeeDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c cjEmployeeDo) FirstOrInit() (*model.CjEmployee, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.CjEmployee), nil
	}
}

func (c cjEmployeeDo) FirstOrCreate() (*model.CjEmployee, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.CjEmployee), nil
	}
}

func (c cjEmployeeDo) FindByPage(offset int, limit int) (result []*model.CjEmployee, count int64, err error) {
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

func (c cjEmployeeDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c cjEmployeeDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c cjEmployeeDo) Delete(models ...*model.CjEmployee) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *cjEmployeeDo) withDO(do gen.Dao) *cjEmployeeDo {
	c.DO = *do.(*gen.DO)
	return c
}
