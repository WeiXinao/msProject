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

type ProjectSaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProjectSaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProjectSaveLogic {
	return &ProjectSaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProjectSaveLogic) ProjectSave(req *types.ProjectSaveReq) (resp *types.ProjectSaveRsp, err error) {
	memerId := l.ctx.Value("memberId").(int64)
	organizationCode := l.ctx.Value("organizationCode").(string)
	saveProjectRsp, err := l.svcCtx.ProjectClient.SaveProject(l.ctx, &projectv1.SaveProjectReq{
		MemberId:         memerId,
		OrganizationCode: organizationCode,
		Name:             req.Name,
		TemplateCode:     req.TemplateCode,
		Description:      req.Description,
		Id:               int64(req.Id),
	})
	if err != nil {
		l.Error("[ProjectSave]", err)
		err = respx.FromStatusErr(err)
		return
	}
	resp = &types.ProjectSaveRsp{}
	err = copier.Copy(&resp, &saveProjectRsp)
	if err != nil {
		err = respx.ErrInternalServer
	}
	return
}
