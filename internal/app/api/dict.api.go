package api

import (
	"ginAdminServer/internal/app/contextx"
	"ginAdminServer/internal/app/ginx"
	"ginAdminServer/internal/app/schema"
	"ginAdminServer/internal/app/service"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var DictSet = wire.NewSet(wire.Struct(new(DictAPI), "*"))

type DictAPI struct {
	DictSrv *service.DictSrv
}

// ----------------------------------------DictData--------------------------------------

func (a *DictAPI) QueryDictData(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.DictDataQueryParam
	if err := ginx.ParseQuery(c, &params); err != nil {
		ginx.ResError(c, err)
		return
	}
	params.Pagination = true
	result, err := a.DictSrv.QueryDictData(ctx, params, schema.DictDataQueryOptions{
		OrderFields: schema.NewOrderFields(
			//schema.NewOrderField("sequence", schema.OrderByDESC),
			schema.NewOrderField("id", schema.OrderByDESC),
		),
	})
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResPage(c, result.Data, result.PageResult)
}

func (a *DictAPI) GetDictData(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.DictSrv.GetDictData(ctx, ginx.ParseParamID(c, "id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, item)
}

func (a *DictAPI) CreateDictData(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.DictData
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}

	item.Creator = contextx.FromUserID(ctx)
	result, err := a.DictSrv.CreateDictData(ctx, item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, result)
}

func (a *DictAPI) UpdateDictData(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.DictData
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}

	err := a.DictSrv.UpdateDictData(ctx, ginx.ParseParamID(c, "id"), item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c)
}

func (a *DictAPI) DeleteDictData(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.DictSrv.DeleteDictData(ctx, ginx.ParseParamID(c, "id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c)
}

// ----------------------------------------DictType--------------------------------------

func (a *DictAPI) QueryDictType(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.DictTypeQueryParam
	if err := ginx.ParseQuery(c, &params); err != nil {
		ginx.ResError(c, err)
		return
	}
	params.Pagination = true
	result, err := a.DictSrv.QueryDictType(ctx, params, schema.DictTypeQueryOptions{
		OrderFields: schema.NewOrderFields(
			//schema.NewOrderField("sequence", schema.OrderByDESC),
			schema.NewOrderField("id", schema.OrderByDESC),
		),
	})
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResPage(c, result.Data, result.PageResult)
}

func (a *DictAPI) GetDictType(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.DictSrv.GetDictType(ctx, ginx.ParseParamID(c, "id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, item)
}

func (a *DictAPI) CreateDictType(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.DictType
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}

	item.Creator = contextx.FromUserID(ctx)
	result, err := a.DictSrv.CreateDictType(ctx, item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, result)
}

func (a *DictAPI) UpdateDictType(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.DictType
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}

	err := a.DictSrv.UpdateDictType(ctx, ginx.ParseParamID(c, "id"), item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c)
}

func (a *DictAPI) DeleteDictType(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.DictSrv.DeleteDictType(ctx, ginx.ParseParamID(c, "id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c)
}
