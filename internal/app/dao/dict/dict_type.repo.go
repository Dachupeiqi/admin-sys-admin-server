package dict

import (
	"context"
	"ginAdminServer/internal/app/dao/util"
	"ginAdminServer/internal/app/schema"
	"ginAdminServer/pkg/errors"

	"github.com/google/wire"
	"gorm.io/gorm"
)

var DictTypeSet = wire.NewSet(wire.Struct(new(DictTypeRepo), "*"))

type DictTypeRepo struct {
	DB *gorm.DB
}

func (a *DictTypeRepo) getQueryOption(opts ...schema.DictTypeQueryOptions) schema.DictTypeQueryOptions {
	var opt schema.DictTypeQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

func (a *DictTypeRepo) Query(ctx context.Context, params schema.DictTypeQueryParam, opts ...schema.DictTypeQueryOptions) (*schema.DictTypeQueryResult, error) {
	opt := a.getQueryOption(opts...)

	db := GetDictTypeDB(ctx, a.DB)
	if v := params.IDs; len(v) > 0 {
		db = db.Where("id IN (?)", v)
	}
	if v := params.Name; v != "" {
		db = db.Where("name=?", v)
	}
	if v := params.Type; v != "" {
		db = db.Where("type=?", v)
	}
	if v := params.Status; v != 0 {
		db = db.Where("status=?", v)
	}

	if len(opt.OrderFields) > 0 {
		db = db.Order(util.ParseOrder(opt.OrderFields))
	}

	var list DictTypes
	pr, err := util.WrapPageQuery(ctx, db, params.PaginationParam, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	qr := &schema.DictTypeQueryResult{
		PageResult: pr,
		Data:       list.ToSchemaDictTypes(),
	}

	return qr, nil
}

func (a *DictTypeRepo) Get(ctx context.Context, id uint64, opts ...schema.DictTypeQueryOptions) (*schema.DictType, error) {
	db := GetDictTypeDB(ctx, a.DB).Where("id=?", id)
	var item DictType
	ok, err := util.FindOne(ctx, db, &item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}

	return item.ToSchemaDictType(), nil
}

func (a *DictTypeRepo) Create(ctx context.Context, item schema.DictType) error {
	eitem := SchemaDictType(item).ToDictType()
	result := GetDictTypeDB(ctx, a.DB).Create(eitem)
	return errors.WithStack(result.Error)
}

func (a *DictTypeRepo) Update(ctx context.Context, id uint64, item schema.DictType) error {
	eitem := SchemaDictType(item).ToDictType()
	result := GetDictTypeDB(ctx, a.DB).Where("id=?", id).Updates(eitem)
	return errors.WithStack(result.Error)
}

func (a *DictTypeRepo) Delete(ctx context.Context, id uint64) error {
	result := GetDictTypeDB(ctx, a.DB).Where("id=?", id).Delete(DictType{})
	return errors.WithStack(result.Error)
}
