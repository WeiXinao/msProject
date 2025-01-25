package dao

import (
	"context"
	"github.com/WeiXinao/msProject/pkg/ormx"
	"github.com/WeiXinao/xkit/slice"
	"xorm.io/xorm"
)

type projectXormDao struct {
	db *xorm.Engine
}

// GetProjectMembersByPid implements ProjectDao.
func (p *projectXormDao) GetProjectMembersByPid(ctx context.Context, pid int64) ([]*ProjectMember, error) {
	pm := make([]*ProjectMember, 0)	
	err := p.db.Where("project_code = ?", pid).Find(&pm)
	return pm, err
}

func (p *projectXormDao) FindTaskStagesTmplsByProjectTmplId(ctx context.Context, templateCode int) ([]MsTaskStagesTemplate, error) {
	var list []MsTaskStagesTemplate
	err := p.db.Where("project_template_code = ?", templateCode).
		OrderBy("sort DESC, id ASC").
		Find(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (p *projectXormDao) UpdateProject(ctx context.Context, project Project) error {
	_, err := p.db.Context(ctx).ID(project.Id).Update(&project)
	return err
}

func (p *projectXormDao) DeleteProjectCollection(ctx context.Context, memberId int64, projectCode int64) error {
	_, err := p.db.Context(ctx).Delete(&ProjectCollection{
		ProjectCode: projectCode,
		MemberCode:  memberId,
	})
	return err
}

func (p *projectXormDao) SaveProjectCollection(ctx context.Context, projectCollection ProjectCollection) error {
	_, err := p.db.InsertOne(&projectCollection)
	return err
}

func (p *projectXormDao) DeleteProject(ctx context.Context, projectId int64, deleted bool) error {
	isDel := 0
	if deleted {
		isDel = 1
	}
	_, err := p.db.Context(ctx).
		Table(new(Project)).
		Where("id = ?", projectId).
		Update(map[string]any{
			"deleted": isDel,
		})
	return err
}

func (p *projectXormDao) GetProjectAndMemberByPidAndMid(ctx context.Context, pid int64, mid int64) (ProjectAndMember, error) {
	pam := ProjectAndMember{}
	has, err := p.db.Context(ctx).Table("ms_project").
		Join("inner", "ms_project_member", "ms_project.id = ms_project_member.project_code").
		Where("ms_project_member.project_code = ? AND ms_project_member.member_code = ? AND deleted = 0", pid, mid).
		Get(&pam)
	if err != nil {
		return ProjectAndMember{}, err
	}
	if !has {
		return ProjectAndMember{}, ErrRecordNotFound
	}
	return pam, nil
}

func (p *projectXormDao) IsCollectedByPidAndMid(ctx context.Context, pid int64, mid int64) (bool, error) {
	return p.db.Context(ctx).
		Where("project_code = ? AND member_code = ?", pid, mid).
		Get(&ProjectCollection{})
}

func (p *projectXormDao) SaveProjectMember(ctx context.Context, projectMember *ProjectMember) error {
	_, err := p.db.Context(ctx).Insert(projectMember)
	return err
}

func (p *projectXormDao) SaveProject(ctx context.Context, project Project, projectMember ProjectMember) (Project, error) {
	oldDB := p.db
	defer func() {
		p.db = oldDB
	}()
	err := ormx.NewTxSession(p.db.Context(ctx)).Tx(func(session any) error {
		sess, ok := session.(*xorm.Session)
		if !ok {
			return ErrTypeConvert
		}
		p.db = sess.Engine()

		_, err := p.db.Context(ctx).Insert(&project)
		if err != nil {
			return err
		}
		projectMember.ProjectCode = project.Id
		return p.SaveProjectMember(ctx, &projectMember)
	})
	return project, err
}

func (p *projectXormDao) FindInProTemIds(ctx context.Context, ids []int) ([]MsTaskStagesTemplate, error) {
	var taskStageTmpls []MsTaskStagesTemplate
	err := p.db.Context(ctx).
		In("project_template_code", slice.Map(ids, func(idx int, src int) any {
			return src
		})...).
		Find(&taskStageTmpls)
	return taskStageTmpls, err
}

func (p *projectXormDao) FindProjectTemplateSystem(ctx context.Context, page int64,
	size int64) (pts []ProjectTemplate, total int64, err error) {
	offset := (page - 1) * size
	query := p.db.Context(ctx).
		Where("is_system = 1")
	err = query.Limit(int(size), int(offset)).
		Find(&pts)
	if err != nil {
		return
	}
	total, err = query.Count(&ProjectTemplate{})
	return
}

func (p *projectXormDao) FindProjectTemplateCustom(ctx context.Context, memId int64, organizationCode int64,
	page int64, size int64) (pts []ProjectTemplate, total int64, err error) {
	offset := (page - 1) * size
	query := p.db.Context(ctx).
		Where("is_system = 0 AND member_code = ? AND organization_code = ?", memId, organizationCode)
	err = query.Limit(int(size), int(offset)).
		Find(&pts)
	if err != nil {
		return
	}
	total, err = query.Count(&ProjectTemplate{})
	return
}

func (p *projectXormDao) FindProjectTemplateAll(ctx context.Context, organizationCode int64,
	page int64, size int64) (pts []ProjectTemplate, total int64, err error) {
	offset := (page - 1) * size
	query := p.db.Context(ctx).
		Where("organization_code = ?", organizationCode)
	err = query.Limit(int(size), int(offset)).
		Find(&pts)
	if err != nil {
		return
	}
	total, err = query.Count(&ProjectTemplate{})
	return
}

func (p *projectXormDao) FindCollectProjectByMemId(ctx context.Context, memId int64, page int64, size int64) ([]*ProjectAndMember, int64, error) {
	offset := (page - 1) * size
	pam := []*ProjectAndMember{}
	err := p.db.Context(ctx).Table("ms_project").
		Join("inner", "ms_project_member", "ms_project.id = ms_project_member.project_code").
		Where("ms_project.id IN (SELECT project_code FROM ms_project_collection WHERE member_code = ?) AND ms_project.deleted = 0", memId).
		OrderBy("`order`").
		Limit(int(size), int(offset)).
		Find(&pam)
	return pam, int64(len(pam)), err
}

func (p *projectXormDao) FindProjectByMemId(ctx context.Context, selectBy string, memId int64, page int64, size int64) ([]*ProjectAndMember, int64, error) {
	pam := []*ProjectAndMember{}
	//pm := new(ProjectMember)
	//total := int64(0)
	var err error

	offset := (page - 1) * size
	listSess := p.db.Context(ctx).Table("ms_project").
		Join("inner", "ms_project_member", "ms_project.id = ms_project_member.project_code").
		Where("ms_project_member.member_code = ?", memId)
	//cntSess := p.db.Context(ctx).Where("member_code = ?", memId)

	if selectBy != "deleted" {
		listSess.Where("deleted = 0")
	}

	if selectBy == "archive" {
		listSess.Where("archive = 1")
	}
	if selectBy == "deleted" {
		listSess.Where("deleted = 1")
	}
	if selectBy == "" || selectBy == "my" || selectBy == "archive" || selectBy == "deleted" {
		err = listSess.OrderBy("`order`").Limit(int(size), int(offset)).
			Find(&pam)
		//total, err = cntSess.Count(pm)
	}
	if selectBy == "collect" {
		pam, _, err = p.FindCollectProjectByMemId(ctx, memId, page, size)
	}
	if err != nil {
		return nil, 0, err
	}

	return pam, int64(len(pam)), err
}

func (p *projectXormDao) GetMenus(ctx context.Context) ([]*ProjectMenu, error) {
	meuns := []*ProjectMenu{}
	err := p.db.Context(ctx).OrderBy("pid, sort ASC, id ASC").Find(&meuns)
	return meuns, err
}

func NewProjectXormDao(engine *xorm.Engine) (ProjectDao, error) {
	err := engine.Sync(
		new(ProjectMenu),
		new(Project),
		new(ProjectMember),
		new(ProjectCollection),
		new(ProjectTemplate),
		new(MsTaskStagesTemplate),
	)
	if err != nil {
		return nil, err
	}
	return &projectXormDao{
		db: engine,
	}, nil
}
