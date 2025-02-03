package task

import (
	"context"

	"github.com/WeiXinao/msProject/task/internal/domain"
	"github.com/WeiXinao/msProject/task/internal/svc"
	"github.com/bytedance/sonic"
	"github.com/zeromicro/go-zero/core/logx"
)

type SaveTaskStagesEvent struct {
	TaskStagesList []*domain.TaskStages
}

type SaveTaskStagesEventCustomer struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext	
} 

func NewSaveTaskStagesEventCustomer(ctx context.Context, svcCtx *svc.ServiceContext) *SaveTaskStagesEventCustomer {
	return &SaveTaskStagesEventCustomer{
		ctx: ctx,
		svcCtx: svcCtx,
	}
}

func (s *SaveTaskStagesEventCustomer) Consume(ctx context.Context, key, value string) error {
	evt := SaveTaskStagesEvent{}
	err := sonic.UnmarshalString(value, &evt)
	taskStagesDmns := evt.TaskStagesList
	err = s.svcCtx.TaskRepo.CreateTaskStagesList(s.ctx, taskStagesDmns)
	if err != nil {
		logx.Errorf("[Custom SaveTaskStagesEvent] %#v", err)
	}
	return err
}