package task

import (
	"context"

	v1 "github.com/WeiXinao/msProject/api/proto/gen/task/v1"
	"github.com/WeiXinao/msProject/bff/internal/svc"
	"github.com/WeiXinao/msProject/bff/internal/types"
	"github.com/WeiXinao/msProject/pkg/respx"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListTaskMemberLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListTaskMemberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListTaskMemberLogic {
	return &ListTaskMemberLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListTaskMemberLogic) ListTaskMember(req *types.ListTaskMemberReq) (resp *types.ListTaskMemberRsp, err error) {
	taskMemberResp, err := l.svcCtx.TaskClient.ListTaskMember(l.ctx, &v1.ListTaskMemberRequest{
		TaskCode: req.TaskCode,
		MemberId: l.ctx.Value("memberId").(int64),
		Page: req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		err = respx.FromStatusErr(err)	
		l.Errorf("[logic ListTaskMember] %#v", err)
		return
	}
	resp = &types.ListTaskMemberRsp{}
	err = copier.Copy(&resp, taskMemberResp)
	if err != nil {
		l.Errorf("[logic ListTaskMember] %#v", err)
		err = respx.ErrInternalServer
		return
	}
	if resp.List == nil {
		resp.List = []*types.TaskMember{}
	}
	resp.Page = req.Page
	return
}
