package department

import (
	"context"

	v1 "github.com/WeiXinao/msProject/api/proto/gen/account/v1"
	"github.com/WeiXinao/msProject/bff/internal/middleware"
	"github.com/WeiXinao/msProject/bff/internal/svc"
	"github.com/WeiXinao/msProject/bff/internal/types"
	"github.com/WeiXinao/msProject/pkg/gozerox"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReadDepartmentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReadDepartmentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReadDepartmentLogic {
	return &ReadDepartmentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReadDepartmentLogic) ReadDepartment(req *types.ReadDepartment) (resp *types.Department, err error) {
	return gozerox.HttpLogicWrapper(l.ctx, l, req, func(methodName string, logic *ReadDepartmentLogic, req *types.ReadDepartment) (*types.Department, error) {
		deptMsg, err := l.svcCtx.AccountClient.ReadDepartment(l.ctx, &v1.ReadDepartmentRequest{
			DepartmentCode: req.DepartmentCode,
			OrganizationCode: l.ctx.Value(middleware.KeyOrganizationCode).(string),
		})
		if err != nil {
			return nil, err
		}
		dept := &types.Department{}
		err = copier.Copy(dept, deptMsg)
		if err != nil {
			return nil, err
		}
		return dept, nil
	})
}
