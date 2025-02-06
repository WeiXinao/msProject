package logic

import (
	"context"

	"github.com/WeiXinao/msProject/api/proto/gen/project/v1"
	"github.com/WeiXinao/msProject/pkg/respx"
	"github.com/WeiXinao/msProject/project/internal/domain"
	"github.com/WeiXinao/msProject/project/internal/svc"
	"github.com/WeiXinao/msProject/user/loginservice"
	"github.com/WeiXinao/xkit/slice"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTaskLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskLogLogic {
	return &TaskLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TaskLogLogic) TaskLog(in *v1.TaskLogRequest) (*v1.TaskLogResponse, error) {
	taskCode, err := l.svcCtx.Encrypter.DecryptInt64(in.GetTaskCode())
	if err != nil {
		l.Errorf("[logic TaskLog] %#v", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}
	var (
		projectLogs []*domain.ProjectLog
		total int64
	)
	switch in.All {
	case 1:
		projectLogs, total, err = l.svcCtx.ProjectRepo.FindLogByTaskCode(l.ctx, taskCode, int(in.GetComment()))
	case 0:
		projectLogs, total, err = l.svcCtx.ProjectRepo.FindLogByTaskCodePagination(l.ctx, taskCode,
			 int(in.GetComment()), in.Page, in.PageSize)
	}
	if err != nil {
		l.Errorf("[logic TaskLog] %#v", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}
	if total == 0 {
		return &v1.TaskLogResponse{}, nil
	}

	mids := slice.Map(projectLogs, func(idx int, src *domain.ProjectLog) int64 {
		return src.MemberCode
	})
	memberInfosRsp, err := l.svcCtx.UserClient.MemberInfosById(l.ctx, &loginservice.MemberInfosByIdRequest{
		MIds: mids,
	})
	idToMemberMsg := slice.ToMap(memberInfosRsp.GetList(), func(element *loginservice.MemberMessage) int64 {
		return element.GetId()
	}) 
	if err != nil {
		l.Errorf("[logic TaskLog] %#v", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}

	projectLogDisplays := slice.Map(projectLogs, func(idx int, src *domain.ProjectLog) *domain.ProjectLogDisplay {
		display := src.ToDisplay(l.svcCtx.Encrypter)
		memberInfo := idToMemberMsg[src.MemberCode]
		member := domain.Member{
			Id: memberInfo.GetId(),
			Name: memberInfo.GetName(),
			Avatar: memberInfo.GetAvatar(),
			Code: memberInfo.GetCode(),
		}
		display.Member = member
		return display
	})
	
	list := make([]*v1.TaskLog, 0)
	err = copier.CopyWithOption(&list, projectLogDisplays, copier.Option{DeepCopy: true})
	if err != nil {
		l.Errorf("[logic TaskLog] %#v", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}

	return &v1.TaskLogResponse{
		List: list,
		Total: total,
	}, nil
}
