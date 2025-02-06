package repo

import (
	"context"
	"github.com/WeiXinao/msProject/pkg/cachex"
	"github.com/WeiXinao/msProject/project/internal/domain"
	"github.com/WeiXinao/msProject/project/internal/repo/dao"
	"github.com/WeiXinao/xkit/slice"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type projectRepo struct {
	cache cachex.Cache
	dao   dao.ProjectDao
}

// FindLogByTaskCode implements ProjectRepo.
func (p *projectRepo) FindLogByTaskCode(ctx context.Context, taskCode int64, comment int) ([]*domain.ProjectLog, int64, error) {
	projectLogs, total, err := p.dao.FindLogByTaskCode(ctx, taskCode, comment)
	if err != nil {
		return nil, 0, err
	}
	projectLogDmns := make([]*domain.ProjectLog, 0)
	err = copier.CopyWithOption(&projectLogDmns, projectLogs, copier.Option{DeepCopy: true})
	if err != nil {
		return nil, 0, err
	}
	return projectLogDmns, total, nil
}

// FindLogByTaskCodePagination implements ProjectRepo.
func (p *projectRepo) FindLogByTaskCodePagination(ctx context.Context, taskCode int64, comment int, page int64, pageSize int64) ([]*domain.ProjectLog, int64, error) {
	projectLogs, total, err := p.dao.FindLogByTaskCodePagination(ctx, taskCode, comment, 
		page, pageSize)
	if err != nil {
		return nil, 0, err
	}
	projectLogDmns := make([]*domain.ProjectLog, 0)
	err = copier.CopyWithOption(&projectLogDmns, projectLogs, copier.Option{DeepCopy: true})
	if err != nil {
		return nil, 0, err
	}
	return projectLogDmns, total, nil
}

// SaveProjectLog implements ProjectRepo.
func (p *projectRepo) SaveProjectLog(ctx context.Context, projectLog domain.ProjectLog) error {
	projectLogEty := dao.ProjectLog{}
	err := copier.Copy(&projectLogEty, projectLog)
	if err != nil {
		return err
	}
	return p.dao.SaveProjectLog(ctx, projectLogEty)
}

// FindProjectByIds implements ProjectRepo.
func (p *projectRepo) FindProjectByIds(ctx context.Context, ids []int64) ([]*domain.Project, error) {
	projects, err := p.dao.FindProjectByIds(ctx, ids)
	if err != nil {
		return nil, err
	}
	projectDmns := make([]*domain.Project, 0)
	err = copier.CopyWithOption(&projectDmns, projects, copier.Option{DeepCopy: true})
	return projectDmns, err
}

// FindProjectById implements ProjectRepo.
func (p *projectRepo) FindProjectById(ctx context.Context, id int64) (*domain.Project, error) {
	project, err := p.dao.FindProjectById(ctx, id)
	if err != nil {
		return nil, err
	}
	projectDmn := &domain.Project{}
	err = copier.Copy(projectDmn, project)
	return projectDmn, err
}

// GetProjectMembersByPid implements ProjectRepo.
func (p *projectRepo) GetProjectMembersByPid(ctx context.Context, pid int64) ([]*domain.ProjectMember, error) {
	pmMdls, err := p.dao.GetProjectMembersByPid(ctx, pid)
	if err != nil {
		return nil, err
	}
	pmDmns := make([]*domain.ProjectMember, 0)
	err = copier.Copy(&pmDmns, pmMdls)
	if err != nil {
		return nil, err
	}
	return pmDmns, nil
}

func (p *projectRepo) FindTaskStagesTmplsByProjectTmplId(ctx context.Context, templateCode int) ([]*domain.MsTaskStagesTemplate, error) {
	taskTmpls, err := p.dao.FindTaskStagesTmplsByProjectTmplId(ctx, templateCode)
	var taskTmplDmns []*domain.MsTaskStagesTemplate
	err = copier.Copy(&taskTmplDmns, taskTmpls)
	if err != nil {
		return nil, err
	}
	return taskTmplDmns, nil
}

func (p *projectRepo) UpdateProject(ctx context.Context, project domain.Project) error {
	projectEntity := dao.Project{}
	err := copier.Copy(&projectEntity, project)
	if err != nil {
		return err
	}
	return p.dao.UpdateProject(ctx, projectEntity)
}

func (p *projectRepo) DeleteProjectCollection(ctx context.Context, memberId int64, projectCode int64) error {
	return p.dao.DeleteProjectCollection(ctx, memberId, projectCode)
}

func (p *projectRepo) FindCollectProjectByMemId(ctx context.Context, memId int64, page int64, size int64) (map[int64]*domain.ProjectAndMember, int64, error) {
	pams, total, err := p.dao.FindCollectProjectByMemId(ctx, memId, page, size)
	IdToProjectAndMemberMap := make(map[int64]*domain.ProjectAndMember)
	for _, pam := range pams {
		pa := &domain.ProjectAndMember{}
		err = copier.Copy(pa, &pam)
		if err != nil {
			logx.WithContext(ctx).Error("[repo FindCollectProjectByMemId]", err)
			return nil, 0, err
		}
		IdToProjectAndMemberMap[pam.Project.Id] = pa
	}
	return IdToProjectAndMemberMap, total, nil
}

func (p *projectRepo) SaveProjectCollection(ctx context.Context, projectCollection domain.ProjectCollection) error {
	pc := dao.ProjectCollection{}
	err := copier.Copy(&pc, projectCollection)
	if err != nil {
		return err
	}
	return p.dao.SaveProjectCollection(ctx, pc)
}

func (p *projectRepo) DeleteProject(ctx context.Context, projectId int64, deleted bool) error {
	return p.dao.DeleteProject(ctx, projectId, deleted)
}

func (p *projectRepo) GetProjectAndMemberByPidAndMid(ctx context.Context, pid int64, mid int64) (domain.ProjectAndMember, error) {
	pam, err := p.dao.GetProjectAndMemberByPidAndMid(ctx, pid, mid)
	if err != nil {
		return domain.ProjectAndMember{}, err
	}
	var pamDomain domain.ProjectAndMember
	err = copier.Copy(&pamDomain, pam)
	return pamDomain, err
}

func (p *projectRepo) IsCollectedByPidAndMid(ctx context.Context, pid int64, mid int64) (bool, error) {
	return p.dao.IsCollectedByPidAndMid(ctx, pid, mid)
}

func (p *projectRepo) SaveProject(ctx context.Context, project domain.Project, projectMember domain.ProjectMember) (domain.Project, error) {
	pro := dao.Project{}
	pm := dao.ProjectMember{}
	err := copier.Copy(&pro, project)
	if err != nil {
		return domain.Project{}, err
	}
	err = copier.Copy(&pm, projectMember)
	if err != nil {
		return domain.Project{}, err
	}
	saveProject, err := p.dao.SaveProject(ctx, pro, pm)
	if err != nil {
		return domain.Project{}, err
	}
	sp := domain.Project{}
	err = copier.Copy(&sp, saveProject)
	return sp, err
}

func (p *projectRepo) FindInProTemIds(ctx context.Context, ids []int) ([]domain.MsTaskStagesTemplate, error) {
	taskTmplModels, err := p.dao.FindInProTemIds(ctx, ids)
	if err != nil {
		return nil, err
	}
	var taskTmplDomains []domain.MsTaskStagesTemplate
	err = copier.Copy(&taskTmplDomains, &taskTmplModels)
	return taskTmplDomains, err
}

func (p *projectRepo) FindProjectTemplateSystem(ctx context.Context, page int64, size int64) ([]domain.ProjectTemplate, int64, error) {
	tmpls, total, err := p.dao.FindProjectTemplateSystem(ctx, page, size)
	if err != nil {
		return nil, 0, err
	}
	var tmplDomains []domain.ProjectTemplate
	err = copier.Copy(&tmplDomains, &tmpls)
	return tmplDomains, total, err
}

func (p *projectRepo) FindProjectTemplateCustom(ctx context.Context, memId int64, organizationCode int64, page int64, size int64) ([]domain.ProjectTemplate, int64, error) {
	tmpls, total, err := p.dao.FindProjectTemplateCustom(ctx, memId, organizationCode, page, size)
	if err != nil {
		return nil, 0, err
	}
	var tmplDomains []domain.ProjectTemplate
	err = copier.Copy(&tmplDomains, &tmpls)
	return tmplDomains, total, err
}

func (p *projectRepo) FindProjectTemplateAll(ctx context.Context, organizationCode int64, page int64, size int64) ([]domain.ProjectTemplate, int64, error) {
	tmpls, total, err := p.dao.FindProjectTemplateAll(ctx, organizationCode, page, size)
	if err != nil {
		return nil, 0, err
	}
	var tmplDomains []domain.ProjectTemplate
	err = copier.Copy(&tmplDomains, &tmpls)
	return tmplDomains, total, err
}

func (p *projectRepo) FindProjectByMemId(ctx context.Context, selectBy string, memId int64, page int64, size int64) ([]*domain.ProjectAndMember, int64, error) {
	pms, total, err := p.dao.FindProjectByMemId(ctx, selectBy, memId, page, size)
	pams := slice.Map[*dao.ProjectAndMember, *domain.ProjectAndMember](pms,
		func(idx int, src *dao.ProjectAndMember) *domain.ProjectAndMember {
			pam := &domain.ProjectAndMember{}
			er := copier.Copy(pam, &src.ProjectMember)
			if er != nil {
				logx.WithContext(ctx).Error("[repo FindProjectByMemId]", er)
			}
			er = copier.Copy(pam, &src.Project)
			if er != nil {
				logx.WithContext(ctx).Error("[repo FindProjectByMemId]", er)
			}
			return pam
		})
	return pams, total, err
}

func (p *projectRepo) GetMenus(ctx context.Context) ([]*domain.ProjectMenu, error) {
	menus, err := p.dao.GetMenus(ctx)
	if err != nil {
		return nil, err
	}
	menuDos := []*domain.ProjectMenu{}
	err = copier.Copy(&menuDos, &menus)
	return menuDos, err
}

func NewProjectRepo(cache cachex.Cache, dao dao.ProjectDao) ProjectRepo {
	return &projectRepo{
		cache: cache,
		dao:   dao,
	}
}
