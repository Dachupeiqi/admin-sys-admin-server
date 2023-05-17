package dict

import (
	"context"
	"ginAdminServer/internal/app/dao/util"
	"ginAdminServer/internal/app/schema"
	"ginAdminServer/pkg/util/structure"

	"gorm.io/gorm"
)

func GetDictTypeDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDBWithModel(ctx, defDB, new(DictType))
}

type SchemaDictType schema.DictType

func (a SchemaDictType) ToDictType() *DictType {
	item := new(DictType)
	structure.Copy(a, item)
	return item
}

type DictType struct {
	util.Model
	Name   string `gorm:"size:128;"` //字典名称
	Type   string `gorm:"size:128;"` //字典类型
	Status int    `gorm:"size:4;"`   //状态（1:启用 2:禁用）
	Remark string `gorm:"size:255;"` //备注
}

func (a DictType) ToSchemaDictType() *schema.DictType {
	item := new(schema.DictType)
	structure.Copy(a, item)
	return item
}

type DictTypes []*DictType

func (a DictTypes) ToSchemaDictTypes() []*schema.DictType {
	list := make([]*schema.DictType, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaDictType()
	}
	return list
}
