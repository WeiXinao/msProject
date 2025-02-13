package logic

import (
	"context"

	"github.com/WeiXinao/msProject/account/internal/svc"
	"github.com/WeiXinao/msProject/api/proto/gen/account/v1"

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
	// todo: add your logic here and delete this line

	return &v1.DepartmentMessage{}, nil
}
