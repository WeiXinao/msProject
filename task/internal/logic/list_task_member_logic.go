package logic

import (
	"context"

	"github.com/WeiXinao/msProject/api/proto/gen/task/v1"
	"github.com/WeiXinao/msProject/pkg/respx"
	"github.com/WeiXinao/msProject/task/internal/domain"
	"github.com/WeiXinao/msProject/task/internal/svc"
	"github.com/WeiXinao/msProject/user/loginservice"
	"github.com/WeiXinao/xkit/slice"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListTaskMemberLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListTaskMemberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListTaskMemberLogic {
	return &ListTaskMemberLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListTaskMemberLogic) ListTaskMember(in *v1.ListTaskMemberRequest) (*v1.ListTaskMemberResponse, error) {
	taskCode, err := l.svcCtx.Encrypter.DecryptInt64(in.TaskCode)
	if err != nil {
		l.Error("[logic ListTaskMember] %#v", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}
	taskMembers, total, err := l.svcCtx.TaskRepo.FindTaskMemberByTaskIdPagination(l.ctx, taskCode, 
		in.Page, in.PageSize)
	if err != nil {
		l.Error("[logic ListTaskMember] %#v", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}
	mids := slice.Map(taskMembers, func(idx int, src *domain.TaskMember) int64 {
		return src.MemberCode
	})
	memberInfosRsp, err := l.svcCtx.UserClient.MemberInfosById(l.ctx, &loginservice.MemberInfosByIdRequest{
		MIds: mids,
	})
	if err != nil {
		l.Errorf("[logic ListTaskMember] %#v", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}
	idMapMemberMsg := slice.ToMap(memberInfosRsp.List, 
	func(element *loginservice.MemberMessage) int64 {
		return element.GetId()
	})
	list := slice.Map(taskMembers, func(idx int, src *domain.TaskMember) *v1.TaskMemberMessage {
		code, err := l.svcCtx.Encrypter.EncryptInt64(src.MemberCode)
		if err != nil {
			l.Errorf("[logic ListTaskMember] %#v", err)
		}
		memberInfo := idMapMemberMsg[src.MemberCode]
		return &v1.TaskMemberMessage{
			Code: code,
			Id: src.Id,
			Name: memberInfo.GetName(),
			Avatar: memberInfo.GetAvatar(),
			IsExecutor: int32(src.IsExecutor),
			IsOwner: int32(src.IsOwner),
		}
	})
	return &v1.ListTaskMemberResponse{
		List: list,
		Total: total,
	}, nil
}
