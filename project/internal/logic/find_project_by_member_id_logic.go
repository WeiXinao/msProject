package logic

import (
	"context"

	"github.com/WeiXinao/msProject/api/proto/gen/project/v1"
	"github.com/WeiXinao/msProject/pkg/gozerox"
	"github.com/WeiXinao/msProject/project/internal/svc"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindProjectByMemberIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindProjectByMemberIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindProjectByMemberIdLogic {
	return &FindProjectByMemberIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindProjectByMemberIdLogic) FindProjectByMemberId(in *v1.FindProjectByMemberIdRequest) (*v1.FindProjectByMemberIdResponse, error) {
	return gozerox.RpcLogicWrapper(l.ctx, l, in, func(methodName string, logic *FindProjectByMemberIdLogic, req *v1.FindProjectByMemberIdRequest) (*v1.FindProjectByMemberIdResponse, error) {
		pam, err := logic.svcCtx.ProjectRepo.GetProjectAndMemberByPidAndMid(logic.ctx, in.GetProjectId(), in.GetMemberId())
		if err != nil {
			return nil, err
		}	
		if pam.MemberCode == 0 && pam.Project.Id == 0 {
			return &v1.FindProjectByMemberIdResponse{
				Project: nil,
				IsOwner: false,
				IsMember: false,
			}, nil
		}
		
		isOwner := false
		if  pam.IsOwner == pam.MemberCode {
			isOwner = true
		}

		projectMsg := &v1.ProjectMessage{}
		err = copier.Copy(projectMsg, pam.Project)
		if err != nil {
			return nil, err
		}

		return &v1.FindProjectByMemberIdResponse{
			Project: projectMsg,
			IsOwner: isOwner,
			IsMember: true,
		}, nil
	})
}
