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

func (p *projectRepo) FindProjectByMemId(ctx context.Context, memId int64, page int64, size int64) ([]*domain.ProjectAndMember, int64, error) {
	pms, total, err := p.dao.FindProjectByMemId(ctx, memId, page, size)
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
