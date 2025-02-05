package logic

import (
	"context"

	"github.com/WeiXinao/msProject/api/proto/gen/task/v1"
	"github.com/WeiXinao/msProject/pkg/respx"
	"github.com/WeiXinao/msProject/project/projectservice"
	"github.com/WeiXinao/msProject/task/internal/domain"
	"github.com/WeiXinao/msProject/task/internal/svc"
	"github.com/WeiXinao/msProject/user/loginservice"
	"github.com/jinzhu/copier"
	"golang.org/x/sync/errgroup"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReadTaskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewReadTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReadTaskLogic {
	return &ReadTaskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ReadTaskLogic) ReadTask(in *v1.ReadTaskRequest) (*v1.TaskMessage, error) {
	taskCode, err := l.svcCtx.Encrypter.DecryptInt64(in.TaskCode)
	if err != nil {
		l.Errorf("[logic ReadTask] %#v", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}

	var (
		eg = errgroup.Group{}
		display *domain.TaskDisplay
		projectMsg *projectservice.ProjectMessage
		taskStages *domain.TaskStages
		executor domain.Executor
	)
	taskInfo, err := l.svcCtx.TaskRepo.FindTaskById(l.ctx, taskCode)
	if err != nil {
		l.Errorf("[logic ReadTask] %#v", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}

	eg.Go(func() error {
		display = taskInfo.ToTaskDisplay(l.svcCtx.Encrypter)
		if taskInfo.Private == 1 {
			// 代表隐私模式
			_, has, err := l.svcCtx.TaskRepo.FindTaskMemberByTaskId(l.ctx, taskInfo.Id, in.MemberId)
			if err != nil {
				return err
			}
			if has {
				display.CanRead = domain.CanRead
			} else {
				display.CanRead = domain.NoCanRead
			}
		}
		
		return nil
	})

	eg.Go(func() error {
		projectMsg, err = l.svcCtx.ProjectClient.FindProjectById(l.ctx, &projectservice.FindProjectByIdRequest{
			ProjectCode: taskInfo.ProjectCode,
		})
		if err != nil {
			return err
		}
		return nil
	})

	eg.Go(func() error {
		taskStages, _, err = l.svcCtx.TaskRepo.FindById(l.ctx, int64(taskInfo.StageCode))
		if err != nil {
			return err
		}
		return nil
	})

	eg.Go(func() error {
		memberMsg, err := l.svcCtx.UserClient.MemberInfo(l.ctx, &loginservice.MemberInfoRequest{
			Id: taskInfo.AssignTo,	
		})
		if err != nil {
			return err
		}
		executor = domain.Executor{
			Name: memberMsg.GetMember().GetName(),
			Avatar: memberMsg.GetMember().GetAvatar(),
		}
		return nil
	})

	err = eg.Wait()
	if err != nil {
		l.Errorf("[logic ReadTask] %#v", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}

	display.ProjectName = projectMsg.GetName()
	display.StageName = taskStages.Name
	display.Executor = executor

	taskMsg := &v1.TaskMessage{}
	err = copier.Copy(taskMsg, display)
	if err != nil {
		l.Errorf("[logic ReadTask] %#v", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}

	return taskMsg, nil
}
