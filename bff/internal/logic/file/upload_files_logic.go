package file

import (
	"context"
	"net/http"
	"os"
	"path/filepath"
	"time"

	v1 "github.com/WeiXinao/msProject/api/proto/gen/file/v1"
	"github.com/WeiXinao/msProject/bff/internal/svc"
	"github.com/WeiXinao/msProject/bff/internal/types"
	"github.com/WeiXinao/msProject/pkg/formatx"
	"github.com/WeiXinao/msProject/pkg/fsx"
	"github.com/WeiXinao/msProject/pkg/respx"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadFilesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r *http.Request
}

func NewUploadFilesLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *UploadFilesLogic {
	return &UploadFilesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r: r,
	}
}

func (l *UploadFilesLogic) UploadFiles(req *types.UploadFileReq) (resp *types.UploadFileRsp, err error) {
	baseUrl := "http://localhost/"
	file, fileHeader, err := l.r.FormFile("file"); 
	if err != nil {
		l.Errorf("[logic UploadFiles] %#v", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}
	defer file.Close()
	
	path := filepath.Join(l.svcCtx.StaticPath, req.ProjectCode, req.TaskCode, formatx.ToDateString(time.Now()))
	if !fsx.IsExist(path) {
		os.MkdirAll(path, os.ModePerm)
	}

	fileName := filepath.Join(path, req.Identifier)
	openedFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		l.Errorf("[logic UploadFiles] %#v", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}
	buf := make([]byte, fileHeader.Size)
	file.Read(buf)
	openedFile.Write(buf)
	openedFile.Close()
	dst := filepath.Join(path, req.Filename)
	if req.TotalChunks == req.ChunkNumber {
		err = os.Rename(fileName, dst)
		if err != nil {
			l.Errorf("[logic UploadFiles] %#v", err)
			return nil, respx.ErrInternalServer
		}
		fileUrl := baseUrl + dst
		_, err := l.svcCtx.FileClient.SaveTaskFile(l.ctx, &v1.TaskFileRequest{
			TaskCode: req.TaskCode,
			ProjectCode: req.ProjectCode,
			OrganizationCode: l.ctx.Value("organizationCode").(string),
			PathName: dst,
			FileName: req.Filename,
			Size: int64(req.TotalSize),
			Extension: filepath.Ext(dst),
			FileUrl: fileUrl,
			FileType: fileHeader.Header.Get("Content-Type"),
			MemberId: l.ctx.Value("memberId").(int64),
		})
		if err != nil {
			err = respx.FromStatusErr(err)
			l.Errorf("[logic UploadFiles] %#v", err)
			return nil, err
		}
	}

	resp = &types.UploadFileRsp{
		File: dst,
		Key: dst,
		Url: baseUrl + dst,
		ProjectName: req.ProjectName,
	}

	return
}
