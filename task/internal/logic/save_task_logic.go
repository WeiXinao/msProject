package logic

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/WeiXinao/msProject/api/proto/gen/task/v1"
	"github.com/WeiXinao/msProject/pkg/respx"
	"github.com/WeiXinao/msProject/project/projectservice"
	"github.com/WeiXinao/msProject/task/internal/domain"
	"github.com/WeiXinao/msProject/task/internal/svc"
	"github.com/WeiXinao/msProject/user/loginservice"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveTaskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveTaskLogic {
	return &SaveTaskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SaveTaskLogic) SaveTask(in *v1.SaveTaskRequest) (*v1.TaskMessage, error) {
	// 参数校验
	if strings.TrimSpace(in.Name) == "" {
		err := respx.ErrEmptyTaskName 
		l.Errorf("[logic SaveTask] %#v", err)
		return nil, respx.ToStatusErr(err)
	}

	stageCodeStr, err := l.svcCtx.Encrypter.Decrypt(in.StageCode)
	if err != nil {
		l.Errorf("[logic SaveTask] %#v", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}	
	stageCode, err := strconv.ParseInt(stageCodeStr, 10, 64)
	if err != nil {
		l.Errorf("[logic SaveTask] %#v", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}
	_, has, err := l.svcCtx.TaskRepo.FindById(l.ctx, stageCode)
	if err != nil {
		l.Errorf("[logic SaveTask] %#v", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}
	if !has {
		err = respx.ErrTaskStageNotExists
		l.Errorf("[logic SaveTask] %#v", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}
	
	projectCode, err := l.svcCtx.Encrypter.DecryptInt64(in.ProjectCode)
	if err != nil {
		l.Errorf("[logic SaveTask] %#v", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}
	msg, err := l.svcCtx.ProjectClient.FindProjectById(l.ctx, &projectservice.FindProjectByIdRequest{
		ProjectCode: projectCode,
	})	
	if err != nil {
		l.Errorf("[logic SaveTask] %#v", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}
	if msg == nil || msg.Deleted == domain.Deleted {
		er := respx.ErrProjectAlreadyDeleted
		l.Errorf("[logic SaveTask] %#v", er)
		return nil, respx.ToStatusErr(er)
	}

	maxIdNum, err := l.svcCtx.TaskRepo.FindTaskMaxIdNum(l.ctx, projectCode)
	if err != nil {
		l.Errorf("[logic SaveTask] %#v", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}

	maxSort, err := l.svcCtx.TaskRepo.FindTaskSort(l.ctx, projectCode, stageCode)
	if err != nil {
		l.Errorf("[logic SaveTask] %#v", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}

	assignTo, err := l.svcCtx.Encrypter.DecryptInt64(in.AssignTo)
	if err != nil {
		l.Errorf("[logic SaveTask] %#v", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}
	
	task := domain.Task{
		Name:        in.GetName(),
		CreateTime:  time.Now().UnixMilli(),
		CreateBy:    in.GetMemberId(),
		AssignTo:    assignTo,
		ProjectCode: projectCode,
		StageCode:   int(stageCode),
		IdNum:       int(maxIdNum + 1),
		Private:     int(msg.OpenTaskPrivate),
		Sort:        int(maxSort + 1),
		BeginTime:   time.Now().UnixMilli(),
		EndTime:     time.Now().Add(2 * 24 * time.Hour).UnixMilli(),
	}	
	taskMember := domain.TaskMember{
		MemberCode: assignTo,
		IsExecutor: domain.NoExecutor,
		JoinTime: time.Now().UnixMilli(),
		IsOwner: domain.Owner,
	}
	if assignTo == in.GetMemberId() {
		taskMember.IsExecutor = domain.IsExecutor
	}
	taskId, taskMemberId, err := l.svcCtx.TaskRepo.CreateTaskAndMember(l.ctx, task, taskMember)
	if err != nil {
		l.Errorf("[logic SaveTask] %#v", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}
	task.Id = taskId
	taskMember.Id = taskMemberId

	display := task.ToTaskDisplay(l.svcCtx.Encrypter)
	memberInfo, err := l.svcCtx.UserClient.MemberInfo(l.ctx, &loginservice.MemberInfoRequest{
		Id: assignTo,
	})
	member := memberInfo.GetMember()
	display.Executor = domain.Executor{
		Name: member.GetName(),
		Avatar: member.GetAvatar(),
		Code: member.GetCode(),
	}
	tm := &v1.TaskMessage{}
	err = copier.Copy(tm, display)
	if err != nil {
		l.Errorf("[logic SaveTask] %#v", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}

	return tm, nil
}
