package logic

import (
	"context"
	"github.com/WeiXinao/msProject/pkg/respx"
	"strconv"

	"github.com/WeiXinao/msProject/api/proto/gen/project/v1"
	"github.com/WeiXinao/msProject/project/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RecycleOrRecoverProjectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRecycleOrRecoverProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecycleOrRecoverProjectLogic {
	return &RecycleOrRecoverProjectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RecycleOrRecoverProjectLogic) RecycleOrRecoverProject(in *v1.RecycleProjectRequest) (*v1.RecycleProjectResponse, error) {
	projectCodeStr, err := l.svcCtx.Encrypter.Decrypt(in.ProjectCode)
	if err != nil {
		l.Error("[logic RecycleOrRecoverProject]", err)
		return nil, respx.ToStatusErr(respx.FromStatusErr(err))
	}
	projectCode, err := strconv.ParseInt(projectCodeStr, 10, 64)
	if err != nil {
		l.Error("[logic RecycleOrRecoverProject]", err)
		return nil, respx.ToStatusErr(respx.FromStatusErr(err))
	}
	err = l.svcCtx.ProjectRepo.DeleteProject(l.ctx, projectCode, in.Deleted)
	return &v1.RecycleProjectResponse{}, nil
}
