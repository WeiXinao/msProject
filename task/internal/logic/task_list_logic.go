package logic

import (
	"context"
	"strconv"

	"github.com/WeiXinao/msProject/api/proto/gen/task/v1"
	"github.com/WeiXinao/msProject/pkg/respx"
	"github.com/WeiXinao/msProject/task/internal/domain"
	"github.com/WeiXinao/msProject/task/internal/svc"
	"github.com/WeiXinao/msProject/user/loginservice"
	"github.com/WeiXinao/xkit/slice"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTaskListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskListLogic {
	return &TaskListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TaskListLogic) TaskList(in *v1.TaskListRequest) (*v1.TaskListResponse, error) {
	stageCodeStr, _ := l.svcCtx.Encrypter.Decrypt(in.StageCode)
	stageCode, err := strconv.Atoi(stageCodeStr)
	if err != nil {
		l.Errorf("[logic TaskList] %#v", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}
	tasks, err := l.svcCtx.TaskRepo.FindTaskByStageCode(l.ctx, stageCode)
	if err != nil {
		l.Errorf("[logic TaskList] %#v", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}

	mIds := slice.Map(tasks, func(idx int, src *domain.Task) int64 {
		return src.AssignTo
	})
	memberInfosRsp, err := l.svcCtx.UserClient.MemberInfosById(l.ctx, &loginservice.MemberInfosByIdRequest{
		MIds: mIds,	
	})
	if err != nil {
		l.Errorf("[logic TaskList] %#v", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}
	idToMemberInfoMap := slice.ToMap(memberInfosRsp.List, func(element *loginservice.MemberMessage) int64 {
		return element.GetId()
	})

	taskDisplayList := slice.Map(tasks, func(idx int, src *domain.Task) *domain.TaskDisplay {
		display := src.ToTaskDisplay(l.svcCtx.Encrypter)
		if src.Private == 1 {
			_, has, err := l.svcCtx.TaskRepo.FindTaskMemberByTaskId(l.ctx, src.Id, in.MemberId)
			if err != nil {
				l.Errorf("[logic TaskList] %#v", err)
			}
			if has {
				display.CanRead = domain.CanRead	
			} else {
				display.CanRead = domain.NoCanRead
			}
		}
		display.Executor = domain.Executor{
			Name: idToMemberInfoMap[src.AssignTo].GetName(),
			Avatar: idToMemberInfoMap[src.AssignTo].GetAvatar(),
		}
		return display
	})

	rsp := &v1.TaskListResponse{}	
	err = copier.CopyWithOption(&rsp.List, taskDisplayList, copier.Option{DeepCopy: true})
	if err != nil {
		l.Errorf("[logic TaskList] %#v", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}
	return rsp, nil
}
