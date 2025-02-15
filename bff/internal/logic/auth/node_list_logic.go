package auth

import (
	"context"

	"github.com/WeiXinao/msProject/bff/internal/svc"
	"github.com/WeiXinao/msProject/bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type NodeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNodeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NodeListLogic {
	return &NodeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NodeListLogic) NodeList() (resp *types.NodeListRsp, err error) {
	// todo: add your logic here and delete this line

	return
}
