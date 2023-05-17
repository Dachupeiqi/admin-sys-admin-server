package dict

import (
	"context"
	"ginAdminServer/internal/app/dao/util"
	"ginAdminServer/internal/app/schema"
	"ginAdminServer/pkg/errors"

	"github.com/google/wire"
	"gorm.io/gorm"
)

var DictDataSet = wire.NewSet(wire.Struct(new(DictDataRepo), "*"))

type DictDataRepo struct {
	DB *gorm.DB
}

func (a *DictDataRepo) getQueryOption(opts ...schema.DictDataQueryOptions) schema.DictDataQueryOptions {
	var opt schema.DictDataQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

func (a *DictDataRepo) Query(ctx context.Context, params schema.DictDataQueryParam, opts ...schema.DictDataQueryOptions) (*schema.DictDataQueryResult, error) {
	opt := a.getQueryOption(opts...)

	db := GetDictDataDB(ctx, a.DB)
	if v := params.IDs; len(v) > 0 {
		db = db.Where("id IN (?)", v)
	}
	if v := params.Label; v != "" {
		db = db.Where("label=?", v)
	}
	if v := params.Value; v != "" {
		db = db.Where("value=?", v)
	}
	if v := params.Type; v != "" {
		db = db.Where("type=?", v)
	}
	if v := params.Status; v != 0 {
		db = db.Where("status=?", v)
	}
	if len(opt.SelectFields) > 0 {
		db = db.Select(opt.SelectFields)
	}

	if len(opt.OrderFields) > 0 {
		db = db.Order(util.ParseOrder(opt.OrderFields))
	}
	if v := params.QueryValue; v != "" {
		v = "%" + v + "%"
		db = db.Where("Label LIKE ?", v)
	}

	var list DictDatas
	pr, err := util.WrapPageQuery(ctx, db, params.PaginationParam, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	qr := &schema.DictDataQueryResult{
		PageResult: pr,
		Data:       list.ToSchemaDictDatas(),
	}

	return qr, nil
}

func (a *DictDataRepo) Get(ctx context.Context, id uint64, opts ...schema.DictDataQueryOptions) (*schema.DictData, error) {
	var item DictData
	ok, err := util.FindOne(ctx, GetDictDataDB(ctx, a.DB).Where("id=?", id), &item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}

	return item.ToSchemaDictData(), nil
}

func (a *DictDataRepo) Create(ctx context.Context, item schema.DictData) error {
	eitem := SchemaDictData(item).ToDictData()
	result := GetDictDataDB(ctx, a.DB).Create(eitem)
	return errors.WithStack(result.Error)
}

func (a *DictDataRepo) Update(ctx context.Context, id uint64, item schema.DictData) error {
	eitem := SchemaDictData(item).ToDictData()
	result := GetDictDataDB(ctx, a.DB).Where("id=?", id).Updates(eitem)
	return errors.WithStack(result.Error)
}

func (a *DictDataRepo) Delete(ctx context.Context, id uint64) error {
	result := GetDictDataDB(ctx, a.DB).Where("id=?", id).Delete(DictData{})
	return errors.WithStack(result.Error)
}

func (a *DictDataRepo) UpdateStatus(ctx context.Context, id uint64, status int) error {
	result := GetDictDataDB(ctx, a.DB).Where("id=?", id).Update("status", status)
	return errors.WithStack(result.Error)
}
