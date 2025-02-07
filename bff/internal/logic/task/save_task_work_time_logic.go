package task

import (
	"context"

	"github.com/WeiXinao/msProject/bff/internal/svc"
	"github.com/WeiXinao/msProject/bff/internal/types"
	"github.com/WeiXinao/msProject/pkg/formatx"
	"github.com/WeiXinao/msProject/pkg/respx"
	"github.com/WeiXinao/msProject/task/taskservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveTaskWorkTimeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveTaskWorkTimeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveTaskWorkTimeLogic {
	return &SaveTaskWorkTimeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveTaskWorkTimeLogic) SaveTaskWorkTime(req *types.SaveTaskWorkTimeReq) (resp *types.SaveTaskWorkTimeRsp, err error) {
	beginTime, err := formatx.ParseDateTimeString(req.BeginTime)
	if err != nil {
		l.Error("[logic SaveTaskWorkTime] %#v", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}
	_, err = l.svcCtx.TaskClient.SaveTaskWorkTime(l.ctx, &taskservice.SaveTaskWorkTimeRequest{
		TaskCode: req.TaskCode,
		MemberId: l.ctx.Value("memberId").(int64),
		Content: req.Content,
		Num: int32(req.Num),
		BeginTime: beginTime,
	})
	if err != nil {
		l.Error("[logic SaveTaskWorkTime] %#v", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}

	resp = &types.SaveTaskWorkTimeRsp{List: []int{}}

	return
}
