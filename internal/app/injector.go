package app

import (
	"ginAdminServer/internal/app/service"
	"ginAdminServer/pkg/auth"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var InjectorSet = wire.NewSet(wire.Struct(new(Injector), "*"))

type Injector struct {
	Engine         *gin.Engine
	Auth           auth.Auther
	CasbinEnforcer *casbin.SyncedEnforcer
	MenuSrv        *service.MenuSrv
	LoginSrv       *service.LoginSrv
	DictSrv        *service.DictSrv
	UserSrv        *service.UserSrv
}
