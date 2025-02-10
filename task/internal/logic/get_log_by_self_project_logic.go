package logic

import (
	"context"
	"time"

	"github.com/WeiXinao/msProject/api/proto/gen/task/v1"
	"github.com/WeiXinao/msProject/project/projectservice"
	"github.com/WeiXinao/msProject/task/internal/domain"
	"github.com/WeiXinao/msProject/task/internal/svc"
	"github.com/WeiXinao/msProject/user/loginservice"
	"github.com/WeiXinao/xkit/slice"
	"github.com/jinzhu/copier"
	"golang.org/x/sync/errgroup"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLogBySelfProjectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLogBySelfProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLogBySelfProjectLogic {
	return &GetLogBySelfProjectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLogBySelfProjectLogic) GetLogBySelfProject(in *v1.GetLogBySelfProjectRequest) (*v1.GetLogBySelfProjectResponse, error) {
	// 根据用户 id 查询当前用户的日志表
	pls, total, err := l.svcCtx.TaskRepo.
	FindLogByMemberCode(l.ctx, in.GetMemberId(), in.GetPage(), in.GetPageSize())
	if err != nil {
		l.Errorf("[logic GetLogBySelfProject] %#v", err)
		return nil, err
	}

	var (
		projectsChan = make(chan map[int64]*projectservice.ProjectMessage, 1)
		membersChan = make(chan map[int64]*loginservice.MemberMessage, 1)
		tasksChan = make(chan map[int64]*domain.Task, 1)
	)
	defer func ()  {
		close(projectsChan)	
		close(membersChan)	
		close(tasksChan)
	}()
	eg := errgroup.Group{}
	eg.Go(func() error {
		pids := slice.Map(pls, func(idx int, src *domain.ProjectLog) int64 {
			return src.ProjectCode
		})
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)	
		defer cancel()
		projectsRep, err := l.svcCtx.ProjectClient.FindProjectByIds(ctx, &projectservice.FindProjectByIdsRequest{
			ProjectCodes: pids,	
		})
		if err != nil {
			return err
		}
		projectsChan <- slice.ToMap(projectsRep.GetProjects(), func(element *projectservice.ProjectMessage) int64 {
			return element.GetId()
		})
		return nil
	})
	
	eg.Go(func() error {
		mids := slice.Map(pls, func(idx int, src *domain.ProjectLog) int64 {
			return src.MemberCode
		})
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		membersRsp, err := l.svcCtx.UserClient.MemberInfosById(ctx, &loginservice.MemberInfosByIdRequest{
			MIds: mids,
		})
		if err != nil {
			return err
		}
		membersChan <- slice.ToMap(membersRsp.GetList(), 
		func(element *loginservice.MemberMessage) int64 {
			return element.GetId()
		})
		return nil
	})

	eg.Go(func() error {
		tids := slice.Map(pls, func(idx int, src *domain.ProjectLog) int64 {
			return src.SourceCode
		})
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)	
		defer cancel()
		tasks, err := l.svcCtx.TaskRepo.FindTaskByIds(ctx, tids)
		if err != nil {
			return err
		}
		tasksChan <- slice.ToMap(tasks, func(element *domain.Task) int64 {
			return element.Id
		})
		return nil
	})

	err = eg.Wait()
	if err != nil {
		l.Errorf("[logic GetLogBySelfProject] %#v", err)
		return nil, err
	}

	idToProject := <-projectsChan
	idToMember := <- membersChan
	idToTask := <- tasksChan
	display := slice.Map(pls, func(idx int, src *domain.ProjectLog) *domain.IndexProjectLogDisplay {
		display := src.ToIndexDisplay(l.svcCtx.Encrypter)
		display.ProjectName = idToProject[src.ProjectCode].GetName()
		memberMsg := idToMember[src.MemberCode]
		display.MemberName = memberMsg.GetName()
		display.MemberAvatar = memberMsg.GetAvatar()
		display.TaskName = idToTask[src.SourceCode].Name
		return display
	})
	list := make([]*v1.ProjectLogMessage, 0) 
	err = copier.CopyWithOption(&list, display, copier.Option{DeepCopy: true})
	if err != nil {
		l.Errorf("[logic GetLogBySelfProject] %#v", err)
		return nil, err
	}
	
	return &v1.GetLogBySelfProjectResponse{
		List: list,
		Total: total,
	}, nil
}
