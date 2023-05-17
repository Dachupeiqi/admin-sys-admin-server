package service

import (
	"context"
	"ginAdminServer/internal/app/dao"
	"ginAdminServer/internal/app/schema"
	"ginAdminServer/pkg/errors"
	"ginAdminServer/pkg/util/snowflake"

	"github.com/google/wire"
)

var DictSet = wire.NewSet(wire.Struct(new(DictSrv), "*"))

type DictSrv struct {
	TransRepo    *dao.TransRepo
	DictDataRepo *dao.DictDataRepo
	DictTypeRepo *dao.DictTypeRepo
}

// ----------------------------------------DictData--------------------------------------

func (a *DictSrv) QueryDictData(ctx context.Context, params schema.DictDataQueryParam, opts ...schema.DictDataQueryOptions) (*schema.DictDataQueryResult, error) {
	return a.DictDataRepo.Query(ctx, params, opts...)
}

func (a *DictSrv) GetDictData(ctx context.Context, id uint64, opts ...schema.DictDataQueryOptions) (*schema.DictData, error) {
	item, err := a.DictDataRepo.Get(ctx, id, opts...)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	return item, nil
}

func (a *DictSrv) CreateDictData(ctx context.Context, item schema.DictData) (*schema.IDResult, error) {

	item.ID = snowflake.MustID()
	err := a.TransRepo.Exec(ctx, func(ctx context.Context) error {
		return a.DictDataRepo.Create(ctx, item)
	})
	if err != nil {
		return nil, err
	}

	return schema.NewIDResult(item.ID), nil
}

func (a *DictSrv) UpdateDictData(ctx context.Context, id uint64, item schema.DictData) error {
	oldItem, err := a.GetDictData(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	item.ID = oldItem.ID
	item.Creator = oldItem.Creator
	item.CreatedAt = oldItem.CreatedAt

	return a.TransRepo.Exec(ctx, func(ctx context.Context) error {
		return a.DictDataRepo.Update(ctx, id, item)
	})
}

func (a *DictSrv) DeleteDictData(ctx context.Context, id uint64) error {
	oldItem, err := a.DictDataRepo.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	return a.TransRepo.Exec(ctx, func(ctx context.Context) error {
		return a.DictDataRepo.Delete(ctx, id)
	})
}

// ----------------------------------------DictType--------------------------------------

func (a *DictSrv) QueryDictType(ctx context.Context, params schema.DictTypeQueryParam, opts ...schema.DictTypeQueryOptions) (*schema.DictTypeQueryResult, error) {
	return a.DictTypeRepo.Query(ctx, params, opts...)
}

func (a *DictSrv) GetDictType(ctx context.Context, id uint64, opts ...schema.DictTypeQueryOptions) (*schema.DictType, error) {
	item, err := a.DictTypeRepo.Get(ctx, id, opts...)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	return item, nil
}

func (a *DictSrv) checkDictType(ctx context.Context, item schema.DictType) error {
	result, err := a.DictTypeRepo.Query(ctx, schema.DictTypeQueryParam{
		PaginationParam: schema.PaginationParam{
			OnlyCount: true,
		},
		Type: item.Type,
	})
	if err != nil {
		return err
	} else if result.PageResult.Total > 0 {
		return errors.New400Response("当前字典类型已经存在！")
	}
	return nil
}

func (a *DictSrv) CreateDictType(ctx context.Context, item schema.DictType) (*schema.IDResult, error) {
	if err := a.checkDictType(ctx, item); err != nil {
		return nil, err
	}

	item.ID = snowflake.MustID()
	err := a.TransRepo.Exec(ctx, func(ctx context.Context) error {
		return a.DictTypeRepo.Create(ctx, item)
	})
	if err != nil {
		return nil, err
	}

	return schema.NewIDResult(item.ID), nil
}

func (a *DictSrv) UpdateDictType(ctx context.Context, id uint64, item schema.DictType) error {
	oldItem, err := a.GetDictType(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	} else if oldItem.Type != item.Type {
		if err := a.checkDictType(ctx, item); err != nil {
			return err
		}
	}

	item.ID = oldItem.ID
	item.Creator = oldItem.Creator
	item.CreatedAt = oldItem.CreatedAt

	return a.TransRepo.Exec(ctx, func(ctx context.Context) error {
		return a.DictTypeRepo.Update(ctx, id, item)
	})
}

func (a *DictSrv) DeleteDictType(ctx context.Context, id uint64) error {
	oldItem, err := a.DictTypeRepo.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}
	result, err := a.DictDataRepo.Query(ctx, schema.DictDataQueryParam{
		Type: oldItem.Type,
	})
	if err != nil {
		return err
	} else if len(result.Data) != 0 {
		return errors.New400Response("存在该字典类型的数据，不能删除！")
	}

	return a.TransRepo.Exec(ctx, func(ctx context.Context) error {
		return a.DictTypeRepo.Delete(ctx, id)
	})
}
