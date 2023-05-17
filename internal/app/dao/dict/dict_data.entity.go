package dict

import (
	"context"
	"ginAdminServer/internal/app/dao/util"
	"ginAdminServer/internal/app/schema"
	"ginAdminServer/pkg/util/structure"

	"gorm.io/gorm"
)

func GetDictDataDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDBWithModel(ctx, defDB, new(DictData))
}

type SchemaDictData schema.DictData

func (a SchemaDictData) ToDictData() *DictData {
	item := new(DictData)
	structure.Copy(a, item)
	return item
}

type DictData struct {
	util.Model
	Sort      int    `gorm:"size:20;"`  //字典排序
	Label     string `gorm:"size:128;"` //字典标签
	Value     string `gorm:"size:255;"` //字典键值
	Type      string `gorm:"size:64;"`  //字典类型
	CssClass  string `gorm:"size:128;"` //样式属性（其他样式扩展）
	ListClass string `gorm:"size:128;"` //表格回显属性
	IsDefault string `gorm:"size:8;"`   //是否默认（Y是N否）
	Status    int    `gorm:"size:4;"`   //状态（1:启用 2:禁用）
	Remark    string `gorm:"size:255;"` //备注
}

func (a DictData) ToSchemaDictData() *schema.DictData {
	item := new(schema.DictData)
	structure.Copy(a, item)
	return item
}

type DictDatas []*DictData

func (a DictDatas) ToSchemaDictDatas() []*schema.DictData {
	list := make([]*schema.DictData, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaDictData()
	}
	return list
}
