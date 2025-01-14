package project

import (
	"context"
	projectv1 "github.com/WeiXinao/msProject/api/proto/gen/project/v1"
	"github.com/WeiXinao/msProject/bff/internal/svc"
	"github.com/WeiXinao/msProject/bff/internal/types"
	"github.com/WeiXinao/msProject/pkg/respx"

	"github.com/zeromicro/go-zero/core/logx"
)

type RecycleProjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRecycleProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecycleProjectLogic {
	return &RecycleProjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RecycleProjectLogic) RecycleProject(req *types.RecycleProjectReq) (resp *types.RecycleProjectRsp, err error) {
	_, err = l.svcCtx.ProjectClient.RecycleOrRecoverProject(l.ctx, &projectv1.RecycleProjectRequest{
		ProjectCode: req.ProjectCode,
		Deleted:     true,
	})
	if err != nil {
		l.Error("[logic RecycleProject]", err)
		err = respx.FromStatusErr(err)
		return
	}
	resp = &types.RecycleProjectRsp{ProjectList: []*types.Project{}}
	return
}
