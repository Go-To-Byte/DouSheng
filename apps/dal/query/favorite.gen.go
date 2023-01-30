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

	"github.com/Go-To-Byte/DouSheng/apps/dal/model"
)

func newFavorite(db *gorm.DB, opts ...gen.DOOption) favorite {
	_favorite := favorite{}

	_favorite.favoriteDo.UseDB(db, opts...)
	_favorite.favoriteDo.UseModel(&model.Favorite{})

	tableName := _favorite.favoriteDo.TableName()
	_favorite.ALL = field.NewAsterisk(tableName)
	_favorite.UserID = field.NewInt64(tableName, "user_id")
	_favorite.VideoID = field.NewInt64(tableName, "video_id")
	_favorite.Flag = field.NewInt64(tableName, "flag")

	_favorite.fillFieldMap()

	return _favorite
}

type favorite struct {
	favoriteDo

	ALL     field.Asterisk
	UserID  field.Int64
	VideoID field.Int64
	Flag    field.Int64

	fieldMap map[string]field.Expr
}

func (f favorite) Table(newTableName string) *favorite {
	f.favoriteDo.UseTable(newTableName)
	return f.updateTableName(newTableName)
}

func (f favorite) As(alias string) *favorite {
	f.favoriteDo.DO = *(f.favoriteDo.As(alias).(*gen.DO))
	return f.updateTableName(alias)
}

func (f *favorite) updateTableName(table string) *favorite {
	f.ALL = field.NewAsterisk(table)
	f.UserID = field.NewInt64(table, "user_id")
	f.VideoID = field.NewInt64(table, "video_id")
	f.Flag = field.NewInt64(table, "flag")

	f.fillFieldMap()

	return f
}

func (f *favorite) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := f.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (f *favorite) fillFieldMap() {
	f.fieldMap = make(map[string]field.Expr, 3)
	f.fieldMap["user_id"] = f.UserID
	f.fieldMap["video_id"] = f.VideoID
	f.fieldMap["flag"] = f.Flag
}

func (f favorite) clone(db *gorm.DB) favorite {
	f.favoriteDo.ReplaceConnPool(db.Statement.ConnPool)
	return f
}

func (f favorite) replaceDB(db *gorm.DB) favorite {
	f.favoriteDo.ReplaceDB(db)
	return f
}

type favoriteDo struct{ gen.DO }

type IFavoriteDo interface {
	gen.SubQuery
	Debug() IFavoriteDo
	WithContext(ctx context.Context) IFavoriteDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IFavoriteDo
	WriteDB() IFavoriteDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IFavoriteDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IFavoriteDo
	Not(conds ...gen.Condition) IFavoriteDo
	Or(conds ...gen.Condition) IFavoriteDo
	Select(conds ...field.Expr) IFavoriteDo
	Where(conds ...gen.Condition) IFavoriteDo
	Order(conds ...field.Expr) IFavoriteDo
	Distinct(cols ...field.Expr) IFavoriteDo
	Omit(cols ...field.Expr) IFavoriteDo
	Join(table schema.Tabler, on ...field.Expr) IFavoriteDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IFavoriteDo
	RightJoin(table schema.Tabler, on ...field.Expr) IFavoriteDo
	Group(cols ...field.Expr) IFavoriteDo
	Having(conds ...gen.Condition) IFavoriteDo
	Limit(limit int) IFavoriteDo
	Offset(offset int) IFavoriteDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IFavoriteDo
	Unscoped() IFavoriteDo
	Create(values ...*model.Favorite) error
	CreateInBatches(values []*model.Favorite, batchSize int) error
	Save(values ...*model.Favorite) error
	First() (*model.Favorite, error)
	Take() (*model.Favorite, error)
	Last() (*model.Favorite, error)
	Find() ([]*model.Favorite, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Favorite, err error)
	FindInBatches(result *[]*model.Favorite, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Favorite) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IFavoriteDo
	Assign(attrs ...field.AssignExpr) IFavoriteDo
	Joins(fields ...field.RelationField) IFavoriteDo
	Preload(fields ...field.RelationField) IFavoriteDo
	FirstOrInit() (*model.Favorite, error)
	FirstOrCreate() (*model.Favorite, error)
	FindByPage(offset int, limit int) (result []*model.Favorite, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IFavoriteDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (f favoriteDo) Debug() IFavoriteDo {
	return f.withDO(f.DO.Debug())
}

func (f favoriteDo) WithContext(ctx context.Context) IFavoriteDo {
	return f.withDO(f.DO.WithContext(ctx))
}

func (f favoriteDo) ReadDB() IFavoriteDo {
	return f.Clauses(dbresolver.Read)
}

func (f favoriteDo) WriteDB() IFavoriteDo {
	return f.Clauses(dbresolver.Write)
}

func (f favoriteDo) Session(config *gorm.Session) IFavoriteDo {
	return f.withDO(f.DO.Session(config))
}

func (f favoriteDo) Clauses(conds ...clause.Expression) IFavoriteDo {
	return f.withDO(f.DO.Clauses(conds...))
}

func (f favoriteDo) Returning(value interface{}, columns ...string) IFavoriteDo {
	return f.withDO(f.DO.Returning(value, columns...))
}

func (f favoriteDo) Not(conds ...gen.Condition) IFavoriteDo {
	return f.withDO(f.DO.Not(conds...))
}

func (f favoriteDo) Or(conds ...gen.Condition) IFavoriteDo {
	return f.withDO(f.DO.Or(conds...))
}

func (f favoriteDo) Select(conds ...field.Expr) IFavoriteDo {
	return f.withDO(f.DO.Select(conds...))
}

func (f favoriteDo) Where(conds ...gen.Condition) IFavoriteDo {
	return f.withDO(f.DO.Where(conds...))
}

func (f favoriteDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IFavoriteDo {
	return f.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (f favoriteDo) Order(conds ...field.Expr) IFavoriteDo {
	return f.withDO(f.DO.Order(conds...))
}

func (f favoriteDo) Distinct(cols ...field.Expr) IFavoriteDo {
	return f.withDO(f.DO.Distinct(cols...))
}

func (f favoriteDo) Omit(cols ...field.Expr) IFavoriteDo {
	return f.withDO(f.DO.Omit(cols...))
}

func (f favoriteDo) Join(table schema.Tabler, on ...field.Expr) IFavoriteDo {
	return f.withDO(f.DO.Join(table, on...))
}

func (f favoriteDo) LeftJoin(table schema.Tabler, on ...field.Expr) IFavoriteDo {
	return f.withDO(f.DO.LeftJoin(table, on...))
}

func (f favoriteDo) RightJoin(table schema.Tabler, on ...field.Expr) IFavoriteDo {
	return f.withDO(f.DO.RightJoin(table, on...))
}

func (f favoriteDo) Group(cols ...field.Expr) IFavoriteDo {
	return f.withDO(f.DO.Group(cols...))
}

func (f favoriteDo) Having(conds ...gen.Condition) IFavoriteDo {
	return f.withDO(f.DO.Having(conds...))
}

func (f favoriteDo) Limit(limit int) IFavoriteDo {
	return f.withDO(f.DO.Limit(limit))
}

func (f favoriteDo) Offset(offset int) IFavoriteDo {
	return f.withDO(f.DO.Offset(offset))
}

func (f favoriteDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IFavoriteDo {
	return f.withDO(f.DO.Scopes(funcs...))
}

func (f favoriteDo) Unscoped() IFavoriteDo {
	return f.withDO(f.DO.Unscoped())
}

func (f favoriteDo) Create(values ...*model.Favorite) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Create(values)
}

func (f favoriteDo) CreateInBatches(values []*model.Favorite, batchSize int) error {
	return f.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (f favoriteDo) Save(values ...*model.Favorite) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Save(values)
}

func (f favoriteDo) First() (*model.Favorite, error) {
	if result, err := f.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Favorite), nil
	}
}

func (f favoriteDo) Take() (*model.Favorite, error) {
	if result, err := f.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Favorite), nil
	}
}

func (f favoriteDo) Last() (*model.Favorite, error) {
	if result, err := f.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Favorite), nil
	}
}

func (f favoriteDo) Find() ([]*model.Favorite, error) {
	result, err := f.DO.Find()
	return result.([]*model.Favorite), err
}

func (f favoriteDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Favorite, err error) {
	buf := make([]*model.Favorite, 0, batchSize)
	err = f.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (f favoriteDo) FindInBatches(result *[]*model.Favorite, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return f.DO.FindInBatches(result, batchSize, fc)
}

func (f favoriteDo) Attrs(attrs ...field.AssignExpr) IFavoriteDo {
	return f.withDO(f.DO.Attrs(attrs...))
}

func (f favoriteDo) Assign(attrs ...field.AssignExpr) IFavoriteDo {
	return f.withDO(f.DO.Assign(attrs...))
}

func (f favoriteDo) Joins(fields ...field.RelationField) IFavoriteDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Joins(_f))
	}
	return &f
}

func (f favoriteDo) Preload(fields ...field.RelationField) IFavoriteDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Preload(_f))
	}
	return &f
}

func (f favoriteDo) FirstOrInit() (*model.Favorite, error) {
	if result, err := f.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Favorite), nil
	}
}

func (f favoriteDo) FirstOrCreate() (*model.Favorite, error) {
	if result, err := f.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Favorite), nil
	}
}

func (f favoriteDo) FindByPage(offset int, limit int) (result []*model.Favorite, count int64, err error) {
	result, err = f.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = f.Offset(-1).Limit(-1).Count()
	return
}

func (f favoriteDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = f.Count()
	if err != nil {
		return
	}

	err = f.Offset(offset).Limit(limit).Scan(result)
	return
}

func (f favoriteDo) Scan(result interface{}) (err error) {
	return f.DO.Scan(result)
}

func (f favoriteDo) Delete(models ...*model.Favorite) (result gen.ResultInfo, err error) {
	return f.DO.Delete(models)
}

func (f *favoriteDo) withDO(do gen.Dao) *favoriteDo {
	f.DO = *do.(*gen.DO)
	return f
}
