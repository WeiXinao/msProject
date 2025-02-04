package user

import (
	"context"
	userv1 "github.com/WeiXinao/msProject/api/proto/gen/user/v1"
	"github.com/WeiXinao/msProject/pkg/respx"
	"github.com/jinzhu/copier"

	"github.com/WeiXinao/msProject/bff/internal/svc"
	"github.com/WeiXinao/msProject/bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginRsp, err error) {
	// 1. 调用 user grpc 完成登陆
	loginReq := &userv1.LoginRequest{}
	err = copier.Copy(loginReq, req)
	if err != nil {
		return nil, err
	}
	loginReq.Ip = l.ctx.Value("ip").(string)
	loginRsp, err := l.svcCtx.UserClient.Login(l.ctx, loginReq)
	if err != nil {
		return nil, respx.FromStatusErr(err)
	}
	resp = &types.LoginRsp{}
	err = copier.Copy(resp, loginRsp)
	if err != nil {
		l.Error("[Login] ", err)
	}
	return
}
