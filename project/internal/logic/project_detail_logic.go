package logic

import (
	"context"
	"github.com/WeiXinao/msProject/pkg/respx"
	"github.com/WeiXinao/msProject/project/internal/domain"
	"github.com/jinzhu/copier"
	"strconv"
	"time"

	"github.com/WeiXinao/msProject/api/proto/gen/project/v1"
	"github.com/WeiXinao/msProject/project/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProjectDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProjectDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProjectDetailLogic {
	return &ProjectDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProjectDetailLogic) ProjectDetail(in *v1.ProjectDetailRequest) (*v1.ProjectDetailResponse, error) {
	// 1. 查项目表
	projectCodeStr, err := l.svcCtx.Encrypter.Decrypt(in.GetProjectCode())
	if err != nil {
		l.Error("[logic ProjectDetail]", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}
	projectCode, err := strconv.ParseInt(projectCodeStr, 10, 64)
	if err != nil {
		l.Error("[logic ProjectDetail]", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}
	memberId := in.GetMemberId()
	projectAndMember, err := l.svcCtx.ProjectRepo.GetProjectAndMemberByPidAndMid(l.ctx, projectCode, memberId)
	if err != nil {
		l.Error("[logic ProjectDetail]", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}

	// 2. 查收藏表 判断收藏状态
	isCollected, err := l.svcCtx.ProjectRepo.IsCollectedByPidAndMid(l.ctx, projectCode, memberId)
	if isCollected {
		projectAndMember.Collected = domain.Collected
	}
	projectDetailRsp := &v1.ProjectDetailResponse{}
	err = copier.Copy(projectDetailRsp, projectAndMember)
	if err != nil {
		l.Error("[logic ProjectDetail]", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}
	projectDetailRsp.CreateTime = time.UnixMilli(projectAndMember.CreateTime).Format(time.DateTime)
	projectDetailRsp.AccessControlType = projectAndMember.GetAccessControlType()
	projectDetailRsp.OrganizationCode, err = l.svcCtx.Encrypter.EncryptInt64(projectAndMember.OrganizationCode)
	if err != nil {
		l.Error("[logic ProjectDetail]", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}
	projectDetailRsp.Code, err = l.svcCtx.Encrypter.EncryptInt64(projectAndMember.Id)
	if err != nil {
		l.Error("[logic ProjectDetail]", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}
	projectDetailRsp.TemplateCode, err = l.svcCtx.Encrypter.Encrypt(strconv.Itoa(projectAndMember.TemplateCode))
	if err != nil {
		l.Error("[logic ProjectDetail]", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}
	return projectDetailRsp, nil
}
