package logic

import (
	"context"

	"github.com/WeiXinao/msProject/account/internal/domain"
	"github.com/WeiXinao/msProject/account/internal/svc"
	"github.com/WeiXinao/msProject/api/proto/gen/account/v1"
	"github.com/WeiXinao/msProject/pkg/gozerox"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuListLogic {
	return &MenuListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MenuListLogic) MenuList(in *v1.MenuRequest) (*v1.MenuResponse, error) {
	return gozerox.RpcLogicWrapper(l.ctx, l, in, func(methodName string, logic *MenuListLogic, req *v1.MenuRequest) (*v1.MenuResponse, error) {
		pms, err := l.svcCtx.AccoutRepo.GetMenus(logic.ctx)
		if err != nil {
			return nil, err
		}
		pmc := domain.CovertChild(pms)
		menuList := []*v1.MenuMessage{}
		err = copier.CopyWithOption(&menuList, pmc, copier.Option{DeepCopy: true})
		if err != nil {
			return nil, err
		}
		return &v1.MenuResponse{
			List: menuList,
		}, nil
	})
}
