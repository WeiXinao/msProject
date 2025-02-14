package logic

import (
	"context"

	"github.com/WeiXinao/msProject/account/internal/svc"
	"github.com/WeiXinao/msProject/api/proto/gen/account/v1"
	"github.com/WeiXinao/msProject/pkg/gozerox"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReadDepartmentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewReadDepartmentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReadDepartmentLogic {
	return &ReadDepartmentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ReadDepartmentLogic) ReadDepartment(in *v1.ReadDepartmentRequest) (*v1.DepartmentMessage, error) {
	return gozerox.RpcLogicWrapper(l.ctx, l, in, func(methodName string, logic *ReadDepartmentLogic, req *v1.ReadDepartmentRequest) (*v1.DepartmentMessage, error) {
		// organCode, err := logic.svcCtx.Encrypter.DecryptInt64(req.OrganizationCode)
		// if err != nil {
		// 	return nil, err
		// }
		deptCode, err := logic.svcCtx.Encrypter.DecryptInt64(req.DepartmentCode)
		if err != nil {
			return nil, err
		}

		dept, err := l.svcCtx.AccoutRepo.FindDepartmentById(l.ctx, deptCode)
		if err != nil {
			return nil, err
		}
		res := &v1.DepartmentMessage{}
		err = copier.Copy(res, dept.ToDisplay(l.svcCtx.Encrypter))
		return res, err
	})	
}
