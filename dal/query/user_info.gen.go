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

	"DouSheng/dal/model"
)

func newUserInfo(db *gorm.DB, opts ...gen.DOOption) userInfo {
	_userInfo := userInfo{}

	_userInfo.userInfoDo.UseDB(db, opts...)
	_userInfo.userInfoDo.UseModel(&model.UserInfo{})

	tableName := _userInfo.userInfoDo.TableName()
	_userInfo.ALL = field.NewAsterisk(tableName)
	_userInfo.Username = field.NewString(tableName, "username")
	_userInfo.User = field.NewInt64(tableName, "user")
	_userInfo.Passwd = field.NewString(tableName, "passwd")
	_userInfo.UserID = field.NewInt64(tableName, "user_id")
	_userInfo.Token = field.NewString(tableName, "token")
	_userInfo.TokenTime = field.NewInt64(tableName, "token_time")

	_userInfo.fillFieldMap()

	return _userInfo
}

type userInfo struct {
	userInfoDo userInfoDo

	ALL       field.Asterisk
	Username  field.String
	User      field.Int64
	Passwd    field.String
	UserID    field.Int64
	Token     field.String
	TokenTime field.Int64

	fieldMap map[string]field.Expr
}

func (u userInfo) Table(newTableName string) *userInfo {
	u.userInfoDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u userInfo) As(alias string) *userInfo {
	u.userInfoDo.DO = *(u.userInfoDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *userInfo) updateTableName(table string) *userInfo {
	u.ALL = field.NewAsterisk(table)
	u.Username = field.NewString(table, "username")
	u.User = field.NewInt64(table, "user")
	u.Passwd = field.NewString(table, "passwd")
	u.UserID = field.NewInt64(table, "user_id")
	u.Token = field.NewString(table, "token")
	u.TokenTime = field.NewInt64(table, "token_time")

	u.fillFieldMap()

	return u
}

func (u *userInfo) WithContext(ctx context.Context) IUserInfoDo { return u.userInfoDo.WithContext(ctx) }

func (u userInfo) TableName() string { return u.userInfoDo.TableName() }

func (u userInfo) Alias() string { return u.userInfoDo.Alias() }

func (u *userInfo) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *userInfo) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 6)
	u.fieldMap["username"] = u.Username
	u.fieldMap["user"] = u.User
	u.fieldMap["passwd"] = u.Passwd
	u.fieldMap["user_id"] = u.UserID
	u.fieldMap["token"] = u.Token
	u.fieldMap["token_time"] = u.TokenTime
}

func (u userInfo) clone(db *gorm.DB) userInfo {
	u.userInfoDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u userInfo) replaceDB(db *gorm.DB) userInfo {
	u.userInfoDo.ReplaceDB(db)
	return u
}

type userInfoDo struct{ gen.DO }

type IUserInfoDo interface {
	gen.SubQuery
	Debug() IUserInfoDo
	WithContext(ctx context.Context) IUserInfoDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IUserInfoDo
	WriteDB() IUserInfoDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IUserInfoDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IUserInfoDo
	Not(conds ...gen.Condition) IUserInfoDo
	Or(conds ...gen.Condition) IUserInfoDo
	Select(conds ...field.Expr) IUserInfoDo
	Where(conds ...gen.Condition) IUserInfoDo
	Order(conds ...field.Expr) IUserInfoDo
	Distinct(cols ...field.Expr) IUserInfoDo
	Omit(cols ...field.Expr) IUserInfoDo
	Join(table schema.Tabler, on ...field.Expr) IUserInfoDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IUserInfoDo
	RightJoin(table schema.Tabler, on ...field.Expr) IUserInfoDo
	Group(cols ...field.Expr) IUserInfoDo
	Having(conds ...gen.Condition) IUserInfoDo
	Limit(limit int) IUserInfoDo
	Offset(offset int) IUserInfoDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IUserInfoDo
	Unscoped() IUserInfoDo
	Create(values ...*model.UserInfo) error
	CreateInBatches(values []*model.UserInfo, batchSize int) error
	Save(values ...*model.UserInfo) error
	First() (*model.UserInfo, error)
	Take() (*model.UserInfo, error)
	Last() (*model.UserInfo, error)
	Find() ([]*model.UserInfo, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserInfo, err error)
	FindInBatches(result *[]*model.UserInfo, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.UserInfo) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IUserInfoDo
	Assign(attrs ...field.AssignExpr) IUserInfoDo
	Joins(fields ...field.RelationField) IUserInfoDo
	Preload(fields ...field.RelationField) IUserInfoDo
	FirstOrInit() (*model.UserInfo, error)
	FirstOrCreate() (*model.UserInfo, error)
	FindByPage(offset int, limit int) (result []*model.UserInfo, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IUserInfoDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (u userInfoDo) Debug() IUserInfoDo {
	return u.withDO(u.DO.Debug())
}

func (u userInfoDo) WithContext(ctx context.Context) IUserInfoDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userInfoDo) ReadDB() IUserInfoDo {
	return u.Clauses(dbresolver.Read)
}

func (u userInfoDo) WriteDB() IUserInfoDo {
	return u.Clauses(dbresolver.Write)
}

func (u userInfoDo) Session(config *gorm.Session) IUserInfoDo {
	return u.withDO(u.DO.Session(config))
}

func (u userInfoDo) Clauses(conds ...clause.Expression) IUserInfoDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userInfoDo) Returning(value interface{}, columns ...string) IUserInfoDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u userInfoDo) Not(conds ...gen.Condition) IUserInfoDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userInfoDo) Or(conds ...gen.Condition) IUserInfoDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userInfoDo) Select(conds ...field.Expr) IUserInfoDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userInfoDo) Where(conds ...gen.Condition) IUserInfoDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userInfoDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IUserInfoDo {
	return u.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (u userInfoDo) Order(conds ...field.Expr) IUserInfoDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userInfoDo) Distinct(cols ...field.Expr) IUserInfoDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userInfoDo) Omit(cols ...field.Expr) IUserInfoDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userInfoDo) Join(table schema.Tabler, on ...field.Expr) IUserInfoDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userInfoDo) LeftJoin(table schema.Tabler, on ...field.Expr) IUserInfoDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userInfoDo) RightJoin(table schema.Tabler, on ...field.Expr) IUserInfoDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userInfoDo) Group(cols ...field.Expr) IUserInfoDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userInfoDo) Having(conds ...gen.Condition) IUserInfoDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userInfoDo) Limit(limit int) IUserInfoDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userInfoDo) Offset(offset int) IUserInfoDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userInfoDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IUserInfoDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userInfoDo) Unscoped() IUserInfoDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userInfoDo) Create(values ...*model.UserInfo) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userInfoDo) CreateInBatches(values []*model.UserInfo, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userInfoDo) Save(values ...*model.UserInfo) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userInfoDo) First() (*model.UserInfo, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserInfo), nil
	}
}

func (u userInfoDo) Take() (*model.UserInfo, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserInfo), nil
	}
}

func (u userInfoDo) Last() (*model.UserInfo, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserInfo), nil
	}
}

func (u userInfoDo) Find() ([]*model.UserInfo, error) {
	result, err := u.DO.Find()
	return result.([]*model.UserInfo), err
}

func (u userInfoDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserInfo, err error) {
	buf := make([]*model.UserInfo, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userInfoDo) FindInBatches(result *[]*model.UserInfo, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userInfoDo) Attrs(attrs ...field.AssignExpr) IUserInfoDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userInfoDo) Assign(attrs ...field.AssignExpr) IUserInfoDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userInfoDo) Joins(fields ...field.RelationField) IUserInfoDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u userInfoDo) Preload(fields ...field.RelationField) IUserInfoDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u userInfoDo) FirstOrInit() (*model.UserInfo, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserInfo), nil
	}
}

func (u userInfoDo) FirstOrCreate() (*model.UserInfo, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserInfo), nil
	}
}

func (u userInfoDo) FindByPage(offset int, limit int) (result []*model.UserInfo, count int64, err error) {
	result, err = u.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = u.Offset(-1).Limit(-1).Count()
	return
}

func (u userInfoDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u userInfoDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u userInfoDo) Delete(models ...*model.UserInfo) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *userInfoDo) withDO(do gen.Dao) *userInfoDo {
	u.DO = *do.(*gen.DO)
	return u
}
