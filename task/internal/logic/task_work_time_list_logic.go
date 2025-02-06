package logic

import (
	"context"

	"github.com/WeiXinao/msProject/api/proto/gen/task/v1"
	"github.com/WeiXinao/msProject/pkg/respx"
	"github.com/WeiXinao/msProject/task/internal/domain"
	"github.com/WeiXinao/msProject/task/internal/svc"
	"github.com/WeiXinao/msProject/user/loginservice"
	"github.com/WeiXinao/xkit/slice"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskWorkTimeListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTaskWorkTimeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskWorkTimeListLogic {
	return &TaskWorkTimeListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TaskWorkTimeListLogic) TaskWorkTimeList(in *v1.TaskWorkTimeRequest) (*v1.TaskWorkTimeResponse, error) {
	taskCode, err := l.svcCtx.Encrypter.DecryptInt64(in.TaskCode)
	if err != nil {
		l.Errorf("[logic TaskWorkTimeList] %#v", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}

	workTimeList, total, err := l.svcCtx.TaskRepo.FindWorkTimeListByTaskId(l.ctx, taskCode)
	if err != nil {
		l.Errorf("[logic TaskWorkTimeList] %#v", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}
	mids := slice.Map(workTimeList, func(idx int, src *domain.TaskWorkTime) int64 {
		return src.MemberCode 
	})

	memberInfosRsp, err := l.svcCtx.UserClient.MemberInfosById(l.ctx, &loginservice.MemberInfosByIdRequest{
		MIds: mids,
	})
	if err != nil {
		l.Errorf("[logic TaskWorkTimeList] %#v", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}
	idMapMemberMsg := slice.ToMap(memberInfosRsp.GetList(), func(element *loginservice.MemberMessage) int64 {
		return element.GetId()
	})

	workTimeDisplayList := slice.Map(workTimeList, func(idx int, src *domain.TaskWorkTime) *domain.TaskWorkTimeDisplay {
		display := src.ToDisplay(l.svcCtx.Encrypter)
		memberInfo := idMapMemberMsg[src.MemberCode]
		member := domain.Member{
			Id: memberInfo.GetId(),
			Name: memberInfo.GetName(),
			Avatar: memberInfo.GetAvatar(),
			Code: memberInfo.GetCode(),
		}
		display.Member = member
		return display
	})

	list := make([]*v1.TaskWorkTime, 0)
	err = copier.Copy(&list, workTimeDisplayList)
	if err != nil {
		l.Errorf("[logic TaskWorkTimeList] %#v", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}
	
	return &v1.TaskWorkTimeResponse{
		List: list,
		Total: total,
	}, nil
}
