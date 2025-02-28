// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.5
// Source: file.proto

package server

import (
	"context"

	"github.com/WeiXinao/msProject/file/internal/logic"
	"github.com/WeiXinao/msProject/file/internal/svc"
	"github.com/WeiXinao/msProject/api/proto/gen/file/v1"
)

type FileServer struct {
	svcCtx *svc.ServiceContext
	v1.UnimplementedFileServer
}

func NewFileServer(svcCtx *svc.ServiceContext) *FileServer {
	return &FileServer{
		svcCtx: svcCtx,
	}
}

func (s *FileServer) SaveTaskFile(ctx context.Context, in *v1.TaskFileRequest) (*v1.TaskFileResponse, error) {
	l := logic.NewSaveTaskFileLogic(ctx, s.svcCtx)
	return l.SaveTaskFile(in)
}

func (s *FileServer) TaskSources(ctx context.Context, in *v1.TaskSourcesRequest) (*v1.TaskSourceResponse, error) {
	l := logic.NewTaskSourcesLogic(ctx, s.svcCtx)
	return l.TaskSources(in)
}
