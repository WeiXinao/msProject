package logic

import (
	"context"

	"github.com/WeiXinao/msProject/api/proto/gen/task/v1"
	"github.com/WeiXinao/msProject/pkg/respx"
	"github.com/WeiXinao/msProject/task/internal/domain"
	"github.com/WeiXinao/msProject/task/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveTaskWorkTimeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveTaskWorkTimeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveTaskWorkTimeLogic {
	return &SaveTaskWorkTimeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SaveTaskWorkTimeLogic) SaveTaskWorkTime(in *v1.SaveTaskWorkTimeRequest) (*v1.SaveTaskWorkTimeResponse, error) {
	taskCode, err := l.svcCtx.Encrypter.DecryptInt64(in.TaskCode)
	if err != nil {
		l.Error("[logic SaveTaskWorkTime] %#v", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}
	twt := domain.TaskWorkTime{
		BeginTime: in.GetBeginTime(),
		Num: int(in.GetNum()),
		Content: in.GetContent(),
		TaskCode: taskCode,
		MemberCode: in.GetMemberId(),
	}
	err = l.svcCtx.TaskRepo.SaveTaskWorkTime(l.ctx, twt)
	if err != nil {
		l.Error("[logic SaveTaskWorkTime] %#v", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}

	return &v1.SaveTaskWorkTimeResponse{}, nil
}
