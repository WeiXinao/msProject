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

type SaveTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveTaskLogic {
	return &SaveTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveTaskLogic) SaveTask(req *types.TaskSaveReq) (resp *types.TaskDisplay, err error) {
	taskMsg, err := l.svcCtx.TaskClient.SaveTask(l.ctx, &v1.SaveTaskRequest{
		ProjectCode: req.ProjectCode,
		Name: req.Name,
		StageCode: req.StageCode,
		AssignTo: req.AssignTo,
		MemberId: l.ctx.Value("memberId").(int64),
	})
	if err != nil {
		err = respx.FromStatusErr(respx.ErrInternalServer)			
		l.Error("[logic SaveTask] %#v", err)
		return
	}
	resp = &types.TaskDisplay{}
	err = copier.CopyWithOption(&resp, taskMsg, copier.Option{DeepCopy: true})
	if err != nil {
		err = respx.ErrInternalServer
		l.Error("[logic SaveTask] %#v", err)
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
