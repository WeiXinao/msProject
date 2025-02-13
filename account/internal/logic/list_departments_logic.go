package logic

import (
	"context"
	"strings"

	"github.com/WeiXinao/msProject/account/internal/domain"
	"github.com/WeiXinao/msProject/account/internal/svc"
	"github.com/WeiXinao/msProject/api/proto/gen/account/v1"
	"github.com/WeiXinao/msProject/pkg/respx"
	"github.com/WeiXinao/xkit/slice"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListDepartmentsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListDepartmentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListDepartmentsLogic {
	return &ListDepartmentsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListDepartmentsLogic) ListDepartments(in *v1.ListDepartmentsReqeust) (*v1.ListDepartmentsResponse, error) {
	organizationCode, err := l.svcCtx.Encrypter.DecryptInt64(in.OrganizationCode)
	if err != nil {
		l.Errorf("[logic ListDepartments] %#v", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}
	parentDepartmentCode := int64(0)
	if strings.TrimSpace(in.ParentDepartmentCode) != "" {
		parentDepartmentCode, err = l.svcCtx.Encrypter.DecryptInt64(in.ParentDepartmentCode)	
		if err != nil {
			l.Errorf("[logic ListDepartments] %#v", err)
			return nil, respx.ToStatusErr(respx.ErrInternalServer)
		}
	}
	
	depts, total, err := l.svcCtx.AccoutRepo.ListDepartments(l.ctx, organizationCode, parentDepartmentCode, in.Page, in.PageSize)
	if err != nil {
		l.Errorf("[logic ListDepartments] %#v", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}
	deptDisplays := slice.Map(depts, func(idx int, src *domain.Department) *domain.DepartmentDisplay {
		return src.ToDisplay(l.svcCtx.Encrypter)
	})
	list := make([]*v1.DepartmentMessage, 0)
	err = copier.CopyWithOption(&list, deptDisplays, copier.Option{DeepCopy: true})
	if err != nil {
		l.Errorf("[logic ListDepartments] %#v", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}
	return &v1.ListDepartmentsResponse{
		List: list,
		Total: total,
	}, nil
}
