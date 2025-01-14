package logic

import (
	"context"
	"github.com/WeiXinao/msProject/pkg/respx"
	"github.com/jinzhu/copier"
	"time"

	userv1 "github.com/WeiXinao/msProject/api/proto/gen/user/v1"
	"github.com/WeiXinao/msProject/user/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberInfoLogic {
	return &MemberInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberInfoLogic) MemberInfo(in *userv1.MemberInfoRequest) (*userv1.MemberInfoResponse, error) {
	m, err := l.svcCtx.UserRepo.GetMemberById(l.ctx, in.GetId())
	if err != nil {
		l.Error("[logic MemberInfo]", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}

	memberMsg := &userv1.MemberMessage{}
	err = copier.Copy(memberMsg, &m)
	if err != nil {
		return nil, err
	}
	memberMsg.LastLoginTime = m.LastLoginTime.Format(time.DateTime)
	memberMsg.Code, _ = l.svcCtx.Encrypter.EncryptInt64(m.Id)
	memberMsg.CreateTime = m.CreateTime.Format(time.DateTime)

	orgs, err := l.svcCtx.UserRepo.GetOrganizationByMemId(l.ctx, m.Id)
	if err != nil {
		l.Error("[logic MemberInfo]", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}
	if len(orgs) > 0 {
		memberMsg.OrganizationCode, err = l.svcCtx.Encrypter.EncryptInt64(orgs[0].Id)
		if err != nil {
			l.Error("[logic MemberInfo]", err)
			return nil, respx.ToStatusErr(respx.ErrInternalServer)
		}
	}

	return &userv1.MemberInfoResponse{
		Member: memberMsg,
	}, nil
}
