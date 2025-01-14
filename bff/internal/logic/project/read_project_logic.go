package project

import (
	"context"
	projectv1 "github.com/WeiXinao/msProject/api/proto/gen/project/v1"
	userv1 "github.com/WeiXinao/msProject/api/proto/gen/user/v1"
	"github.com/WeiXinao/msProject/pkg/respx"
	"github.com/jinzhu/copier"

	"github.com/WeiXinao/msProject/bff/internal/svc"
	"github.com/WeiXinao/msProject/bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReadProjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReadProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReadProjectLogic {
	return &ReadProjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReadProjectLogic) ReadProject(req *types.ReadProjectReq) (resp *types.ReadProjectRsp, err error) {
	memberId := l.ctx.Value("memberId").(int64)
	projectDetailRsp, err := l.svcCtx.ProjectClient.ProjectDetail(l.ctx, &projectv1.ProjectDetailRequest{
		ProjectCode: req.ProjectCode,
		MemberId:    memberId,
	})
	if err != nil {
		l.Error("[logic ReadProject]", err)
		return nil, respx.FromStatusErr(err)
	}
	memberInfoRsp, err := l.svcCtx.UserClient.MemberInfo(l.ctx, &userv1.MemberInfoRequest{
		Id: projectDetailRsp.GetIsOwner(),
	})
	if err != nil {
		l.Error("[logic ReadProject]", err)
		return nil, respx.FromStatusErr(err)
	}
	resp = &types.ReadProjectRsp{}
	err = copier.Copy(resp, projectDetailRsp)
	if err != nil {
		l.Error("[logic ReadProject]", err)
		return nil, respx.FromStatusErr(err)
	}
	resp.OwnerName = memberInfoRsp.GetMember().GetName()
	resp.OwnerAvatar = memberInfoRsp.GetMember().GetAvatar()
	return
}
