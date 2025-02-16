package auth

import (
	"context"

	v1 "github.com/WeiXinao/msProject/api/proto/gen/account/v1"
	"github.com/WeiXinao/msProject/bff/internal/svc"
	"github.com/WeiXinao/msProject/bff/internal/types"
	"github.com/WeiXinao/msProject/pkg/gozerox"
	"github.com/jinzhu/copier"

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
	return gozerox.HttpLogicWrapperWithoutReq[*NodeListLogic, *types.NodeListRsp](l.ctx, l, func(methodName string, logic *NodeListLogic) (*types.NodeListRsp, error) {
		rsp, err := l.svcCtx.AccountClient.NodeList(logic.ctx, &v1.NodeListRequest{})
		if err != nil {
			return nil, err
		}
		res := &types.NodeListRsp{}
		err = copier.CopyWithOption(res, rsp, copier.Option{DeepCopy: true})
		if err != nil {
			return nil, err
		}
		return res, nil
	})
}
