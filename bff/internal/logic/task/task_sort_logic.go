package task

import (
	"context"

	v1 "github.com/WeiXinao/msProject/api/proto/gen/task/v1"
	"github.com/WeiXinao/msProject/bff/internal/svc"
	"github.com/WeiXinao/msProject/bff/internal/types"
	"github.com/WeiXinao/msProject/pkg/respx"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskSortLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTaskSortLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskSortLogic {
	return &TaskSortLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TaskSortLogic) TaskSort(req *types.TaskSortReq) (resp *types.TaskSortRsp, err error) {
	_, err = l.svcCtx.TaskClient.TaskSort(l.ctx, &v1.TaskSortRequest{
		PreTaskCode: req.PreTaskCode,
		NextTaskCode: req.NextTaskCode,
		ToStageCode: req.ToStageCode,
	})
	if err != nil {
		err = respx.FromStatusErr(err)
		l.Error("[logic TaskSort] %#v", err)
	}
	resp = &types.TaskSortRsp{}
	return
}
