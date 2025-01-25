package logic

import (
	"context"
	"github.com/WeiXinao/msProject/pkg/respx"
	"github.com/WeiXinao/msProject/project/internal/domain"
	"strconv"

	"github.com/WeiXinao/msProject/api/proto/gen/project/v1"
	"github.com/WeiXinao/msProject/project/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProjectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProjectLogic {
	return &UpdateProjectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateProjectLogic) UpdateProject(in *v1.UpdateProjectRequest) (*v1.UpdateProjectResponse, error) {
	projectCodeStr, err := l.svcCtx.Encrypter.Decrypt(in.ProjectCode)
	if err != nil {
		l.Error("[logic UpdateCollectProject]", err)
		return nil, respx.ToStatusErr(respx.FromStatusErr(err))
	}
	projectCode, err := strconv.ParseInt(projectCodeStr, 10, 64)
	if err != nil {
		l.Error("[logic UpdateCollectProject]", err)
		return nil, respx.ToStatusErr(respx.FromStatusErr(err))
	}
	project := domain.Project{
		Id:                 projectCode,
		Cover:              in.Cover,
		Name:               in.Name,
		Description:        in.Description,
		Schedule:           in.Schedule,
		Private:            int(in.Private),
		Prefix:             in.Prefix,
		OpenPrefix:         int(in.OpenPrefix),
		OpenBeginTime:      int(in.OpenBeginTime),
		OpenTaskPrivate:    int(in.OpenTaskPrivate),
		TaskBoardTheme:     in.TaskBoardTheme,
		AutoUpdateSchedule: int(in.AutoUpdateSchedule),
	}
	err = l.svcCtx.ProjectRepo.UpdateProject(l.ctx, project)
	if err != nil {
		l.Error("[logic UpdateProject]", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}

	return &v1.UpdateProjectResponse{}, nil
}
