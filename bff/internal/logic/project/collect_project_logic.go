package project

import (
	"context"
	projectv1 "github.com/WeiXinao/msProject/api/proto/gen/project/v1"
	"github.com/WeiXinao/msProject/pkg/respx"

	"github.com/WeiXinao/msProject/bff/internal/svc"
	"github.com/WeiXinao/msProject/bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CollectProjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCollectProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CollectProjectLogic {
	return &CollectProjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CollectProjectLogic) CollectProject(req *types.CollectProjectReq) (resp *types.CollectProjectRsp, err error) {
	memberId := l.ctx.Value("memberId").(int64)
	_, err = l.svcCtx.ProjectClient.UpdateCollectProject(l.ctx, &projectv1.UpdateCollectProjectRequest{
		ProjectCode: req.ProjectCode,
		CollectType: req.Type,
		MemberId:    memberId,
	})
	if err != nil {
		l.Error("[logic CollectProject]", err)
		err = respx.FromStatusErr(err)
		return
	}
	resp = &types.CollectProjectRsp{ProjectList: []*types.Project{}}
	return
}
