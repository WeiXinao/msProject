package logic

import (
	"context"
	"strings"
	"time"

	"github.com/WeiXinao/msProject/account/internal/domain"
	"github.com/WeiXinao/msProject/account/internal/svc"
	"github.com/WeiXinao/msProject/api/proto/gen/account/v1"
	"github.com/WeiXinao/msProject/pkg/gozerox"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveDepartmentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveDepartmentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveDepartmentLogic {
	return &SaveDepartmentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SaveDepartmentLogic) SaveDepartment(in *v1.SaveDepartmentRequest) (*v1.DepartmentMessage, error) {
	return gozerox.RpcLogicWrapper(l.ctx, l, in, func(methodName string, logic *SaveDepartmentLogic, req *v1.SaveDepartmentRequest) (*v1.DepartmentMessage, error) {
		organzationCode, err := l.svcCtx.Encrypter.DecryptInt64(in.OrganizationCode)
		if err != nil {
			return nil, err
		}
		departmentCode := int64(0)
		if strings.TrimSpace(in.DepartmentCode) != "" {
			departmentCode, err = l.svcCtx.Encrypter.DecryptInt64(in.DepartmentCode)
			if err != nil {
				return nil, err	
			}
		}
		parentDepartmentCode := int64(0)
		if strings.TrimSpace(in.ParentDepartmentCode) != "" {
			parentDepartmentCode, err = l.svcCtx.Encrypter.DecryptInt64(in.ParentDepartmentCode)
			if err != nil {
				return nil, err
			}
		}
		dept, err := l.svcCtx.AccoutRepo.SaveDepartment(l.ctx, domain.Department{
			Id: departmentCode,
			Name: in.Name,
			OrganizationCode: organzationCode,
			Pcode: parentDepartmentCode,
			CreateTime: time.Now().UnixMilli(),
		})
		if err != nil {
			return nil, err
		}
		display := dept.ToDisplay(l.svcCtx.Encrypter)
		deptMsg := &v1.DepartmentMessage{}
		err = copier.Copy(deptMsg, display)
		if err != nil {
			return nil, err
		}
		return deptMsg, nil
	})
}
