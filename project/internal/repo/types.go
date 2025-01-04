package repo

import (
	"context"
	"github.com/WeiXinao/msProject/project/internal/domain"
)

type ProjectRepo interface {
	GetMenus(ctx context.Context) ([]*domain.ProjectMenu, error)
	FindProjectByMemId(ctx context.Context, memId int64, page int64,
		size int64) ([]*domain.ProjectAndMember, int64, error)
}
