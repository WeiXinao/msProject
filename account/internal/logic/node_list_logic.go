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

type NodeListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewNodeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NodeListLogic {
	return &NodeListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *NodeListLogic) NodeList(in *v1.NodeListRequest) (*v1.ProjectNodeResponse, error) {
	return gozerox.HttpLogicWrapperWithoutReq(l.ctx, l, func(methodName string, logic *NodeListLogic) (*v1.ProjectNodeResponse, error) {
		// node 表都查出来
		pns, err := l.svcCtx.AccoutRepo.FindAllProjectNodes(logic.ctx)
		if err != nil {
			return nil, err
		}

		// 转换成 treelist 结构
		nodeTreeList := domain.ToNodeTreeList(pns)
		nodes := []*v1.ProjectNodeMessage{}
		err = copier.Copy(&nodes, nodeTreeList)
		if err != nil {
			return nil, err
		}
		return &v1.ProjectNodeResponse{
			Nodes: nodes,
		}, nil
	})
}
