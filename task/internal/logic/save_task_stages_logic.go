package logic

import (
	"context"

	"github.com/WeiXinao/msProject/api/proto/gen/task/v1"
	"github.com/WeiXinao/msProject/pkg/respx"
	"github.com/WeiXinao/msProject/task/internal/domain"
	"github.com/WeiXinao/msProject/task/internal/svc"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveTaskStagesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveTaskStagesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveTaskStagesLogic {
	return &SaveTaskStagesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SaveTaskStagesLogic) SaveTaskStages(in *v1.SaveTaskStagesRequest) (*v1.SaveTaskStagesResponse, error) {
	taskStagesDmns := make([]*domain.TaskStages, 0)
	err := copier.CopyWithOption(&taskStagesDmns, in.List, copier.Option{DeepCopy: true})
	if err != nil {
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}

	err = l.svcCtx.TaskRepo.CreateTaskStagesList(l.ctx, taskStagesDmns)
	if err != nil {
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}
		
	return &v1.SaveTaskStagesResponse{}, nil
}
