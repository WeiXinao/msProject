package logic

import (
	"context"
	"github.com/WeiXinao/msProject/pkg/respx"
	"github.com/WeiXinao/msProject/user/internal/domain"
	"github.com/WeiXinao/xkit/slice"
	"github.com/jinzhu/copier"
	"time"

	userv1 "github.com/WeiXinao/msProject/api/proto/gen/user/v1"
	"github.com/WeiXinao/msProject/user/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type MyOrgListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMyOrgListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MyOrgListLogic {
	return &MyOrgListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MyOrgListLogic) MyOrgList(in *userv1.MyOrgListRequest) (*userv1.MyOrgListResponse, error) {
	memId := in.MemId
	organizations, err := l.svcCtx.UserRepo.GetOrganizationByMemId(l.ctx, memId)
	if err != nil {
		l.Error("[logic MyOrgList]", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}
	orgs := slice.Map(organizations, func(idx int, src domain.Organization) *userv1.OrganizationMessage {
		org := &userv1.OrganizationMessage{}
		err = copier.Copy(&org, &src)
		if err != nil {
			l.Error("[logic MyOrgList]", err)
		}
		org.CreateTime = src.CTime.Format(time.DateTime)
		return org
	})

	for _, org := range orgs {
		org.Code, err = l.svcCtx.Encrypter.EncryptInt64(org.Id)
		if err != nil {
			l.Error("[logic MyOrgList]", err)
		}
	}
	return &userv1.MyOrgListResponse{
		OrganizationList: orgs,
	}, nil
}
