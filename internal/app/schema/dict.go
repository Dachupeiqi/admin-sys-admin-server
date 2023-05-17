package schema

import (
	"ginAdminServer/pkg/util/json"
	"time"
)

// Dict 字典对象
type DictData struct {
	ID        uint64    `json:"id,string"`  //id
	Sort      int       `json:"Sort"`       //字典排序
	Label     string    `json:"Label"`      //字典标签
	Value     string    `json:"Value"`      //字典键值
	Type      string    `json:"Type"`       //字典类型
	CssClass  string    `json:"cssClass"`   //样式属性（其他样式扩展）
	ListClass string    `json:"listClass"`  //表格回显属性
	IsDefault string    `json:"isDefault"`  //是否默认（Y是N否）
	Status    int       `json:"status"`     //状态（1:启用 2:禁用）
	Remark    string    `json:"remark"`     //备注
	Creator   uint64    `json:"creator"`    //创建者
	CreatedAt time.Time `json:"created_at"` //创建时间
	UpdatedAt time.Time `json:"updated_at"` //更新时间
}

func (a *DictData) String() string {
	return json.MarshalToString(a)
}

// DictDataQueryParam 查询条件
type DictDataQueryParam struct {
	PaginationParam
	IDs        []uint64 `form:"-"`          // 唯一标识列表
	Label      string   `form:"Label"`      //字典标签
	Value      string   `form:"Value"`      //字典键值
	Type       string   `form:"Type"`       //字典类型
	Status     int      `form:"status"`     //状态(0:启用 1:禁用)
	QueryValue string   `form:"QueryValue"` //模糊查询字段
}

// DictDataQueryOptions 查询可选参数项
type DictDataQueryOptions struct {
	OrderFields  []*OrderField
	SelectFields []string
}

// DictDataQueryResult 查询结果
type DictDataQueryResult struct {
	Data       DictDatas
	PageResult *PaginationResult
}

// Menus 菜单列表
type DictDatas []*DictData

// ----------------------------------------DictType--------------------------------------

// DictType 字典类型对象
type DictType struct {
	ID        uint64    `json:"id,string"`  //id
	Name      string    `json:"name"`       //字典名称
	Type      string    `json:"type"`       //字典类型
	Status    int       `json:"status"`     //状态（1:启用 2:禁用）
	Remark    string    `json:"remark"`     //备注
	Creator   uint64    `json:"creator"`    //创建者
	CreatedAt time.Time `json:"created_at"` //创建时间                         // 创建时间
	UpdatedAt time.Time `json:"updated_at"` //更新时间
}

// DictTypeQueryParam 查询条件
type DictTypeQueryParam struct {
	PaginationParam
	IDs    []uint64 `form:"-"`      // 唯一标识列表
	Name   string   `form:"name"`   //字典键值
	Type   string   `form:"type"`   //字典类型
	Status int      `form:"status"` //状态(0:启用 1:禁用)
}

// DictTypeQueryOptions 查询可选参数项
type DictTypeQueryOptions struct {
	OrderFields []*OrderField // 排序字段
}

// DictTypeQueryResult 查询结果
type DictTypeQueryResult struct {
	Data       DictTypes
	PageResult *PaginationResult
}

// DictTypes 菜单动作管理列表
type DictTypes []*DictType
