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
	FindTaskByStageCode(ctx context.Context, stageCode int) ([]*domain.Task, error)
	FindTaskMemberByTaskId(ctx context.Context, taskCode int64, memberId int64) ([]*domain.TaskMember, bool, error)
	FindById(ctx context.Context, id int64) (*domain.TaskStages, bool, error)
	FindTaskMaxIdNum(ctx context.Context, projectCode int64) (int64, error)
	FindTaskSort(ctx context.Context, projectCode int64, stageCode int64) (int64, error)
	CreateTaskAndMember(ctx context.Context, task domain.Task, taskMember domain.TaskMember) (taskId int64, taskMemberId int64, err error)
	FindTaskById(ctx context.Context, id int64) (domain.Task, error)
	UpdateTask(ctx context.Context, ts domain.Task) error
	Move(ctx context.Context, toStageCode int, task domain.Task, nextTask domain.Task) error
	FindTaskByAssignTo(ctx context.Context, memberId int64, done int,
		page int64, pageSize int64) ([]*domain.Task, int64, error)
	FindTaskByMemberCode(ctx context.Context, memberId int64, done int,
		page int64, pageSize int64) ([]*domain.Task, int64, error)
	FindTaskByCreateBy(ctx context.Context, memberId int64, done int,
		page int64, pageSize int64) ([]*domain.Task, int64, error)
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

// FindTaskMemberByTaskIdAndMemberId implements TaskRepo.
func (t *taskRepo) FindTaskMemberByTaskIdAndMemberId(ctx context.Context, taskId int64, taskMemberId int64) (domain.TaskMember, bool, error) {
	taskMember, has, err := t.dao.FindTaskMemberByTaskId(ctx, taskId, taskMemberId)
	if err != nil {
		return domain.TaskMember{}, false, err
	}
	taskDmn := domain.TaskMember{}
	err = copier.Copy(&taskDmn, taskMember)
	if err != nil {
		return domain.TaskMember{}, false, err
	}
	return taskDmn, has, nil
}

// FindTaskByAssignTo implements TaskRepo.
func (t *taskRepo) FindTaskByAssignTo(ctx context.Context, memberId int64, done int,
	page int64, pageSize int64) ([]*domain.Task, int64, error) {
	tasks, total, err := t.dao.FindTaskByAssignTo(ctx, memberId, done, page, pageSize)
	if err != nil {
		return nil, 0, err
	}
	taskDmns := make([]*domain.Task, 0)
	err = copier.CopyWithOption(&taskDmns, tasks, copier.Option{DeepCopy: true})
	return taskDmns, total, err
}

// FindTaskByCreateBy implements TaskRepo.
func (t *taskRepo) FindTaskByCreateBy(ctx context.Context, memberId int64, done int,
	page int64, pageSize int64) ([]*domain.Task, int64, error) {
	tasks, total, err := t.dao.FindTaskByCreateBy(ctx, memberId, done, page, pageSize)
	if err != nil {
		return nil, 0, err
	}
	taskDmns := make([]*domain.Task, 0)
	err = copier.CopyWithOption(&taskDmns, tasks, copier.Option{DeepCopy: true})
	return taskDmns, total, err
}

// FindTaskByMemberCode implements TaskRepo.
func (t *taskRepo) FindTaskByMemberCode(ctx context.Context, memberId int64, done int,
	page int64, pageSize int64) ([]*domain.Task, int64, error) {
	tasks, total, err := t.dao.FindTaskByMemberCode(ctx, memberId, done, page, pageSize)
	if err != nil {
		return nil, 0, err
	}
	taskDmns := make([]*domain.Task, 0)
	err = copier.CopyWithOption(&taskDmns, tasks, copier.Option{DeepCopy: true})
	return taskDmns, total, err
}

// Move implements TaskRepo.
func (t *taskRepo) Move(ctx context.Context, toStageCode int, task domain.Task, nextTask domain.Task) error {
	taskEty, nextTaskEty := dao.Task{}, dao.Task{}
	err := copier.Copy(&taskEty, task)
	if err != nil {
		return err
	}
	err = copier.Copy(&nextTaskEty, nextTask)
	if err != nil {
		return err
	}
	return t.dao.Move(ctx, toStageCode, taskEty, nextTaskEty)
}

// UpdateTaskSort implements TaskRepo.
func (t *taskRepo) UpdateTask(ctx context.Context, ts domain.Task) error {
	task := dao.Task{}
	err := copier.Copy(&task, ts)
	if err != nil {
		return err
	}
	return t.dao.UpdateTask(ctx, task)
}

// FindTaskById implements TaskRepo.
func (t *taskRepo) FindTaskById(ctx context.Context, id int64) (domain.Task, error) {
	ts, err := t.dao.FindTaskById(ctx, id)
	if err != nil {
		return domain.Task{}, err
	}
	taskDmn := domain.Task{}
	err = copier.Copy(&taskDmn, ts)
	if err != nil {
		return domain.Task{}, err
	}
	return taskDmn, nil
}

// CreateTaskAndMember implements TaskRepo.
func (t *taskRepo) CreateTaskAndMember(ctx context.Context, task domain.Task, taskMember domain.TaskMember) (taskId int64, taskMemberId int64, err error) {
	var (
		taskEty       = dao.Task{}
		taskMemberEty = dao.TaskMember{}
	)
	err = copier.Copy(&taskEty, task)
	if err != nil {
		return 0, 0, err
	}
	err = copier.Copy(&taskMemberEty, taskMember)
	if err != nil {
		return 0, 0, err
	}
	err = t.dao.CreateTaskAndMember(ctx, &taskEty, &taskMemberEty)
	if err != nil {
		return 0, 0, err
	}
	return taskEty.Id, taskMemberEty.Id, nil
}

func (t *taskRepo) FindTaskSort(ctx context.Context, projectCode int64, stageCode int64) (int64, error) {
	return t.dao.FindTaskSort(ctx, projectCode, stageCode)
}

// FindTaskMaxIdNum implements TaskRepo.
func (t *taskRepo) FindTaskMaxIdNum(ctx context.Context, projectCode int64) (int64, error) {
	return t.dao.FindTaskMaxIdNum(ctx, projectCode)
}

// FindById implements TaskRepo.
func (t *taskRepo) FindById(ctx context.Context, id int64) (*domain.TaskStages, bool, error) {
	mts, has, err := t.dao.FindById(ctx, id)
	if err != nil {
		return nil, false, err
	}
	ts := &domain.TaskStages{}
	err = copier.CopyWithOption(ts, mts, copier.Option{DeepCopy: true})
	if err != nil {
		return nil, false, err
	}
	return ts, has, nil
}

// FindTaskMemberByTaskId implements TaskRepo.
func (t *taskRepo) FindTaskMemberByTaskId(ctx context.Context, taskCode int64, memberId int64) ([]*domain.TaskMember, bool, error) {
	tm, has, err := t.dao.FindTaskMemberByTaskId(ctx, taskCode, memberId)
	if err != nil {
		return nil, false, err
	}
	tmsDmn := make([]*domain.TaskMember, 0)
	err = copier.CopyWithOption(&tmsDmn, tm, copier.Option{DeepCopy: true})
	if err != nil {
		return nil, false, err
	}
	return tmsDmn, has, nil
}

// FindTaskByStageCode implements TaskRepo.
func (t *taskRepo) FindTaskByStageCode(ctx context.Context, stageCode int) ([]*domain.Task, error) {
	tasksEntity, err := t.dao.FindTaskByStageCode(ctx, stageCode)
	if err != nil {
		return nil, err
	}
	tasks := make([]*domain.Task, 0)
	err = copier.Copy(&tasks, tasksEntity)
	return tasks, err
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
