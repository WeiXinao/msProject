package repo

import (
	"context"
	"github.com/WeiXinao/msProject/project/internal/domain"
)

type ProjectRepo interface {
	GetMenus(ctx context.Context) ([]*domain.ProjectMenu, error)
}
