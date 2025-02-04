package logic

import (
	"context"

	"github.com/WeiXinao/msProject/api/proto/gen/project/v1"
	"github.com/WeiXinao/msProject/pkg/respx"
	"github.com/WeiXinao/msProject/project/internal/svc"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindProjectByIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindProjectByIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindProjectByIdsLogic {
	return &FindProjectByIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindProjectByIdsLogic) FindProjectByIds(in *v1.FindProjectByIdsRequest) (*v1.FindProjectByIdsResponse, error) {
	projects, err := l.svcCtx.ProjectRepo.FindProjectByIds(l.ctx, in.ProjectCodes)
	if err != nil {
		l.Error("[logic FindProjectByIds] %#v", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}
	rsp := &v1.FindProjectByIdsResponse{}
	rsp.Projects = []*v1.ProjectMessage{}
	err = copier.CopyWithOption(&rsp.Projects, projects, copier.Option{DeepCopy: true})
	if err != nil {
		l.Error("[logic FindProjectByIds] %#v", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}
	return rsp, nil
}
