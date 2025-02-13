package logic

import (
	"context"

	"github.com/WeiXinao/msProject/account/internal/svc"
	"github.com/WeiXinao/msProject/api/proto/gen/account/v1"

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
	// todo: add your logic here and delete this line

	return &v1.ListDepartmentsResponse{}, nil
}
