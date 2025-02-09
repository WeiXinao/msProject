package task

import (
	"context"

	v1 "github.com/WeiXinao/msProject/api/proto/gen/file/v1"
	"github.com/WeiXinao/msProject/bff/internal/svc"
	"github.com/WeiXinao/msProject/bff/internal/types"
	"github.com/WeiXinao/msProject/pkg/respx"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskSourcesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTaskSourcesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskSourcesLogic {
	return &TaskSourcesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TaskSourcesLogic) TaskSources(req *types.TaskSourcesReq) (resp *types.TaskSourcesRsp, err error) {
	sourceListRsp, err := l.svcCtx.FileClient.TaskSources(l.ctx, &v1.TaskSourcesRequest{
		TaskCode: req.TaskCode,
	})
	if err != nil {
		l.Errorf("[logic TaskSources] %#v", err)
		return nil, respx.FromStatusErr(respx.ErrInternalServer)
	}
	resp = &types.TaskSourcesRsp{}
	err = copier.CopyWithOption(&resp, sourceListRsp, copier.Option{DeepCopy: true})
	if err != nil {
		l.Errorf("[logic TaskSources] %#v", err)
		return nil, respx.FromStatusErr(respx.ErrInternalServer)
	}
	if resp.List == nil {
		resp.List = []*types.SourceLink{}
	}
	return
}
