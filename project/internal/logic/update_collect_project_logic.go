package logic

import (
	"context"
	"github.com/WeiXinao/msProject/pkg/respx"
	"github.com/WeiXinao/msProject/project/internal/domain"
	"strconv"
	"time"

	"github.com/WeiXinao/msProject/api/proto/gen/project/v1"
	"github.com/WeiXinao/msProject/project/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCollectProjectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCollectProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCollectProjectLogic {
	return &UpdateCollectProjectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCollectProjectLogic) UpdateCollectProject(in *v1.UpdateCollectProjectRequest) (*v1.UpdateCollectProjectResponse, error) {
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
	if "collect" == in.CollectType {
		pc := domain.ProjectCollection{
			ProjectCode: projectCode,
			MemberCode:  in.MemberId,
			CreateTime:  time.Now().UnixMilli(),
		}
		err = l.svcCtx.ProjectRepo.SaveProjectCollection(l.ctx, pc)
		if err != nil {
			l.Error("[logic UpdateCollectProject]", err)
			return nil, respx.ToStatusErr(respx.FromStatusErr(err))
		}
	}
	if "cancel" == in.CollectType {
		err = l.svcCtx.ProjectRepo.DeleteProjectCollection(l.ctx, in.MemberId, projectCode)
		if err != nil {
			l.Error("[logic UpdateCollectProject]", err)
			return nil, respx.ToStatusErr(respx.FromStatusErr(err))
		}
	}
	return &v1.UpdateCollectProjectResponse{}, nil
}
