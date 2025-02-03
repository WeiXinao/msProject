package logic

import (
	"context"
	"strconv"
	"time"

	"github.com/WeiXinao/msProject/pkg/respx"
	"github.com/WeiXinao/msProject/project/internal/domain"
	"github.com/WeiXinao/msProject/project/internal/events/task"

	"github.com/WeiXinao/msProject/api/proto/gen/project/v1"
	// taskv1 "github.com/WeiXinao/msProject/api/proto/gen/task/v1"
	"github.com/WeiXinao/msProject/project/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveProjectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveProjectLogic {
	return &SaveProjectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SaveProjectLogic) SaveProject(in *v1.SaveProjectReq) (*v1.SaveProjectRsp, error) {
	organizationCodeStr, err := l.svcCtx.Encrypter.Decrypt(in.OrganizationCode)
	if err != nil {
		l.Error("[logic SaveProject]", err)
		return nil, respx.ErrInternalServer
	}
	organizationCode, err := strconv.ParseInt(organizationCodeStr, 10, 64)
	if err != nil {
		l.Error("[logic SaveProject]", err)
		return nil, respx.ErrInternalServer
	}
	templateCodeStr, err := l.svcCtx.Encrypter.Decrypt(in.TemplateCode)
	if err != nil {
		l.Error("[logic SaveProject]", err)
		return nil, respx.ErrInternalServer
	}
	templateCode, err := strconv.Atoi(templateCodeStr)
	if err != nil {
		l.Error("[logic SaveProject]", err)
		return nil, respx.ErrInternalServer
	}
	// 1. 保存项目表
	project := domain.Project{
		Name:              in.Name,
		Cover:             "https://img2.baidu.com/it/u=792555388,2449797505&fm=253&fmt=auto&app=138&f=JPEG?w=667&h=500",
		Description:       in.Description,
		AccessControlType: domain.AccessControlTypeOpen,
		Deleted:           domain.Undeleted,
		TemplateCode:      templateCode,
		CreateTime:        time.Now().UnixMilli(),
		OrganizationCode:  organizationCode,
		Archive:           domain.Unarchived,
		TaskBoardTheme:    domain.TaskBoardThemeSimple,
	}
	// 2. 保存项目和成员的关联表
	projectMember := domain.ProjectMember{
		MemberCode: in.MemberId,
		JoinTime:   time.Now().UnixMilli(),
		IsOwner:    in.MemberId,
	}
	savedProject, err := l.svcCtx.ProjectRepo.SaveProject(l.ctx, project, projectMember)
	if err != nil {
		l.Error("[logic SaveProject]", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer) 
	}

	code, err := l.svcCtx.Encrypter.EncryptInt64(savedProject.Id)
	if err != nil {
		l.Error("[logic SaveProject]", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer) 
	}
	// 3. 生成任务的步骤
	taskStageTmplateList, err := l.svcCtx.
	ProjectRepo.
	FindTaskStagesTmplsByProjectTmplId(l.ctx, templateCode)
	if err != nil {
		l.Error("[logic SaveProject]", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer) 
	}
	// taskStageList := make([]*taskv1.SaveTaskStagesMessage, 0)
	taskStageList := make([]*domain.TaskStages, 0)
	for idx, mtst := range taskStageTmplateList {
		taskStage := &domain.TaskStages{
			ProjectCode: savedProject.Id,
			Name: mtst.Name,	
			Sort: idx + 1,
			CreateTime: time.Now().UnixMilli(),	
			Deleted: domain.Undeleted,
		}
		taskStageList = append(taskStageList, taskStage)
	}
	// _, err = l.svcCtx.TaskClient.SaveTaskStages(l.ctx, &taskv1.SaveTaskStagesRequest{
	// 	List: taskStageList,
	// })

	err = l.svcCtx.TaskProducer.ProduceSaveTaskStagesEvent(l.ctx, task.SaveTaskStagesEvent{
		TaskStagesList: taskStageList,
	})
	if err != nil {
		l.Error("[logic SaveProject]", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer) 
	}

	return &v1.SaveProjectRsp{
		Id:               savedProject.Id,
		Cover:            savedProject.Cover,
		Name:             savedProject.Name,
		Description:      savedProject.Description,
		Code:             code,
		CreateTime:       time.UnixMilli(savedProject.CreateTime).Format(time.DateTime),
		TaskBoardTheme:   savedProject.TaskBoardTheme,
		OrganizationCode: organizationCodeStr,
	}, nil
}
