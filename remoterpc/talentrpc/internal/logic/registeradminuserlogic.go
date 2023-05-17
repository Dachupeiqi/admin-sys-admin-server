package logic

import (
	"context"
	"ginAdminServer/internal/app/schema"
	"ginAdminServer/remoterpc/talentrpc/internal/svc"
	pb "ginAdminServer/remoterpc/talentrpc/pb"
	"github.com/LyricTian/gin-admin/v8/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterAdminUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterAdminUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterAdminUserLogic {
	return &RegisterAdminUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterAdminUserLogic) RegisterAdminUser(in *pb.RegisterReq) (*pb.RegisterResp, error) {
	// todo: add your logic here and delete this line
	userName := in.UserName
	userType := in.Type

	talentUserRole := &schema.UserRole{
		RoleID: 88902574807187461,
	}
	EnterpriseUserRole := &schema.UserRole{
		RoleID: 88902734861828101,
	}
	userRoles := make([]*schema.UserRole, 1)

	if userType == "talent" {
		copy(userRoles, []*schema.UserRole{talentUserRole})
	} else if userType == "enterprise" {
		copy(userRoles, []*schema.UserRole{EnterpriseUserRole})
	} else {
		return nil, errors.New("用户类型传入不正确!")
	}

	idResult, err := l.svcCtx.UserSrv.CreateTalentUser(l.ctx, schema.User{
		UserName:  userName,
		RealName:  userName,
		Password:  userName,
		UserRoles: userRoles,
	})
	if err != nil {
		return nil, err
	}
	return &pb.RegisterResp{
		Status:  "OK",
		AdminId: idResult.ID,
	}, nil
}
