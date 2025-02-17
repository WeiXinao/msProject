package auth

import (
	"context"
	"strings"

	v1 "github.com/WeiXinao/msProject/api/proto/gen/account/v1"
	"github.com/WeiXinao/msProject/bff/internal/svc"
	"github.com/WeiXinao/msProject/bff/internal/types"
	"github.com/WeiXinao/msProject/pkg/gozerox"
	"github.com/bytedance/sonic"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type ApplyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewApplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApplyLogic {
	return &ApplyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ApplyLogic) Apply(req *types.AuthApplyReq) (resp *types.AuthApplyRsp, err error) {
	return gozerox.HttpLogicWrapper(l.ctx, l, req, func(methodName string, logic *ApplyLogic, req *types.AuthApplyReq) (*types.AuthApplyRsp, error) {
		nodes := make([]string, 0)
		if strings.TrimSpace(req.Nodes) != "" {
			err := sonic.UnmarshalString(req.Nodes, &nodes)
			if err != nil {
				return nil, err
			}
		}

		applyRep, err := logic.svcCtx.AccountClient.Apply(logic.ctx, &v1.AuthReqMessage{
			Action: req.Action,
			AuthId: req.Id,
			Nodes: nodes,
		})
		if err != nil {
			return nil, err
		}
		rsp := &types.AuthApplyRsp{}
		err = copier.CopyWithOption(rsp, applyRep, copier.Option{DeepCopy: true})
		if err != nil {
			return nil, err
		}
		return rsp, nil
	})
}
