package department

import (
	"context"

	"github.com/WeiXinao/msProject/bff/internal/svc"
	"github.com/WeiXinao/msProject/bff/internal/types"

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
	// todo: add your logic here and delete this line

	return
}
