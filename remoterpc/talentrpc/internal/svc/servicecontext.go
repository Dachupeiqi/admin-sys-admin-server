package svc

import (
	"ginAdminServer/internal/app/config"
	"ginAdminServer/internal/app/service"
	"ginAdminServer/remoterpc"
	RpcConfig "ginAdminServer/remoterpc/talentrpc/internal/config"
)

const (
	configFile = "etc/config.toml"
)

type ServiceContext struct {
	//zeroRpc配置类
	RpcConfig RpcConfig.Config
	LoginSrv  *service.LoginSrv
	DictSrv   *service.DictSrv
	UserSrv   *service.UserSrv
}

func NewServiceContext(c RpcConfig.Config) *ServiceContext {

	//加载wire项目注入配置类
	config.MustLoad(configFile)
	//构建依赖，拿出Srv
	injector, _, _ := remoterpc.BuildInjectorForTalent()
	loginSrv := injector.LoginSrv
	userSrv := injector.UserSrv
	dictSrv := injector.DictSrv

	return &ServiceContext{
		RpcConfig: c,
		LoginSrv:  loginSrv,
		DictSrv:   dictSrv,
		UserSrv:   userSrv,
	}
}
