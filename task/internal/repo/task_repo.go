package repo

import (
	"context"

	"github.com/WeiXinao/msProject/task/internal/domain"
	"github.com/WeiXinao/msProject/task/internal/repo/dao"
	"github.com/jinzhu/copier"
)

// SaveComment implements TaskRepo.
func (p *taskRepo) SaveComment(ctx context.Context, comment domain.ProjectLog) error {
	commentEty := dao.ProjectLog{}
	err := copier.Copy(&commentEty, comment)
	if err != nil {
		return err
	}
	return p.dao.SaveComment(ctx, commentEty)
}

// FindLogByTaskCode implements ProjectRepo.
func (p *taskRepo) FindLogByTaskCode(ctx context.Context, taskCode int64, comment int) ([]*domain.ProjectLog, int64, error) {
	projectLogs, total, err := p.dao.FindLogByTaskCode(ctx, taskCode, comment)
	if err != nil {
		return nil, 0, err
	}
	projectLogDmns := make([]*domain.ProjectLog, 0)
	err = copier.CopyWithOption(&projectLogDmns, projectLogs, copier.Option{DeepCopy: true})
	if err != nil {
		return nil, 0, err
	}
	return projectLogDmns, total, nil
}

// FindLogByTaskCodePagination implements ProjectRepo.
func (p *taskRepo) FindLogByTaskCodePagination(ctx context.Context, taskCode int64, comment int, page int64, pageSize int64) ([]*domain.ProjectLog, int64, error) {
	projectLogs, total, err := p.dao.FindLogByTaskCodePagination(ctx, taskCode, comment,
		page, pageSize)
	if err != nil {
		return nil, 0, err
	}
	projectLogDmns := make([]*domain.ProjectLog, 0)
	err = copier.CopyWithOption(&projectLogDmns, projectLogs, copier.Option{DeepCopy: true})
	if err != nil {
		return nil, 0, err
	}
	return projectLogDmns, total, nil
}

// SaveProjectLog implements ProjectRepo.
func (p *taskRepo) SaveProjectLog(ctx context.Context, projectLog domain.ProjectLog) error {
	projectLogEty := dao.ProjectLog{}
	err := copier.Copy(&projectLogEty, projectLog)
	if err != nil {
		return err
	}
	return p.dao.SaveProjectLog(ctx, projectLogEty)
}

// FindWorkTimeListByTaskId implements TaskRepo.
func (t *taskRepo) FindWorkTimeListByTaskId(ctx context.Context, taskId int64) ([]*domain.TaskWorkTime, int64, error) {
	taskWorkTimeList, total, err := t.dao.FindWorkTimeListByTaskId(ctx, taskId)
	if err != nil {
		return nil, 0, err
	}
	taskWorkTimeDmns := make([]*domain.TaskWorkTime, 0)
	err = copier.Copy(&taskWorkTimeDmns, taskWorkTimeList)
	if err != nil {
		return nil, 0, err
	}
	return taskWorkTimeDmns, total, nil
}

// SaveTaskWorkTime implements TaskRepo.
func (t *taskRepo) SaveTaskWorkTime(ctx context.Context, taskWorkTime domain.TaskWorkTime) error {
	taskWorkTimeEty := dao.TaskWorkTime{}
	err := copier.Copy(&taskWorkTimeEty, taskWorkTime)
	if err != nil {
		return err
	}
	return t.dao.SaveTaskWorkTime(ctx, taskWorkTimeEty)
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

// FindTaskMemberByTaskIdPagination implements TaskRepo.
func (t *taskRepo) FindTaskMemberByTaskIdPagination(ctx context.Context, taskId int64, page int64, pageSize int64) ([]*domain.TaskMember, int64, error) {
	taskMembers, total, err := t.dao.FindTaskMemberByTaskIdPagination(ctx, taskId,
		page, pageSize)
	if err != nil {
		return nil, 0, err
	}
	taskDmns := make([]*domain.TaskMember, 0)
	err = copier.Copy(&taskDmns, taskMembers)
	if err != nil {
		return nil, 0, err
	}
	return taskDmns, total, nil
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

type taskRepo struct {
	dao dao.TaskDao
}


func NewTaskRepo(dao dao.TaskDao) TaskRepo {
	return &taskRepo{
		dao: dao,
	}
}
