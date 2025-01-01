package repo

import (
	"context"
	"github.com/WeiXinao/msProject/pkg/cachex"
	"github.com/WeiXinao/msProject/project/internal/domain"
	"github.com/WeiXinao/msProject/project/internal/repo/dao"
	"github.com/jinzhu/copier"
)

type projectRepo struct {
	cache cachex.Cache
	dao   dao.ProjectDao
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
