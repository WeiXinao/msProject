package logic

import (
	"context"

	"github.com/WeiXinao/msProject/account/internal/svc"
	"github.com/WeiXinao/msProject/api/proto/gen/account/v1"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAuthListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthListLogic {
	return &AuthListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AuthListLogic) AuthList(in *v1.AuthListRequest) (*v1.AuthListResponse, error) {
	// todo: add your logic here and delete this line

	return &v1.AuthListResponse{}, nil
}
