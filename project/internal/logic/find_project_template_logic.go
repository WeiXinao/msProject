package logic

import (
	"context"
	"github.com/WeiXinao/msProject/pkg/respx"
	"github.com/WeiXinao/msProject/project/internal/domain"
	"github.com/WeiXinao/xkit/slice"
	"github.com/jinzhu/copier"
	"strconv"

	"github.com/WeiXinao/msProject/api/proto/gen/project/v1"
	"github.com/WeiXinao/msProject/project/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindProjectTemplateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindProjectTemplateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindProjectTemplateLogic {
	return &FindProjectTemplateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindProjectTemplateLogic) FindProjectTemplate(in *v1.FindProjectTemplateRequest) (*v1.FindProjectTemplateResponse, error) {
	// 1. 根据 viewType 去查询项目模板表，得到 list
	var (
		pts   []domain.ProjectTemplate
		total int64
		err   error
	)

	organizationCodeStr, err := l.svcCtx.Encrypter.Decrypt(in.OrganizationCode)
	if err != nil {
		l.Error("[logic FindProjectTemplate]", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}
	organizationCode, err := strconv.ParseInt(organizationCodeStr, 10, 64)
	if err != nil {
		l.Error("[logic FindProjectTemplate]", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}
	page := in.Page
	pageSize := in.PageSize

	switch in.ViewType {
	case -1:
		// 查询所有模板
		pts, total, err = l.svcCtx.ProjectRepo.FindProjectTemplateAll(l.ctx, organizationCode, page, pageSize)
	case 0:
		// 查询自定义模板
		pts, total, err = l.svcCtx.ProjectRepo.FindProjectTemplateCustom(l.ctx, in.MemberId, organizationCode, page, pageSize)
	case 1:
		// 查询系统模板
		pts, total, err = l.svcCtx.ProjectRepo.FindProjectTemplateSystem(l.ctx, page, pageSize)
	}
	if err != nil {
		l.Error("[logic FindProjectTemplate]", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}
	// 2. 模型转换，拿到模板 id 列表，去 任务步骤模板表 进行查询
	taskStageTmpls, err := l.svcCtx.ProjectRepo.FindInProTemIds(l.ctx, domain.ToProjectTemplateIds(pts))
	if err != nil {
		l.Error("[logic FindProjectTemplate]", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}
	// 3. 组装数据
	ptas := slice.Map(pts, func(idx int, src domain.ProjectTemplate) *domain.ProjectTemplateAll {
		// 写代码，该谁做的事情一定要交出去
		return src.Convert(l.svcCtx.Encrypter, domain.CovertProjectMap(taskStageTmpls)[src.Id])
	})
	var pmMsgs []*v1.ProjectTemplateMessage
	err = copier.CopyWithOption(&pmMsgs, &ptas, copier.Option{DeepCopy: true})
	if err != nil {
		l.Error("[logic FindProjectTemplate]", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}
	return &v1.FindProjectTemplateResponse{
		Ptm:   pmMsgs,
		Total: total,
	}, nil
}
