package project

import (
	"context"
	projectv1 "github.com/WeiXinao/msProject/api/proto/gen/project/v1"
	"github.com/WeiXinao/msProject/pkg/respx"

	"github.com/WeiXinao/msProject/bff/internal/svc"
	"github.com/WeiXinao/msProject/bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RecoverProjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRecoverProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecoverProjectLogic {
	return &RecoverProjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RecoverProjectLogic) RecoverProject(req *types.RecycleProjectReq) (resp *types.RecycleProjectRsp, err error) {
	_, err = l.svcCtx.ProjectClient.RecycleOrRecoverProject(l.ctx, &projectv1.RecycleProjectRequest{
		ProjectCode: req.ProjectCode,
		Deleted:     false,
	})
	if err != nil {
		l.Error("[logic RecoverProject]", err)
		err = respx.FromStatusErr(err)
		return
	}
	resp = &types.RecycleProjectRsp{ProjectList: []*types.Project{}}
	return
}
