package project

import (
	"context"
	projectv1 "github.com/WeiXinao/msProject/api/proto/gen/project/v1"
	"github.com/WeiXinao/msProject/pkg/respx"
	"github.com/jinzhu/copier"

	"github.com/WeiXinao/msProject/bff/internal/svc"
	"github.com/WeiXinao/msProject/bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditProjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEditProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditProjectLogic {
	return &EditProjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EditProjectLogic) EditProject(req *types.EditProjectReq) (resp *types.EditProjectRsp, err error) {
	memberId := l.ctx.Value("memberId").(int64)
	request := &projectv1.UpdateProjectRequest{}
	err = copier.Copy(request, req)
	if err != nil {
		l.Error("logic EditProject", err)
		err = respx.ErrInternalServer
		return
	}
	request.MemberId = memberId

	_, err = l.svcCtx.ProjectClient.UpdateProject(l.ctx, request)
	if err != nil {
		l.Error("logic EditProject", err)
		err = respx.FromStatusErr(err)
		return
	}
	resp = &types.EditProjectRsp{
		ProjectList: []*types.Project{},
	}
	return
}
