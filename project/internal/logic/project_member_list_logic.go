package logic

import (
	"context"
	"strconv"

	"github.com/WeiXinao/msProject/api/proto/gen/project/v1"
	userv1 "github.com/WeiXinao/msProject/api/proto/gen/user/v1"
	"github.com/WeiXinao/msProject/pkg/respx"
	"github.com/WeiXinao/msProject/project/internal/domain"
	"github.com/WeiXinao/msProject/project/internal/svc"
	"github.com/WeiXinao/xkit/slice"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProjectMemberListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProjectMemberListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProjectMemberListLogic {
	return &ProjectMemberListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProjectMemberListLogic) ProjectMemberList(in *v1.ProjectMemberListRequest) (*v1.ProjectMemberListResponse, error) {
	// 1. 去 project_member 表去查询用户 id 列表
	projectCodeStr, err := l.svcCtx.Encrypter.Decrypt(in.ProjectCode)
	if err != nil {
		l.Error("[logic ProjectMemberList] %#v", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}
	projectCode, err := strconv.ParseInt(projectCodeStr, 10, 64)
	if err != nil {
		l.Error("[logic ProjectMemberList] %#v", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}
	pm, err := l.svcCtx.ProjectRepo.GetProjectMembersByPid(l.ctx, projectCode)
	if err != nil {
		l.Error("[logic ProjectMemberList] %#v", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}
	if pm == nil || len(pm) <= 0 {
		return &v1.ProjectMemberListResponse{}, nil
	}

	// 2.	拿上用户 id 列表去请求用户信息
	mIds := slice.Map[*domain.ProjectMember, int64](pm, 
		func(idx int, src *domain.ProjectMember) int64 {
			return src.MemberCode
	})

	memberInfosResp, err := l.svcCtx.UserClient.MemberInfosById(l.ctx, &userv1.MemberInfosByIdRequest{
		MIds: mIds,
	})
	if err != nil {
		l.Error("[logic ProjectMemberList] %#v", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}
	mIdToPmMap := slice.ToMap(pm, func(element *domain.ProjectMember) int64 {
		return element.MemberCode
	})

	return &v1.ProjectMemberListResponse{
		List: slice.Map(memberInfosResp.GetList(), func(idx int, src *userv1.MemberMessage) *v1.ProjectMemberMessage {
			owner := mIdToPmMap[src.Id].IsOwner
			isOwner := domain.NotOwner
			if owner == src.Id {
				isOwner = domain.Owner
			}
			return &v1.ProjectMemberMessage{
				MemberCode: src.GetId(),
				Name: src.GetName(),
				Avatar: src.GetAvatar(),
				Email: src.GetEmail(),
				Code: src.GetCode(),
				IsOwner: int32(isOwner),
			}
		}),
		Total: int64(len(mIds)),
	}, nil
}