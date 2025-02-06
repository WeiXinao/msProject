package task

import (
	"context"

	"github.com/WeiXinao/msProject/bff/internal/svc"
	"github.com/WeiXinao/msProject/bff/internal/types"
	"github.com/WeiXinao/msProject/pkg/respx"
	"github.com/WeiXinao/msProject/task/taskservice"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskWorkTimeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTaskWorkTimeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskWorkTimeListLogic {
	return &TaskWorkTimeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TaskWorkTimeListLogic) TaskWorkTimeList(req *types.TaskWorkTimeListReq) (resp *types.TaskWorkTimeListRsp, err error) {
	taskWorkTimeRsp, err := l.svcCtx.TaskClient.TaskWorkTimeList(l.ctx, &taskservice.TaskWorkTimeRequest{
		TaskCode: req.TaskCode,	
		MemberId: l.ctx.Value("memberId").(int64),
	})	
	if err != nil {
		err = respx.FromStatusErr(err)
		l.Errorf("[logic TaskWorkTimeList] %#v", err)
		return
	}
	resp = &types.TaskWorkTimeListRsp{}
	err = copier.Copy(resp, taskWorkTimeRsp)
	if err != nil {
		err = respx.FromStatusErr(err)
		l.Errorf("[logic TaskWorkTimeList] %#v", err)
		return
	}
	
	if resp.List == nil {
		resp.List = []*types.TaskWorkTime{}
	}

	return
}
