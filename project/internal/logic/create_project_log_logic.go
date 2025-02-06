package logic

import (
	"context"

	"github.com/WeiXinao/msProject/api/proto/gen/project/v1"
	"github.com/WeiXinao/msProject/pkg/respx"
	"github.com/WeiXinao/msProject/project/internal/domain"
	"github.com/WeiXinao/msProject/project/internal/svc"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateProjectLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateProjectLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateProjectLogLogic {
	return &CreateProjectLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateProjectLogLogic) CreateProjectLog(in *v1.CreateProjectLogRequest) (*v1.CreateProjectLogResponse, error) {
	projectLogDmn := domain.ProjectLog{}
	err := copier.Copy(&projectLogDmn, in)
	if err != nil {
		l.Errorf("[logic CreateProjectLog] %#v", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}

	err = l.svcCtx.ProjectRepo.SaveProjectLog(l.ctx, projectLogDmn)
	if err != nil {
		l.Errorf("[logic CreateProjectLog] %#v", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}

	return &v1.CreateProjectLogResponse{}, nil
}
