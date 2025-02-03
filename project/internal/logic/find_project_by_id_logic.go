package logic

import (
	"context"
	"time"

	"github.com/WeiXinao/msProject/api/proto/gen/project/v1"
	"github.com/WeiXinao/msProject/pkg/respx"
	"github.com/WeiXinao/msProject/project/internal/svc"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindProjectByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindProjectByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindProjectByIdLogic {
	return &FindProjectByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindProjectByIdLogic) FindProjectById(in *v1.FindProjectByIdRequest) (*v1.ProjectMessage, error) {
	project, err := l.svcCtx.ProjectRepo.FindProjectById(l.ctx, in.GetProjectCode())
	if err != nil {
		l.Error("[logic FindProjectById] %#v", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}
	projectMsg := &v1.ProjectMessage{}
	err = copier.Copy(&projectMsg, project)
	if err != nil {
		l.Error("[logic FindProjectById] %#v", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}
	if projectMsg == nil {
		return nil, nil
	}
	projectMsg.AccessControlType = project.GetAccessControlType()
	projectMsg.Order = int32(project.Order)
	projectMsg.CreateTime = time.UnixMilli(project.CreateTime).Format(time.DateTime)
	projectMsg.Code, _ = l.svcCtx.Encrypter.EncryptInt64(project.Id)
	projectMsg.TemplateCode, _ = l.svcCtx.Encrypter.EncryptInt64(int64(project.TemplateCode))
	return projectMsg, nil
}
