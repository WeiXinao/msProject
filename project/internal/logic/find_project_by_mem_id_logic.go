package logic

import (
	"context"
	"github.com/WeiXinao/msProject/api/proto/gen/project/v1"
	"github.com/WeiXinao/msProject/pkg/respx"
	"github.com/WeiXinao/msProject/project/internal/domain"
	"github.com/WeiXinao/msProject/project/internal/svc"
	"github.com/jinzhu/copier"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindProjectByMemIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindProjectByMemIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindProjectByMemIdLogic {
	return &FindProjectByMemIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindProjectByMemIdLogic) FindProjectByMemId(in *v1.ProjectRequest) (*v1.ProjectResponse, error) {
	memberId := in.MemberId
	page := in.Page
	pageSize := in.PageSize
	pms, total, err := l.svcCtx.ProjectRepo.FindProjectByMemId(l.ctx, in.SelectBy, memberId, page, pageSize)
	if err != nil {
		l.Error("[FindProjectByMemId]", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}
	if pms == nil {
		return &v1.ProjectResponse{Pm: []*v1.ProjectMessage{}, Total: total}, nil
	}
	pm := []*v1.ProjectMessage{}
	err = copier.CopyWithOption(&pm, &pms, copier.Option{DeepCopy: true})
	if err != nil {
		l.Error("[FindProjectByMemId]", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}
	for _, p := range pm {
		p.Code, err = l.svcCtx.Encrypter.EncryptInt64(p.Id)
		if err != nil {
			l.Error("[FindProjectByMemId]", err)
		}
		pam := domain.ToMap(pms)[p.Id]
		p.AccessControlType = pam.GetAccessControlType()
		p.JoinTime = time.UnixMilli(pam.JoinTime).Format(time.DateTime)
		p.OwnerName = in.MemberName
		p.Order = int32(pam.Order)
		p.CreateTime = time.UnixMilli(pam.CreateTime).Format(time.DateTime)
	}
	return &v1.ProjectResponse{Pm: pm, Total: total}, nil
}
