package auth

import (
	"context"

	v1 "github.com/WeiXinao/msProject/api/proto/gen/account/v1"
	"github.com/WeiXinao/msProject/bff/internal/middleware"
	"github.com/WeiXinao/msProject/bff/internal/svc"
	"github.com/WeiXinao/msProject/bff/internal/types"
	"github.com/WeiXinao/msProject/pkg/gozerox"
	"github.com/jinzhu/copier"

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

func (l *AuthListLogic) AuthList(req *types.AuthListReq) (resp *types.AuthListRsp, err error) {
	return gozerox.HttpLogicWrapper(l.ctx, l, req, func(methodName string, logic *AuthListLogic, req *types.AuthListReq) (*types.AuthListRsp, error) {
		organCode := l.ctx.Value(middleware.KeyOrganizationCode).(string)
		resp, err := l.svcCtx.AccountClient.AuthList(logic.ctx, &v1.AuthListRequest{
			OrganizationCode: organCode,
			Page: req.Page,
			PageSize: req.PageSize,
		})
		if err != nil {
			return nil, err
		}
		res := &types.AuthListRsp{}
		err = copier.Copy(&res, resp)
		if err != nil {
			return nil, err
		}
		if res.List == nil {
			res.List = []*types.ProjectAuth{}
		}
		res.Page = req.Page
		return res, nil
	})
}
