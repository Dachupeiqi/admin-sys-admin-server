package remoterpc

import (
	"ginAdminServer/internal/app/service"
)

type InjectorForTalent struct {
	MenuSrv  *service.MenuSrv
	LoginSrv *service.LoginSrv
	DictSrv  *service.DictSrv
	UserSrv  *service.UserSrv
}
