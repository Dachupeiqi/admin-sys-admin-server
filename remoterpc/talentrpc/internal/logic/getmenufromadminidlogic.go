package logic

import (
	"context"
	"encoding/json"
	__ "ginAdminServer/remoterpc/talentrpc/pb"

	"ginAdminServer/remoterpc/talentrpc/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuFromAdminIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMenuFromAdminIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuFromAdminIdLogic {
	return &GetMenuFromAdminIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMenuFromAdminIdLogic) GetMenuFromAdminId(in *__.GetMenuTreeReq) (*__.GetMenuTreeResp, error) {
	// todo: add your logic here and delete this line
	menuTree, err := l.svcCtx.LoginSrv.QueryUserMenuTree(l.ctx, in.AdminId)
	if err != nil {
		return nil, err
	}
	menuTreeJsonStr, _ := json.Marshal(menuTree)

	return &__.GetMenuTreeResp{
		List: string(menuTreeJsonStr),
	}, nil
}
