package auth

import (
	"context"

	"github.com/WeiXinao/msProject/bff/internal/svc"
	"github.com/WeiXinao/msProject/bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAuthListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthListLogic {
	return &AuthListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AuthListLogic) AuthList() (resp *types.AuthListRsp, err error) {
	// todo: add your logic here and delete this line

	return
}
