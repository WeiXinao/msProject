package logic

import (
	"context"

	projectv1 "github.com/WeiXinao/msProject/api/proto/gen/project/v1"
	"github.com/WeiXinao/msProject/api/proto/gen/task/v1"
	userv1 "github.com/WeiXinao/msProject/api/proto/gen/user/v1"
	"github.com/WeiXinao/msProject/pkg/respx"
	"github.com/WeiXinao/msProject/task/internal/domain"
	"github.com/WeiXinao/msProject/task/internal/svc"
	"github.com/WeiXinao/xkit/slice"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type MyTaskListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMyTaskListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MyTaskListLogic {
	return &MyTaskListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MyTaskListLogic) MyTaskList(in *v1.MyTaskListRequest) (*v1.MyTaskListResponse, error) {
	var (
		tasks []*domain.Task
		err error
		total int64
	)
	switch in.TaskType {
	case 1:
		tasks, total, err = l.svcCtx.TaskRepo.FindTaskByAssignTo(l.ctx, in.GetMemberId(), int(in.GetType()), in.GetPage(), in.GetPageSize())
	case 2:
		tasks, total, err = l.svcCtx.TaskRepo.FindTaskByMemberCode(l.ctx, in.GetMemberId(), int(in.GetType()), in.GetPage(), in.GetPageSize())
	case 3:
		tasks, total, err = l.svcCtx.TaskRepo.FindTaskByCreateBy(l.ctx, in.GetMemberId(), int(in.GetType()), in.GetPage(), in.GetPageSize())
	}
	if err != nil {
		l.Error("[logic MyTaskList] %#v", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}
	if tasks == nil || len(tasks) <= 0 {
		return &v1.MyTaskListResponse{List: nil, Total: 0}, nil
	}

	pids := slice.Map(tasks, func(idx int, src *domain.Task) int64 {
		return src.ProjectCode
	})
	mids := slice.Map(tasks, func(idx int, src *domain.Task) int64 {
		return src.AssignTo
	})
	projectsRsp, err := l.svcCtx.ProjectClient.FindProjectByIds(l.ctx, &projectv1.FindProjectByIdsRequest{
		ProjectCodes: pids,
	})
	if err != nil {
		l.Error("[logic MyTaskList] %#v", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}
	memberInfosRsp, err := l.svcCtx.UserClient.MemberInfosById(l.ctx, &userv1.MemberInfosByIdRequest{
		MIds: mids,
	})
	if err != nil {
		l.Error("[logic MyTaskList] %#v", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}
	membersMap := slice.ToMap(memberInfosRsp.List, func(element *userv1.MemberMessage) int64 {
		return element.GetId()
	})
	projectsMap := slice.ToMap(projectsRsp.Projects, func(element *projectv1.ProjectMessage) int64 {
		return element.GetId()
	})

	myTaskDisplayList := slice.Map(tasks, func(idx int, src *domain.Task) *domain.MyTaskDisplay {
		memberMsg := membersMap[src.AssignTo]
		name := memberMsg.GetName()
		avatar := memberMsg.GetAvatar()
		project := domain.Project{}
		err = copier.Copy(&project, projectsMap[src.ProjectCode])
		if err != nil {
			l.Error("[logic MyTaskList] %#v", err)
			return nil
		}
		return src.ToMyTaskDisplay(l.svcCtx.Encrypter, project, name, avatar)
	})
	
	myTaskMsgs := make([]*v1.MyTaskMessage, 0)
	err = copier.Copy(&myTaskMsgs, myTaskDisplayList)
	if err != nil {
		l.Error("[logic MyTaskList] %#v", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}
	
	return &v1.MyTaskListResponse{
		List: myTaskMsgs,
		Total: total,
	}, nil
}
