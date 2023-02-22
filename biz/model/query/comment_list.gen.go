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

	"douyin/biz/model/orm_gen"
)

func newCommentList(db *gorm.DB, opts ...gen.DOOption) commentList {
	_commentList := commentList{}

	_commentList.commentListDo.UseDB(db, opts...)
	_commentList.commentListDo.UseModel(&orm_gen.CommentList{})

	tableName := _commentList.commentListDo.TableName()
	_commentList.ALL = field.NewAsterisk(tableName)
	_commentList.ID = field.NewInt64(tableName, "id")
	_commentList.VideoID = field.NewInt64(tableName, "video_id")

	_commentList.fillFieldMap()

	return _commentList
}

type commentList struct {
	commentListDo

	ALL     field.Asterisk
	ID      field.Int64 // 评论id
	VideoID field.Int64 // 视频id

	fieldMap map[string]field.Expr
}

func (c commentList) Table(newTableName string) *commentList {
	c.commentListDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c commentList) As(alias string) *commentList {
	c.commentListDo.DO = *(c.commentListDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *commentList) updateTableName(table string) *commentList {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewInt64(table, "id")
	c.VideoID = field.NewInt64(table, "video_id")

	c.fillFieldMap()

	return c
}

func (c *commentList) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *commentList) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 2)
	c.fieldMap["id"] = c.ID
	c.fieldMap["video_id"] = c.VideoID
}

func (c commentList) clone(db *gorm.DB) commentList {
	c.commentListDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c commentList) replaceDB(db *gorm.DB) commentList {
	c.commentListDo.ReplaceDB(db)
	return c
}

type commentListDo struct{ gen.DO }

type ICommentListDo interface {
	gen.SubQuery
	Debug() ICommentListDo
	WithContext(ctx context.Context) ICommentListDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ICommentListDo
	WriteDB() ICommentListDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ICommentListDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ICommentListDo
	Not(conds ...gen.Condition) ICommentListDo
	Or(conds ...gen.Condition) ICommentListDo
	Select(conds ...field.Expr) ICommentListDo
	Where(conds ...gen.Condition) ICommentListDo
	Order(conds ...field.Expr) ICommentListDo
	Distinct(cols ...field.Expr) ICommentListDo
	Omit(cols ...field.Expr) ICommentListDo
	Join(table schema.Tabler, on ...field.Expr) ICommentListDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ICommentListDo
	RightJoin(table schema.Tabler, on ...field.Expr) ICommentListDo
	Group(cols ...field.Expr) ICommentListDo
	Having(conds ...gen.Condition) ICommentListDo
	Limit(limit int) ICommentListDo
	Offset(offset int) ICommentListDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ICommentListDo
	Unscoped() ICommentListDo
	Create(values ...*orm_gen.CommentList) error
	CreateInBatches(values []*orm_gen.CommentList, batchSize int) error
	Save(values ...*orm_gen.CommentList) error
	First() (*orm_gen.CommentList, error)
	Take() (*orm_gen.CommentList, error)
	Last() (*orm_gen.CommentList, error)
	Find() ([]*orm_gen.CommentList, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*orm_gen.CommentList, err error)
	FindInBatches(result *[]*orm_gen.CommentList, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*orm_gen.CommentList) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ICommentListDo
	Assign(attrs ...field.AssignExpr) ICommentListDo
	Joins(fields ...field.RelationField) ICommentListDo
	Preload(fields ...field.RelationField) ICommentListDo
	FirstOrInit() (*orm_gen.CommentList, error)
	FirstOrCreate() (*orm_gen.CommentList, error)
	FindByPage(offset int, limit int) (result []*orm_gen.CommentList, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ICommentListDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c commentListDo) Debug() ICommentListDo {
	return c.withDO(c.DO.Debug())
}

func (c commentListDo) WithContext(ctx context.Context) ICommentListDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c commentListDo) ReadDB() ICommentListDo {
	return c.Clauses(dbresolver.Read)
}

func (c commentListDo) WriteDB() ICommentListDo {
	return c.Clauses(dbresolver.Write)
}

func (c commentListDo) Session(config *gorm.Session) ICommentListDo {
	return c.withDO(c.DO.Session(config))
}

func (c commentListDo) Clauses(conds ...clause.Expression) ICommentListDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c commentListDo) Returning(value interface{}, columns ...string) ICommentListDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c commentListDo) Not(conds ...gen.Condition) ICommentListDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c commentListDo) Or(conds ...gen.Condition) ICommentListDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c commentListDo) Select(conds ...field.Expr) ICommentListDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c commentListDo) Where(conds ...gen.Condition) ICommentListDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c commentListDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ICommentListDo {
	return c.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (c commentListDo) Order(conds ...field.Expr) ICommentListDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c commentListDo) Distinct(cols ...field.Expr) ICommentListDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c commentListDo) Omit(cols ...field.Expr) ICommentListDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c commentListDo) Join(table schema.Tabler, on ...field.Expr) ICommentListDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c commentListDo) LeftJoin(table schema.Tabler, on ...field.Expr) ICommentListDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c commentListDo) RightJoin(table schema.Tabler, on ...field.Expr) ICommentListDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c commentListDo) Group(cols ...field.Expr) ICommentListDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c commentListDo) Having(conds ...gen.Condition) ICommentListDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c commentListDo) Limit(limit int) ICommentListDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c commentListDo) Offset(offset int) ICommentListDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c commentListDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ICommentListDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c commentListDo) Unscoped() ICommentListDo {
	return c.withDO(c.DO.Unscoped())
}

func (c commentListDo) Create(values ...*orm_gen.CommentList) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c commentListDo) CreateInBatches(values []*orm_gen.CommentList, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c commentListDo) Save(values ...*orm_gen.CommentList) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c commentListDo) First() (*orm_gen.CommentList, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*orm_gen.CommentList), nil
	}
}

func (c commentListDo) Take() (*orm_gen.CommentList, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*orm_gen.CommentList), nil
	}
}

func (c commentListDo) Last() (*orm_gen.CommentList, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*orm_gen.CommentList), nil
	}
}

func (c commentListDo) Find() ([]*orm_gen.CommentList, error) {
	result, err := c.DO.Find()
	return result.([]*orm_gen.CommentList), err
}

func (c commentListDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*orm_gen.CommentList, err error) {
	buf := make([]*orm_gen.CommentList, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c commentListDo) FindInBatches(result *[]*orm_gen.CommentList, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c commentListDo) Attrs(attrs ...field.AssignExpr) ICommentListDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c commentListDo) Assign(attrs ...field.AssignExpr) ICommentListDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c commentListDo) Joins(fields ...field.RelationField) ICommentListDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c commentListDo) Preload(fields ...field.RelationField) ICommentListDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c commentListDo) FirstOrInit() (*orm_gen.CommentList, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*orm_gen.CommentList), nil
	}
}

func (c commentListDo) FirstOrCreate() (*orm_gen.CommentList, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*orm_gen.CommentList), nil
	}
}

func (c commentListDo) FindByPage(offset int, limit int) (result []*orm_gen.CommentList, count int64, err error) {
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

func (c commentListDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c commentListDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c commentListDo) Delete(models ...*orm_gen.CommentList) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *commentListDo) withDO(do gen.Dao) *commentListDo {
	c.DO = *do.(*gen.DO)
	return c
}
