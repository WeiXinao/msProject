package logic

import (
	"context"
	"time"

	"github.com/WeiXinao/msProject/api/proto/gen/file/v1"
	"github.com/WeiXinao/msProject/file/internal/domain"
	"github.com/WeiXinao/msProject/file/internal/svc"
	"github.com/WeiXinao/msProject/pkg/respx"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveTaskFileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveTaskFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveTaskFileLogic {
	return &SaveTaskFileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SaveTaskFileLogic) SaveTaskFile(in *v1.TaskFileRequest) (*v1.TaskFileResponse, error) {
	organizationCode, err := l.svcCtx.Encrypter.DecryptInt64(in.GetOrganizationCode())
	if err != nil {
		l.Errorf("[logic SaveTaskFile]", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}
	taskCode, err := l.svcCtx.Encrypter.DecryptInt64(in.GetTaskCode())
	if err != nil {
		l.Errorf("[logic SaveTaskFile]", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}
	projectCode, err := l.svcCtx.Encrypter.DecryptInt64(in.GetProjectCode())
	if err != nil {
		l.Errorf("[logic SaveTaskFile]", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}
	err = l.svcCtx.FileRepo.SaveFileAndSourceLink(l.ctx, domain.File{
		PathName: in.GetPathName(),
		Title: in.GetFileName(),
		Extension: in.GetExtension(),
		Size: int(in.GetSize()),
		ObjectType: "",
		OrganizationCode: organizationCode,
		TaskCode: taskCode,
		ProjectCode: projectCode,
		CreateBy: in.GetMemberId(),
		CreateTime: time.Now().UnixMilli(),
		Downloads: 0,
		Extra: "",
		Deleted: domain.Undeleted,
		FileType: in.GetFileType(),
		FileUrl: in.GetFileUrl(),
		DeletedTime: 0,
	}, domain.SourceLink{
		SourceType: "file",
		LinkType: "task",
		LinkCode: taskCode,
		OrganizationCode: organizationCode,
		CreateBy: in.GetMemberId(),
		CreateTime: time.Now().UnixMilli(),
		Sort: 0,
	})
	if err != nil {
		l.Errorf("[logic SaveTaskFile]", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}
	return &v1.TaskFileResponse{}, nil
}
