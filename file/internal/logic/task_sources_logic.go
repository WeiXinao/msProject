package logic

import (
	"context"

	"github.com/WeiXinao/msProject/api/proto/gen/file/v1"
	"github.com/WeiXinao/msProject/file/internal/domain"
	"github.com/WeiXinao/msProject/file/internal/svc"
	"github.com/WeiXinao/msProject/pkg/respx"
	"github.com/WeiXinao/xkit/slice"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskSourcesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTaskSourcesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskSourcesLogic {
	return &TaskSourcesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TaskSourcesLogic) TaskSources(in *v1.TaskSourcesRequest) (*v1.TaskSourceResponse, error) {
	taskCode, err := l.svcCtx.Encrypter.DecryptInt64(in.GetTaskCode())
	if err != nil {
		l.Error("[logic TaskSources] %#v", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}
	links, err := l.svcCtx.FileRepo.FindByTaskCode(l.ctx, taskCode)
	if err != nil {
		l.Error("[logic TaskSources] %#v", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}
	if len(links) == 0 {
		return &v1.TaskSourceResponse{}, nil
	}
	fids := slice.Map(links, func(idx int, src *domain.SourceLink) int64 {
		return src.SourceCode
	})

	files, err := l.svcCtx.FileRepo.FindByIds(l.ctx, fids)
	if err != nil {
		l.Error("[logic TaskSources] %#v", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}
	idToFiles := slice.ToMap(files, func(element *domain.File) int64 {
		return element.Id
	})

	linkDisplay := slice.Map(links, func(idx int, src *domain.SourceLink) *domain.SourceLinkDisplay {
		return src.ToDisplay(idToFiles[src.SourceCode], l.svcCtx.Encrypter)
	})
	list := make([]*v1.TaskSourceMessage, 0)
	err = copier.CopyWithOption(&list, linkDisplay, copier.Option{DeepCopy: true})
	if err != nil {
		l.Error("[logic TaskSources] %#v", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}
	return &v1.TaskSourceResponse{
		List: list,
	}, nil
}
