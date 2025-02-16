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

type ApplyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewApplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApplyLogic {
	return &ApplyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ApplyLogic) Apply(in *v1.AuthReqMessage) (*v1.ApplyResponse, error) {
	return gozerox.RpcLogicWrapper(l.ctx, l, in, func(methodName string, logic *ApplyLogic, req *v1.AuthReqMessage) (*v1.ApplyResponse, error) {
		res := &v1.ApplyResponse{}
		if in.Action == "getnode" {
			pns, err := l.svcCtx.AccoutRepo.FindAllProjectNodes(logic.ctx)
			if err != nil {
				return nil, err
			}
			authNodes, err := l.svcCtx.AccoutRepo.FindAuthNodeStringList(l.ctx, in.GetAuthId())
			if err != nil {
				return nil, err
			}
			pnats := domain.ToAuthNodeTreeList(pns, authNodes)
			pnms := []*v1.ProjectNodeMessage{}
			err = copier.CopyWithOption(&pnms, pnats, copier.Option{DeepCopy: true})
			if err != nil {
				return nil, err
			}
			res = &v1.ApplyResponse{
				List: pnms,
				CheckedList: authNodes,
			}
		}
		return res, nil
	})
}
