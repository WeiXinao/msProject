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

type TaskStagesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTaskStagesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskStagesLogic {
	return &TaskStagesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TaskStagesLogic) TaskStages(req *types.TaskStagesReq) (resp *types.TaskStagesResp, err error) {
	// 1. 调用 rpc 服务
	taskRsp, err := l.svcCtx.TaskClient.TaskStages(l.ctx, &v1.TaskStagesRequest{
		MemberId: l.ctx.Value("memberId").(int64),
		ProjectCode: req.ProjectCode,
		Page: req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		l.Errorf("[TaskStages logic] %#v", err)
		return nil, respx.FromStatusErr(respx.ErrInternalServer)
	}

	// 2. 处理响应
	resp = &types.TaskStagesResp{}
	err = copier.CopyWithOption(resp, taskRsp, copier.Option{DeepCopy: true})
	if err != nil {
		l.Errorf("[TaskStages logic] %#v", err)
		return nil, respx.FromStatusErr(respx.ErrInternalServer)
	}
	if resp.List == nil {
		resp.List = []*types.TaskStages{}
	}

	for i := range resp.List {
		resp.List[i].TasksLoading = true
		resp.List[i].FixedCreator = false
		resp.List[i].ShowTaskCard = false
		resp.List[i].Tasks = []int{}
		resp.List[i].DoneTasks = []int{}
		resp.List[i].UnDoneTasks = []int{}
	}

	return
}
