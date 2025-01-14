package repo

import (
	"context"
	"github.com/WeiXinao/msProject/project/internal/domain"
)

type ProjectRepo interface {
	GetMenus(ctx context.Context) ([]*domain.ProjectMenu, error)

	FindProjectByMemId(ctx context.Context, selectBy string, memId int64, page int64,
		size int64) ([]*domain.ProjectAndMember, int64, error)
	SaveProject(ctx context.Context, project domain.Project, projectMember domain.ProjectMember) (domain.Project, error)

	FindProjectTemplateSystem(ctx context.Context, page int64, size int64) ([]domain.ProjectTemplate, int64, error)
	FindProjectTemplateCustom(ctx context.Context, memId int64, organizationCode int64, page int64, size int64) ([]domain.ProjectTemplate, int64, error)
	FindProjectTemplateAll(ctx context.Context, organizationCode int64, page int64, size int64) ([]domain.ProjectTemplate, int64, error)
	FindInProTemIds(ctx context.Context, ids []int) ([]domain.MsTaskStagesTemplate, error)
	GetProjectAndMemberByPidAndMid(ctx context.Context, pid int64, mid int64) (domain.ProjectAndMember, error)
	IsCollectedByPidAndMid(ctx context.Context, pid int64, mid int64) (bool, error)
}
