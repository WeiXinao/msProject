package repo

import (
	"context"

	"github.com/WeiXinao/msProject/task/internal/domain"
	"github.com/WeiXinao/msProject/task/internal/repo/dao"
	"github.com/jinzhu/copier"
)

type TaskRepo interface {
	CreateTaskStagesList(ctx context.Context, taskStagesList []*domain.TaskStages) error
	FindStagesByProjectIdPagination(ctx context.Context, projectCode int64,
		page int64, pageSize int64) ([]*domain.TaskStages, int64, error)
}

func (t *taskRepo) CreateTaskStagesList(ctx context.Context,
	taskStagesList []*domain.TaskStages) error {
	taskStageMdls := make([]*dao.MsTaskStages, 0)
	err := copier.CopyWithOption(&taskStageMdls, taskStagesList,
		copier.Option{DeepCopy: true})
	if err != nil {
		return err
	}
	return t.dao.CreateTaskStagesList(ctx,
		taskStageMdls)
}

type taskRepo struct {
	dao dao.TaskDao
}

func (t *taskRepo) FindStagesByProjectIdPagination(ctx context.Context, projectCode int64, 
	page int64, pageSize int64) ([]*domain.TaskStages, int64, error) {
	taskStages, total, err := t.dao.FindStagesByProjectIdPagination(ctx, projectCode, 
	page, pageSize)
	if err != nil {
		return nil, 0, err
	}
	taskStagesDmns := make([]*domain.TaskStages, 0)	
	err = copier.CopyWithOption(&taskStagesDmns, taskStages, copier.Option{DeepCopy: true})
	if err != nil {
		return nil, 0, err
	}
	return taskStagesDmns, total, nil
}

func NewTaskRepo(dao dao.TaskDao) TaskRepo {
	return &taskRepo{
		dao: dao,
	}
}
