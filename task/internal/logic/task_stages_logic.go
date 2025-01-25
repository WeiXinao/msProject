package logic

import (
	"context"
	"strconv"
	"time"

	"github.com/WeiXinao/msProject/api/proto/gen/task/v1"
	"github.com/WeiXinao/msProject/pkg/respx"
	"github.com/WeiXinao/msProject/task/internal/domain"
	"github.com/WeiXinao/msProject/task/internal/svc"
	"github.com/WeiXinao/xkit/slice"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskStagesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTaskStagesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskStagesLogic {
	return &TaskStagesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TaskStagesLogic) TaskStages(in *v1.TaskStagesRequest) (*v1.TaskStagesResponse, error) {
	projectCodeStr, err := l.svcCtx.Encrypter.Decrypt(in.ProjectCode)
	if err != nil {
		l.Errorf("[logic TaskStages] %#v", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}

	projectCode, err := strconv.ParseInt(projectCodeStr, 10, 64)
	if err != nil {
		l.Errorf("[logic TaskStages] %#v", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}
	page := in.Page
	pageSize := in.PageSize
	ts, total, err := l.svcCtx.TaskRepo.FindStagesByProjectIdPagination(l.ctx, projectCode, page, pageSize)
	if err != nil {
		l.Errorf("[logic TaskStages] %#v", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}

	taskStagesMsgs := slice.Map(ts, func(idx int, src *domain.TaskStages) *v1.TaskStagesMessage {
		code, err := l.svcCtx.Encrypter.EncryptInt64(int64(src.Id))
		if err != nil {
			l.Error("[logic TaskStages] %#w", err)	
		}
		return &v1.TaskStagesMessage{
			Code: code,
			Name: src.Name,
			ProjectCode: strconv.FormatInt(src.ProjectCode, 10),
			Sort: int32(src.Sort),
			Description: src.Description,
			CreateTime: time.UnixMilli(src.CreateTime).Format(time.DateTime),
			Deleted: int32(src.Deleted),
			Id: int32(src.Id),
		}
	})
	return &v1.TaskStagesResponse{
		List: taskStagesMsgs,
		Total: total,
	}, nil
}