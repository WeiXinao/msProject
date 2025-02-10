package task

import (
	"context"

	"github.com/WeiXinao/msProject/bff/internal/svc"
	"github.com/WeiXinao/msProject/bff/internal/types"
	"github.com/WeiXinao/msProject/pkg/respx"
	"github.com/WeiXinao/msProject/task/taskservice"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLogBySelfProjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLogBySelfProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLogBySelfProjectLogic {
	return &GetLogBySelfProjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLogBySelfProjectLogic) GetLogBySelfProject(req *types.GetLogBySelfProjectReq) (resp *types.GetLogBySelfProjectRsp, err error) {
	getLogBySelfProjectRsp, err := l.svcCtx.TaskClient.GetLogBySelfProject(l.ctx, &taskservice.GetLogBySelfProjectRequest{
		MemberId: l.ctx.Value("memberId").(int64),
		Page: req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		err = respx.FromStatusErr(err)
		l.Errorf("[logic GetLogBySelfProject] %#v", err)
		return 
	}
	resp = &types.GetLogBySelfProjectRsp{}
	err = copier.CopyWithOption(resp, getLogBySelfProjectRsp, copier.Option{DeepCopy: true})
	if err != nil {
		err = respx.FromStatusErr(err)
		l.Errorf("[logic GetLogBySelfProject] %#v", err)
		return 
	}
	if resp.List == nil {
		resp.List = []*types.ProjectLog{}
	}

	return
}
