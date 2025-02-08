package logic

import (
	"context"

	"github.com/WeiXinao/msProject/file/internal/svc"
	"github.com/WeiXinao/msProject/api/proto/gen/file/v1"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskSourcesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTaskSourcesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskSourcesLogic {
	return &TaskSourcesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TaskSourcesLogic) TaskSources(in *v1.TaskSourcesRequest) (*v1.TaskSourceResponse, error) {
	// todo: add your logic here and delete this line

	return &v1.TaskSourceResponse{}, nil
}
