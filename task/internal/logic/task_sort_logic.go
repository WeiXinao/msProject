package logic

import (
	"context"
	"strings"

	"github.com/WeiXinao/msProject/api/proto/gen/task/v1"
	"github.com/WeiXinao/msProject/pkg/respx"
	"github.com/WeiXinao/msProject/task/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskSortLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTaskSortLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskSortLogic {
	return &TaskSortLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TaskSortLogic) TaskSort(in *v1.TaskSortRequest) (*v1.TaskSortResponse, error) {
	// 移动的任务 id 肯定有 preTaskCode
	preTaskCode, err := l.svcCtx.Encrypter.DecryptInt64(in.PreTaskCode)
	if err != nil {
		l.Errorf("[logic TaskSort] %#v", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}
	nextTaskCode := int64(0)
	if len(strings.TrimSpace(in.NextTaskCode)) == 0 {
		nextTaskCode = 0
	} else {
		nextTaskCode, err = l.svcCtx.Encrypter.DecryptInt64(in.NextTaskCode)
		if err != nil {
			l.Errorf("[logic TaskSort] %#v", err)
			return nil, respx.ToStatusErr(respx.ErrInternalServer)
		}
	}
	toStageCode, err := l.svcCtx.Encrypter.DecryptInt64(in.ToStageCode)
	if err != nil {
		l.Errorf("[logic TaskSort] %#v", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}
	task, err := l.svcCtx.TaskRepo.FindTaskById(l.ctx, preTaskCode)
	if err != nil {
		l.Errorf("[logic TaskSort] %#v", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}
	nextTask, err := l.svcCtx.TaskRepo.FindTaskById(l.ctx, nextTaskCode)
	if err != nil {
		l.Errorf("[logic TaskSort] %#v", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}
	err = l.svcCtx.TaskRepo.Move(l.ctx, int(toStageCode), task, nextTask)
	if err != nil {
		l.Errorf("[logic TaskSort] %#v", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}

	return &v1.TaskSortResponse{}, nil
}
