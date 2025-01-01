package dao

import (
	"context"
	"xorm.io/xorm"
)

type projectXormDao struct {
	db *xorm.Engine
}

func (p *projectXormDao) GetMenus(ctx context.Context) ([]*ProjectMenu, error) {
	meuns := []*ProjectMenu{}
	err := p.db.Context(ctx).Find(&meuns)
	return meuns, err
}

func NewProjectXormDao(engine *xorm.Engine) (ProjectDao, error) {
	err := engine.Sync(
		new(ProjectMenu),
	)
	if err != nil {
		return nil, err
	}
	return &projectXormDao{
		db: engine,
	}, nil
}
