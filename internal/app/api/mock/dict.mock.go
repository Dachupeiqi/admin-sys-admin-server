package mock

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var DictSet = wire.NewSet(wire.Struct(new(DictMock), "*"))

type DictMock struct {
}

// ----------------------------------------DictData--------------------------------------

// QueryDictData
// @Tags DictAPI
// @Summary 字典数据列表查询
// @Description 获取JSON
// @Param current query int true "分页索引" default(1)
// @Param pageSize query int true "分页大小" default(10)
// @Param queryValue query string false "查询值"
// @Param status query string false "状态(0:启用 1:禁用)"
// @Param type query string false "字典类型"
// @Success 200 {object} schema.ListResult{list=[]schema.DictData} "查询结果"
// @Failure 400 {object} schema.ErrorResult "{error:{code:0,message:bad request}}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:9999,message:invalid signature}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:internal server error}}"
// @Router /api/v1/dict/data [get]
func (a *DictMock) QueryDictData(c *gin.Context) {
}

// GetDictData
// @Tags DictAPI
// @Summary 通过id获取字典数据
// @Param id path int true "字典数据id"
// @Success 200 {object} schema.DictData
// @Failure 401 {object} schema.ErrorResult "{error:{code:9999,message:invalid signature}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:internal server error}}"
// @Router /api/v1/dict/data/{id} [get]
func (a *DictMock) GetDictData(c *gin.Context) {
}

// CreateDictData
// @Tags DictAPI
// @Summary 创建字典数据
// @Security ApiKeyAuth
// @Param body body schema.DictData true "创建数据"
// @Success 200 {object} schema.IDResult
// @Failure 400 {object} schema.ErrorResult "{error:{code:0,message:bad request}}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:9999,message:invalid signature}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:internal server error}}"
// @Router /api/v1/dict/data [post]
func (a *DictMock) CreateDictData(c *gin.Context) {
}

// UpdateDictData
// @Tags DictAPI
// @Summary 更新数据
// @Security ApiKeyAuth
// @Param id path int true "唯一标识"
// @Param body body schema.DictData true "更新数据"
// @Success 200 {object} schema.StatusResult "{status:OK}"
// @Failure 400 {object} schema.ErrorResult "{error:{code:0,message:bad request}}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:9999,message:invalid signature}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:internal server error}}"
// @Router /api/v1/dict/data/{id} [put]
func (a *DictMock) UpdateDictData(c *gin.Context) {
}

// DeleteDictData
// @Tags DictAPI
// @Summary 删除数据
// @Security ApiKeyAuth
// @Param id path int true "唯一标识"
// @Success 200 {object} schema.StatusResult "{status:OK}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:9999,message:invalid signature}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:internal server error}}"
// @Router /api/v1/dict/data/{id} [delete]
func (a *MenuMock) DeleteDictData(c *gin.Context) {
}

// ----------------------------------------DictType--------------------------------------

// QueryDictType
// @Tags DictAPI
// @Summary 字典数据列表查询
// @Description 获取JSON
// @Param current query int true "分页索引" default(1)
// @Param pageSize query int true "分页大小" default(10)
// @Param queryValue query string false "查询值"
// @Param Name query string false "字典名称"
// @Param Id query string false "字典类型Id"
// @Param Type query string false "字典类型"
// @Success 200 {object} schema.ListResult{list=[]schema.DictType} "查询结果"
// @Failure 400 {object} schema.ErrorResult "{error:{code:0,message:bad request}}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:9999,message:invalid signature}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:internal server error}}"
// @Router /api/v1/dict/type [get]
func (a *DictMock) QueryDictType(c *gin.Context) {
}

// GetDictType
// @Tags DictAPI
// @Summary 通过id获取字典类型
// @Param id path int true "字典数据id"
// @Success 200 {object} schema.DictType
// @Failure 401 {object} schema.ErrorResult "{error:{code:9999,message:invalid signature}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:internal server error}}"
// @Router /api/v1/dict/type/{id} [get]
func (a *DictMock) GetDictType(c *gin.Context) {
}

// CreateDictType
// @Tags DictAPI
// @Summary 创建字典数据
// @Security ApiKeyAuth
// @Param body body schema.DictType true "创建数据"
// @Success 200 {object} schema.IDResult
// @Failure 400 {object} schema.ErrorResult "{error:{code:0,message:bad request}}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:9999,message:invalid signature}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:internal server error}}"
// @Router /api/v1/dict/type [post]
func (a *DictMock) CreateDictType(c *gin.Context) {
}

// UpdateDictType
// @Tags DictAPI
// @Summary 更新数据
// @Security ApiKeyAuth
// @Param id path int true "唯一标识"
// @Param body body schema.DictType true "更新数据"
// @Success 200 {object} schema.StatusResult "{status:OK}"
// @Failure 400 {object} schema.ErrorResult "{error:{code:0,message:bad request}}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:9999,message:invalid signature}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:internal server error}}"
// @Router /api/v1/dict/type/{id} [put]
func (a *DictMock) UpdateDictType(c *gin.Context) {
}

// DeleteDictType
// @Tags DictAPI
// @Summary 删除数据
// @Security ApiKeyAuth
// @Param id path int true "唯一标识"
// @Success 200 {object} schema.StatusResult "{status:OK}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:9999,message:invalid signature}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:internal server error}}"
// @Router /api/v1/dict/type/{id} [delete]
func (a *MenuMock) DeleteDictType(c *gin.Context) {
}
