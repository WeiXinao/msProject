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

type TaskLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTaskLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskLogLogic {
	return &TaskLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TaskLogLogic) TaskLog(req *types.TaskLogReq) (resp *types.TaskLogRsp, err error) {
	taskLogRsp, err := l.svcCtx.TaskClient.TaskLog(l.ctx, &taskservice.TaskLogRequest{
		TaskCode: req.TaskCode,
		MemberId: l.ctx.Value("memberId").(int64),
		Page: int64(req.Page),
		PageSize: int64(req.PageSize),
		All: int32(req.All),
		Comment: int32(req.Comment),
	})
	if err != nil {
		l.Errorf("[logic TaskLog] %#v", err)	
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}
	resp = &types.TaskLogRsp{}
	err = copier.Copy(&resp, taskLogRsp)
	if err != nil {
		l.Errorf("[logic TaskLog] %#v", err)	
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}
	if resp.List == nil {
		resp.List = []*types.TaskLog{}
	}
	resp.Page = req.Page
	return
}
