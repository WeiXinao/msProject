package repo

import (
	"context"

	"github.com/WeiXinao/msProject/task/internal/domain"
)

type TaskRepo interface {
	CreateTaskStagesList(ctx context.Context, taskStagesList []*domain.TaskStages) error
	FindStagesByProjectIdPagination(ctx context.Context, projectCode int64,
		page int64, pageSize int64) ([]*domain.TaskStages, int64, error)
	FindTaskByStageCode(ctx context.Context, stageCode int) ([]*domain.Task, error)
	FindTaskMemberByTaskId(ctx context.Context, taskCode int64,
		memberId int64) ([]*domain.TaskMember, bool, error)
	FindById(ctx context.Context, id int64) (*domain.TaskStages, bool, error)
	FindTaskMaxIdNum(ctx context.Context, projectCode int64) (int64, error)
	FindTaskSort(ctx context.Context, projectCode int64, stageCode int64) (int64, error)
	CreateTaskAndMember(ctx context.Context, task domain.Task,
		taskMember domain.TaskMember) (taskId int64, taskMemberId int64, err error)
	FindTaskById(ctx context.Context, id int64) (domain.Task, error)
	UpdateTask(ctx context.Context, ts domain.Task) error
	Move(ctx context.Context, toStageCode int, task domain.Task, nextTask domain.Task) error
	FindTaskByAssignTo(ctx context.Context, memberId int64, done int,
		page int64, pageSize int64) ([]*domain.Task, int64, error)
	FindTaskByMemberCode(ctx context.Context, memberId int64, done int,
		page int64, pageSize int64) ([]*domain.Task, int64, error)
	FindTaskByCreateBy(ctx context.Context, memberId int64, done int,
		page int64, pageSize int64) ([]*domain.Task, int64, error)
	FindTaskMemberByTaskIdPagination(ctx context.Context, taskId int64,
		page int64, pageSize int64) ([]*domain.TaskMember, int64, error)
	SaveTaskWorkTime(ctx context.Context, taskWorkTime domain.TaskWorkTime) error
	FindWorkTimeListByTaskId(ctx context.Context, taskId int64) ([]*domain.TaskWorkTime, int64, error)
}