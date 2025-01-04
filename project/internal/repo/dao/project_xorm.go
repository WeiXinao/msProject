package dao

import (
	"context"
	"xorm.io/xorm"
)

type projectXormDao struct {
	db *xorm.Engine
}

func (p *projectXormDao) FindProjectByMemId(ctx context.Context, memId int64, page int64, size int64) ([]*ProjectAndMember, int64, error) {
	pam := []*ProjectAndMember{}
	offset := (page - 1) * size
	err := p.db.Context(ctx).Table("ms_project").
		Join("inner", "ms_project_member", "ms_project.id = ms_project_member.project_code").
		Where("ms_project_member.member_code = ?", memId).Limit(int(size), int(offset)).
		Find(&pam)
	if err != nil {
		return nil, 0, err
	}
	pm := new(ProjectMember)
	total, err := p.db.Context(ctx).Where("member_code = ?", memId).Count(pm)
	return pam, total, err
}

func (p *projectXormDao) GetMenus(ctx context.Context) ([]*ProjectMenu, error) {
	meuns := []*ProjectMenu{}
	err := p.db.Context(ctx).Find(&meuns)
	return meuns, err
}

func NewProjectXormDao(engine *xorm.Engine) (ProjectDao, error) {
	err := engine.Sync(
		new(ProjectMenu),
		new(Project),
		new(ProjectMember),
	)
	if err != nil {
		return nil, err
	}
	return &projectXormDao{
		db: engine,
	}, nil
}
