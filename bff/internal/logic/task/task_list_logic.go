package task

import (
	"context"

	v1 "github.com/WeiXinao/msProject/api/proto/gen/task/v1"
	"github.com/WeiXinao/msProject/bff/internal/svc"
	"github.com/WeiXinao/msProject/bff/internal/types"
	"github.com/WeiXinao/msProject/pkg/respx"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTaskListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskListLogic {
	return &TaskListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TaskListLogic) TaskList(req *types.TaskReq) (resp *types.TaskRsp, err error) {
	taskListRsp, err := l.svcCtx.TaskClient.TaskList(l.ctx, &v1.TaskListRequest{
		StageCode: req.StageCode,
	})
	if err != nil {
		err = respx.FromStatusErr(err)
		return
	}
	resp = &types.TaskRsp{}
	err = copier.CopyWithOption(&resp.List, taskListRsp.List, copier.Option{DeepCopy: true})
	if err != nil {
		err = respx.ErrInternalServer
		return
	}
	if resp.List == nil {
		resp.List = []*types.TaskDisplay{}
	}
	for i, v := range resp.List {
		if v.Tags == nil {
			resp.List[i].Tags = []int{}
		}
		if v.ChildCount == nil {
			resp.List[i].ChildCount = []int{}
		}	
	}
	return
}
