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

type MyTaskListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMyTaskListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MyTaskListLogic {
	return &MyTaskListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MyTaskListLogic) MyTaskList(req *types.MyTaskListReq) (resp *types.MyTaskListRsp, err error) {
	myTaskListRsp, err := l.svcCtx.TaskClient.MyTaskList(l.ctx, &v1.MyTaskListRequest{
		MemberId: l.ctx.Value("memberId").(int64),
		Page: req.Page,
		PageSize: req.PageSize,
		TaskType: int32(req.TaskType),
		Type: int32(req.Type),
	})
	if err != nil {
		err = respx.FromStatusErr(err)
		l.Errorf("[logic MyTaskList] %#v", err)
		return
	}
	resp = &types.MyTaskListRsp{}
	err = copier.CopyWithOption(resp, myTaskListRsp, copier.Option{DeepCopy: true})
	if err != nil {
		err = respx.FromStatusErr(err)
		l.Errorf("[logic MyTaskList] %#v", err)
		return
	}
	if resp.List == nil {
		resp.List = []*types.MyTaskDisplay{}
	}
	for _, r := range resp.List {
		r.ProjectInfo = types.ProjectInfo{
			Name: r.ProjectName,
			Code: r.ProjectCode,
		}
	}
	return
}
