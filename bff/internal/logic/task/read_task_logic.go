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

type ReadTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReadTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReadTaskLogic {
	return &ReadTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReadTaskLogic) ReadTask(req *types.ReadTaskReq) (resp *types.TaskDisplay, err error) {
	taskMsg, err := l.svcCtx.TaskClient.ReadTask(l.ctx, &v1.ReadTaskRequest{
		TaskCode: req.TaskCode,
		MemberId: l.ctx.Value("memberId").(int64),
	})
	if err != nil {
		err = respx.FromStatusErr(err)
		l.Error("[logic ReadTask] %#v", err)
		return
	}

	resp = &types.TaskDisplay{}
	err = copier.Copy(resp, taskMsg)
	if err != nil {
		l.Error("[logic ReadTask] %#v", err)
		err = respx.ErrInternalServer
		return
	}

	if resp != nil {
		if resp.Tags == nil {
			resp.Tags = []int{}
		}
		if resp.ChildCount == nil {
			resp.ChildCount = []int{}
		}
	}

	return
}
