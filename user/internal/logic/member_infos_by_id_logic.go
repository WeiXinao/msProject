package logic

import (
	"context"
	"time"

	userv1 "github.com/WeiXinao/msProject/api/proto/gen/user/v1"
	"github.com/WeiXinao/msProject/user/internal/domain"
	"github.com/WeiXinao/msProject/user/internal/svc"
	"github.com/WeiXinao/xkit/slice"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberInfosByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberInfosByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberInfosByIdLogic {
	return &MemberInfosByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberInfosByIdLogic) MemberInfosById(in *userv1.MemberInfosByIdRequest) (*userv1.MemberInfosByIdResponse, error) {
	memberDmns, err := l.svcCtx.UserRepo.GetMemberByIds(l.ctx, in.MIds)
	if err != nil {
		l.Error("[logic MemberInfosById] %w", err)	
		return nil, err
	}

	list := slice.Map(memberDmns, func(idx int, src *domain.Member) *userv1.MemberMessage {
		memberMsg := &userv1.MemberMessage{}
		err = copier.Copy(&memberMsg, src)
		if err != nil {
			l.Error("[logic MemberInfosById] %w", err)	
			return nil
		}
		memberMsg.LastLoginTime = src.LastLoginTime.Format(time.DateTime)
		memberMsg.Code, _ = l.svcCtx.Encrypter.EncryptInt64(src.Id)
		memberMsg.CreateTime = src.CreateTime.Format(time.DateTime)
		return memberMsg
	})
	return &userv1.MemberInfosByIdResponse{
		List: list,
	}, nil
}
