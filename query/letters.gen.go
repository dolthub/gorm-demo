// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/max-hoffman/gorm-dolt/model"
)

func newLetter(db *gorm.DB, opts ...gen.DOOption) letter {
	_letter := letter{}

	_letter.letterDo.UseDB(db, opts...)
	_letter.letterDo.UseModel(&model.Letter{})

	tableName := _letter.letterDo.TableName()
	_letter.ALL = field.NewAsterisk(tableName)
	_letter.Letter = field.NewString(tableName, "letter")
	_letter.CreatedAt = field.NewTime(tableName, "createdAt")
	_letter.CountryID = field.NewString(tableName, "countryID")
	_letter.ID = field.NewString(tableName, "id")

	_letter.fillFieldMap()

	return _letter
}

type letter struct {
	letterDo

	ALL       field.Asterisk
	Letter    field.String
	CreatedAt field.Time
	CountryID field.String
	ID        field.String

	fieldMap map[string]field.Expr
}

func (l letter) Table(newTableName string) *letter {
	l.letterDo.UseTable(newTableName)
	return l.updateTableName(newTableName)
}

func (l letter) As(alias string) *letter {
	l.letterDo.DO = *(l.letterDo.As(alias).(*gen.DO))
	return l.updateTableName(alias)
}

func (l *letter) updateTableName(table string) *letter {
	l.ALL = field.NewAsterisk(table)
	l.Letter = field.NewString(table, "letter")
	l.CreatedAt = field.NewTime(table, "createdAt")
	l.CountryID = field.NewString(table, "countryID")
	l.ID = field.NewString(table, "id")

	l.fillFieldMap()

	return l
}

func (l *letter) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := l.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (l *letter) fillFieldMap() {
	l.fieldMap = make(map[string]field.Expr, 4)
	l.fieldMap["letter"] = l.Letter
	l.fieldMap["createdAt"] = l.CreatedAt
	l.fieldMap["countryID"] = l.CountryID
	l.fieldMap["id"] = l.ID
}

func (l letter) clone(db *gorm.DB) letter {
	l.letterDo.ReplaceConnPool(db.Statement.ConnPool)
	return l
}

func (l letter) replaceDB(db *gorm.DB) letter {
	l.letterDo.ReplaceDB(db)
	return l
}

type letterDo struct{ gen.DO }

type ILetterDo interface {
	gen.SubQuery
	Debug() ILetterDo
	WithContext(ctx context.Context) ILetterDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ILetterDo
	WriteDB() ILetterDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ILetterDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ILetterDo
	Not(conds ...gen.Condition) ILetterDo
	Or(conds ...gen.Condition) ILetterDo
	Select(conds ...field.Expr) ILetterDo
	Where(conds ...gen.Condition) ILetterDo
	Order(conds ...field.Expr) ILetterDo
	Distinct(cols ...field.Expr) ILetterDo
	Omit(cols ...field.Expr) ILetterDo
	Join(table schema.Tabler, on ...field.Expr) ILetterDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ILetterDo
	RightJoin(table schema.Tabler, on ...field.Expr) ILetterDo
	Group(cols ...field.Expr) ILetterDo
	Having(conds ...gen.Condition) ILetterDo
	Limit(limit int) ILetterDo
	Offset(offset int) ILetterDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ILetterDo
	Unscoped() ILetterDo
	Create(values ...*model.Letter) error
	CreateInBatches(values []*model.Letter, batchSize int) error
	Save(values ...*model.Letter) error
	First() (*model.Letter, error)
	Take() (*model.Letter, error)
	Last() (*model.Letter, error)
	Find() ([]*model.Letter, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Letter, err error)
	FindInBatches(result *[]*model.Letter, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Letter) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ILetterDo
	Assign(attrs ...field.AssignExpr) ILetterDo
	Joins(fields ...field.RelationField) ILetterDo
	Preload(fields ...field.RelationField) ILetterDo
	FirstOrInit() (*model.Letter, error)
	FirstOrCreate() (*model.Letter, error)
	FindByPage(offset int, limit int) (result []*model.Letter, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ILetterDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (l letterDo) Debug() ILetterDo {
	return l.withDO(l.DO.Debug())
}

func (l letterDo) WithContext(ctx context.Context) ILetterDo {
	return l.withDO(l.DO.WithContext(ctx))
}

func (l letterDo) ReadDB() ILetterDo {
	return l.Clauses(dbresolver.Read)
}

func (l letterDo) WriteDB() ILetterDo {
	return l.Clauses(dbresolver.Write)
}

func (l letterDo) Session(config *gorm.Session) ILetterDo {
	return l.withDO(l.DO.Session(config))
}

func (l letterDo) Clauses(conds ...clause.Expression) ILetterDo {
	return l.withDO(l.DO.Clauses(conds...))
}

func (l letterDo) Returning(value interface{}, columns ...string) ILetterDo {
	return l.withDO(l.DO.Returning(value, columns...))
}

func (l letterDo) Not(conds ...gen.Condition) ILetterDo {
	return l.withDO(l.DO.Not(conds...))
}

func (l letterDo) Or(conds ...gen.Condition) ILetterDo {
	return l.withDO(l.DO.Or(conds...))
}

func (l letterDo) Select(conds ...field.Expr) ILetterDo {
	return l.withDO(l.DO.Select(conds...))
}

func (l letterDo) Where(conds ...gen.Condition) ILetterDo {
	return l.withDO(l.DO.Where(conds...))
}

func (l letterDo) Order(conds ...field.Expr) ILetterDo {
	return l.withDO(l.DO.Order(conds...))
}

func (l letterDo) Distinct(cols ...field.Expr) ILetterDo {
	return l.withDO(l.DO.Distinct(cols...))
}

func (l letterDo) Omit(cols ...field.Expr) ILetterDo {
	return l.withDO(l.DO.Omit(cols...))
}

func (l letterDo) Join(table schema.Tabler, on ...field.Expr) ILetterDo {
	return l.withDO(l.DO.Join(table, on...))
}

func (l letterDo) LeftJoin(table schema.Tabler, on ...field.Expr) ILetterDo {
	return l.withDO(l.DO.LeftJoin(table, on...))
}

func (l letterDo) RightJoin(table schema.Tabler, on ...field.Expr) ILetterDo {
	return l.withDO(l.DO.RightJoin(table, on...))
}

func (l letterDo) Group(cols ...field.Expr) ILetterDo {
	return l.withDO(l.DO.Group(cols...))
}

func (l letterDo) Having(conds ...gen.Condition) ILetterDo {
	return l.withDO(l.DO.Having(conds...))
}

func (l letterDo) Limit(limit int) ILetterDo {
	return l.withDO(l.DO.Limit(limit))
}

func (l letterDo) Offset(offset int) ILetterDo {
	return l.withDO(l.DO.Offset(offset))
}

func (l letterDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ILetterDo {
	return l.withDO(l.DO.Scopes(funcs...))
}

func (l letterDo) Unscoped() ILetterDo {
	return l.withDO(l.DO.Unscoped())
}

func (l letterDo) Create(values ...*model.Letter) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Create(values)
}

func (l letterDo) CreateInBatches(values []*model.Letter, batchSize int) error {
	return l.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (l letterDo) Save(values ...*model.Letter) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Save(values)
}

func (l letterDo) First() (*model.Letter, error) {
	if result, err := l.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Letter), nil
	}
}

func (l letterDo) Take() (*model.Letter, error) {
	if result, err := l.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Letter), nil
	}
}

func (l letterDo) Last() (*model.Letter, error) {
	if result, err := l.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Letter), nil
	}
}

func (l letterDo) Find() ([]*model.Letter, error) {
	result, err := l.DO.Find()
	return result.([]*model.Letter), err
}

func (l letterDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Letter, err error) {
	buf := make([]*model.Letter, 0, batchSize)
	err = l.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (l letterDo) FindInBatches(result *[]*model.Letter, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return l.DO.FindInBatches(result, batchSize, fc)
}

func (l letterDo) Attrs(attrs ...field.AssignExpr) ILetterDo {
	return l.withDO(l.DO.Attrs(attrs...))
}

func (l letterDo) Assign(attrs ...field.AssignExpr) ILetterDo {
	return l.withDO(l.DO.Assign(attrs...))
}

func (l letterDo) Joins(fields ...field.RelationField) ILetterDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Joins(_f))
	}
	return &l
}

func (l letterDo) Preload(fields ...field.RelationField) ILetterDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Preload(_f))
	}
	return &l
}

func (l letterDo) FirstOrInit() (*model.Letter, error) {
	if result, err := l.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Letter), nil
	}
}

func (l letterDo) FirstOrCreate() (*model.Letter, error) {
	if result, err := l.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Letter), nil
	}
}

func (l letterDo) FindByPage(offset int, limit int) (result []*model.Letter, count int64, err error) {
	result, err = l.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = l.Offset(-1).Limit(-1).Count()
	return
}

func (l letterDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = l.Count()
	if err != nil {
		return
	}

	err = l.Offset(offset).Limit(limit).Scan(result)
	return
}

func (l letterDo) Scan(result interface{}) (err error) {
	return l.DO.Scan(result)
}

func (l letterDo) Delete(models ...*model.Letter) (result gen.ResultInfo, err error) {
	return l.DO.Delete(models)
}

func (l *letterDo) withDO(do gen.Dao) *letterDo {
	l.DO = *do.(*gen.DO)
	return l
}
