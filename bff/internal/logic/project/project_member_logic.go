package project

import (
	"context"

	v1 "github.com/WeiXinao/msProject/api/proto/gen/project/v1"
	"github.com/WeiXinao/msProject/bff/internal/svc"
	"github.com/WeiXinao/msProject/bff/internal/types"
	"github.com/WeiXinao/msProject/pkg/respx"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProjectMemberLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProjectMemberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProjectMemberLogic {
	return &ProjectMemberLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProjectMemberLogic) ProjectMember(req *types.ProjectMemberListReq) (resp *types.ProjectMemberListRsp, err error) {
	projectMemberListRsp, err := l.svcCtx.ProjectClient.ProjectMemberList(l.ctx, &v1.ProjectMemberListRequest{
		MemberId: l.ctx.Value("memberId").(int64),
		ProjectCode: req.ProjectCode,
		Page: req.Page,
		PageSize: req.PageSize,	
	})
	if err != nil {
		l.Error("[logic ProjectMember] %w", err)	
		err = respx.FromStatusErr(respx.ErrInternalServer)
		return
	}
	resp = &types.ProjectMemberListRsp{
	}
	err = copier.CopyWithOption(resp, projectMemberListRsp, copier.Option{DeepCopy: true})
	if err != nil {
		l.Error("[logic ProjectMember] %w", err)	
		err = respx.FromStatusErr(respx.ErrInternalServer)
		return
	}
	if resp.List == nil {
		resp.List = []*types.ProjectListMember{}
	}
	return 
}
