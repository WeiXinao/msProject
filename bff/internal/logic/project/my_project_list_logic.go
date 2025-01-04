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

type MyProjectListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMyProjectListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MyProjectListLogic {
	return &MyProjectListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MyProjectListLogic) MyProjectList(req *types.SelfListReq) (resp *types.SelfListRsp, err error) {
	keyMemberId := "memberId"
	memberId, _ := l.ctx.Value(keyMemberId).(int64)
	myProjectResp, err := l.svcCtx.ProjectClient.FindProjectByMemId(l.ctx,
		&projectv1.ProjectRequest{
			MemberId: memberId,
			Page:     req.Page,
			PageSize: req.PageSize,
		})
	if err != nil {
		l.Error("[MyProjectList]", err)
		err = respx.FromStatusErr(err)
		resp = &types.SelfListRsp{List: []*types.ProjectAndMember{}}
		return
	}
	if myProjectResp.Pm == nil {
		myProjectResp.Pm = []*projectv1.ProjectMessage{}
	}
	resp = &types.SelfListRsp{List: []*types.ProjectAndMember{}}
	err = copier.CopyWithOption(&resp.List, &myProjectResp.Pm, copier.Option{DeepCopy: true})
	if err != nil {
		l.Error("[MyProjectList]", err)
		err = respx.FromStatusErr(err)
		resp = &types.SelfListRsp{List: []*types.ProjectAndMember{}}
		return
	}
	resp.Total = myProjectResp.Total
	return
}
