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

type ProjectTemplateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProjectTemplateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProjectTemplateLogic {
	return &ProjectTemplateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProjectTemplateLogic) ProjectTemplate(req *types.ProjectTemplateReq) (resp *types.ProjectTemplateRsp, err error) {
	memberId := l.ctx.Value("memberId").(int64)
	memberName := l.ctx.Value("memberName").(string)
	organizationCode := l.ctx.Value("organizationCode").(string)

	ptmRsp, err := l.svcCtx.ProjectClient.FindProjectTemplate(l.ctx, &projectv1.FindProjectTemplateRequest{
		MemberId:         memberId,
		MemberName:       memberName,
		OrganizationCode: organizationCode,
		Page:             req.Page,
		PageSize:         req.PageSize,
		ViewType:         req.ViewType,
	})
	if err != nil {
		return &types.ProjectTemplateRsp{Ptm: []*types.ProjectTemplate{}}, respx.FromStatusErr(err)
	}

	var ptms []*types.ProjectTemplate
	err = copier.CopyWithOption(&ptms, &ptmRsp.Ptm, copier.Option{DeepCopy: true})
	if err != nil {
		l.Error("[logic ProjectTemplate]", err)
		return &types.ProjectTemplateRsp{Ptm: []*types.ProjectTemplate{}}, respx.ErrInternalServer
	}
	if ptms == nil {
		ptms = []*types.ProjectTemplate{}
	}
	for _, v := range ptms {
		if v.TaskStages == nil {
			v.TaskStages = []*types.TaskStagesOnlyName{}
		}
	}
	resp = &types.ProjectTemplateRsp{
		Ptm:   ptms,
		Total: ptmRsp.Total,
	}
	return
}
