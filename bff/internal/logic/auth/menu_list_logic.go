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

type MenuListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuListLogic {
	return &MenuListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MenuListLogic) MenuList() (resp *types.MenuListRsp, err error) {
	return gozerox.HttpLogicWrapperWithoutReq[*MenuListLogic, *types.MenuListRsp](l.ctx, l, func(methodName string, logic *MenuListLogic) (*types.MenuListRsp, error) {
		res, err := l.svcCtx.AccountClient.MenuList(logic.ctx, &v1.MenuRequest{})
		if err != nil {
			return nil, err
		}
		resp = &types.MenuListRsp{}
		err = copier.CopyWithOption(&resp, res, copier.Option{DeepCopy: true})
		if err != nil {
			return nil, err
		}
		if resp.List == nil {
			resp.List = []*types.MenuMessage{}
		}
		return resp, nil
	})
}
