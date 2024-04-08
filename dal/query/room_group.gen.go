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

	"DiTing-Go/dal/model"
)

func newRoomGroup(db *gorm.DB, opts ...gen.DOOption) roomGroup {
	_roomGroup := roomGroup{}

	_roomGroup.roomGroupDo.UseDB(db, opts...)
	_roomGroup.roomGroupDo.UseModel(&model.RoomGroup{})

	tableName := _roomGroup.roomGroupDo.TableName()
	_roomGroup.ALL = field.NewAsterisk(tableName)
	_roomGroup.ID = field.NewInt64(tableName, "id")
	_roomGroup.RoomID = field.NewInt64(tableName, "room_id")
	_roomGroup.Name = field.NewString(tableName, "name")
	_roomGroup.Avatar = field.NewString(tableName, "avatar")
	_roomGroup.ExtJSON = field.NewString(tableName, "ext_json")
	_roomGroup.DeleteStatus = field.NewInt32(tableName, "delete_status")
	_roomGroup.CreateTime = field.NewTime(tableName, "create_time")
	_roomGroup.UpdateTime = field.NewTime(tableName, "update_time")

	_roomGroup.fillFieldMap()

	return _roomGroup
}

// roomGroup 群聊房间表
type roomGroup struct {
	roomGroupDo roomGroupDo

	ALL          field.Asterisk
	ID           field.Int64  // id
	RoomID       field.Int64  // 房间id
	Name         field.String // 群名称
	Avatar       field.String // 群头像
	ExtJSON      field.String // 额外信息（根据不同类型房间有不同存储的东西）
	DeleteStatus field.Int32  // 逻辑删除(0-正常,1-删除)
	CreateTime   field.Time   // 创建时间
	UpdateTime   field.Time   // 修改时间

	fieldMap map[string]field.Expr
}

func (r roomGroup) Table(newTableName string) *roomGroup {
	r.roomGroupDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r roomGroup) As(alias string) *roomGroup {
	r.roomGroupDo.DO = *(r.roomGroupDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *roomGroup) updateTableName(table string) *roomGroup {
	r.ALL = field.NewAsterisk(table)
	r.ID = field.NewInt64(table, "id")
	r.RoomID = field.NewInt64(table, "room_id")
	r.Name = field.NewString(table, "name")
	r.Avatar = field.NewString(table, "avatar")
	r.ExtJSON = field.NewString(table, "ext_json")
	r.DeleteStatus = field.NewInt32(table, "delete_status")
	r.CreateTime = field.NewTime(table, "create_time")
	r.UpdateTime = field.NewTime(table, "update_time")

	r.fillFieldMap()

	return r
}

func (r *roomGroup) WithContext(ctx context.Context) IRoomGroupDo {
	return r.roomGroupDo.WithContext(ctx)
}

func (r roomGroup) TableName() string { return r.roomGroupDo.TableName() }

func (r roomGroup) Alias() string { return r.roomGroupDo.Alias() }

func (r roomGroup) Columns(cols ...field.Expr) gen.Columns { return r.roomGroupDo.Columns(cols...) }

func (r *roomGroup) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *roomGroup) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 8)
	r.fieldMap["id"] = r.ID
	r.fieldMap["room_id"] = r.RoomID
	r.fieldMap["name"] = r.Name
	r.fieldMap["avatar"] = r.Avatar
	r.fieldMap["ext_json"] = r.ExtJSON
	r.fieldMap["delete_status"] = r.DeleteStatus
	r.fieldMap["create_time"] = r.CreateTime
	r.fieldMap["update_time"] = r.UpdateTime
}

func (r roomGroup) clone(db *gorm.DB) roomGroup {
	r.roomGroupDo.ReplaceConnPool(db.Statement.ConnPool)
	return r
}

func (r roomGroup) replaceDB(db *gorm.DB) roomGroup {
	r.roomGroupDo.ReplaceDB(db)
	return r
}

type roomGroupDo struct{ gen.DO }

type IRoomGroupDo interface {
	gen.SubQuery
	Debug() IRoomGroupDo
	WithContext(ctx context.Context) IRoomGroupDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IRoomGroupDo
	WriteDB() IRoomGroupDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IRoomGroupDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IRoomGroupDo
	Not(conds ...gen.Condition) IRoomGroupDo
	Or(conds ...gen.Condition) IRoomGroupDo
	Select(conds ...field.Expr) IRoomGroupDo
	Where(conds ...gen.Condition) IRoomGroupDo
	Order(conds ...field.Expr) IRoomGroupDo
	Distinct(cols ...field.Expr) IRoomGroupDo
	Omit(cols ...field.Expr) IRoomGroupDo
	Join(table schema.Tabler, on ...field.Expr) IRoomGroupDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IRoomGroupDo
	RightJoin(table schema.Tabler, on ...field.Expr) IRoomGroupDo
	Group(cols ...field.Expr) IRoomGroupDo
	Having(conds ...gen.Condition) IRoomGroupDo
	Limit(limit int) IRoomGroupDo
	Offset(offset int) IRoomGroupDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IRoomGroupDo
	Unscoped() IRoomGroupDo
	Create(values ...*model.RoomGroup) error
	CreateInBatches(values []*model.RoomGroup, batchSize int) error
	Save(values ...*model.RoomGroup) error
	First() (*model.RoomGroup, error)
	Take() (*model.RoomGroup, error)
	Last() (*model.RoomGroup, error)
	Find() ([]*model.RoomGroup, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.RoomGroup, err error)
	FindInBatches(result *[]*model.RoomGroup, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.RoomGroup) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IRoomGroupDo
	Assign(attrs ...field.AssignExpr) IRoomGroupDo
	Joins(fields ...field.RelationField) IRoomGroupDo
	Preload(fields ...field.RelationField) IRoomGroupDo
	FirstOrInit() (*model.RoomGroup, error)
	FirstOrCreate() (*model.RoomGroup, error)
	FindByPage(offset int, limit int) (result []*model.RoomGroup, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IRoomGroupDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (r roomGroupDo) Debug() IRoomGroupDo {
	return r.withDO(r.DO.Debug())
}

func (r roomGroupDo) WithContext(ctx context.Context) IRoomGroupDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r roomGroupDo) ReadDB() IRoomGroupDo {
	return r.Clauses(dbresolver.Read)
}

func (r roomGroupDo) WriteDB() IRoomGroupDo {
	return r.Clauses(dbresolver.Write)
}

func (r roomGroupDo) Session(config *gorm.Session) IRoomGroupDo {
	return r.withDO(r.DO.Session(config))
}

func (r roomGroupDo) Clauses(conds ...clause.Expression) IRoomGroupDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r roomGroupDo) Returning(value interface{}, columns ...string) IRoomGroupDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r roomGroupDo) Not(conds ...gen.Condition) IRoomGroupDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r roomGroupDo) Or(conds ...gen.Condition) IRoomGroupDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r roomGroupDo) Select(conds ...field.Expr) IRoomGroupDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r roomGroupDo) Where(conds ...gen.Condition) IRoomGroupDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r roomGroupDo) Order(conds ...field.Expr) IRoomGroupDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r roomGroupDo) Distinct(cols ...field.Expr) IRoomGroupDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r roomGroupDo) Omit(cols ...field.Expr) IRoomGroupDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r roomGroupDo) Join(table schema.Tabler, on ...field.Expr) IRoomGroupDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r roomGroupDo) LeftJoin(table schema.Tabler, on ...field.Expr) IRoomGroupDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r roomGroupDo) RightJoin(table schema.Tabler, on ...field.Expr) IRoomGroupDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r roomGroupDo) Group(cols ...field.Expr) IRoomGroupDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r roomGroupDo) Having(conds ...gen.Condition) IRoomGroupDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r roomGroupDo) Limit(limit int) IRoomGroupDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r roomGroupDo) Offset(offset int) IRoomGroupDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r roomGroupDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IRoomGroupDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r roomGroupDo) Unscoped() IRoomGroupDo {
	return r.withDO(r.DO.Unscoped())
}

func (r roomGroupDo) Create(values ...*model.RoomGroup) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r roomGroupDo) CreateInBatches(values []*model.RoomGroup, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r roomGroupDo) Save(values ...*model.RoomGroup) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r roomGroupDo) First() (*model.RoomGroup, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.RoomGroup), nil
	}
}

func (r roomGroupDo) Take() (*model.RoomGroup, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.RoomGroup), nil
	}
}

func (r roomGroupDo) Last() (*model.RoomGroup, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.RoomGroup), nil
	}
}

func (r roomGroupDo) Find() ([]*model.RoomGroup, error) {
	result, err := r.DO.Find()
	return result.([]*model.RoomGroup), err
}

func (r roomGroupDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.RoomGroup, err error) {
	buf := make([]*model.RoomGroup, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r roomGroupDo) FindInBatches(result *[]*model.RoomGroup, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r roomGroupDo) Attrs(attrs ...field.AssignExpr) IRoomGroupDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r roomGroupDo) Assign(attrs ...field.AssignExpr) IRoomGroupDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r roomGroupDo) Joins(fields ...field.RelationField) IRoomGroupDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r roomGroupDo) Preload(fields ...field.RelationField) IRoomGroupDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r roomGroupDo) FirstOrInit() (*model.RoomGroup, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.RoomGroup), nil
	}
}

func (r roomGroupDo) FirstOrCreate() (*model.RoomGroup, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.RoomGroup), nil
	}
}

func (r roomGroupDo) FindByPage(offset int, limit int) (result []*model.RoomGroup, count int64, err error) {
	result, err = r.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = r.Offset(-1).Limit(-1).Count()
	return
}

func (r roomGroupDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r roomGroupDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r roomGroupDo) Delete(models ...*model.RoomGroup) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *roomGroupDo) withDO(do gen.Dao) *roomGroupDo {
	r.DO = *do.(*gen.DO)
	return r
}