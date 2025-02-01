package dao

import (
	"context"
	"xorm.io/xorm"
)

type taskXormDao struct {
	db *xorm.Engine
}

// FindTaskMemberByTaskId implements TaskDao.
func (t *taskXormDao) FindTaskMemberByTaskId(ctx context.Context, taskCode int64, memberId int64) ([]*TaskMember, bool, error) {
	taskMembers := make([]*TaskMember, 0)
	has, err := t.db.Where("task_code = ? AND member_code = ?", taskCode, memberId).Get(&taskMembers)
	return taskMembers, has, err
}

// FindTaskByStageCode implements TaskDao.
func (t *taskXormDao) FindTaskByStageCode(ctx context.Context, stageCode int) ([]*Task, error) {
	tasks := make([]*Task, 0)
	err := t.db.Where("stage_code = ? AND deleted = 0", stageCode).
		OrderBy("sort ASC").
		Find(&tasks)
	return tasks, err
}

// FindStagesByProjectIdPagination implements TaskDao.
func (t *taskXormDao) FindStagesByProjectIdPagination(ctx context.Context, projectCode int64, page int64, pageSize int64) ([]*MsTaskStages, int64, error) {
	taskStagesList := make([]*MsTaskStages, 0)
	offset := (page - 1) * pageSize
	sqlClause := t.db.Where("project_code = ?", projectCode)
	total, err := sqlClause.Count(new(MsTaskStages))
	if err != nil {
		return nil, 0, err
	}

	err = sqlClause.
		OrderBy("sort ASC").
		Limit(int(pageSize), int(offset)).
		Find(&taskStagesList)
	if err != nil {
		total = 0
	}

	return taskStagesList, total, err
}

// CreateTaskStagesList implements TaskDao.
func (t *taskXormDao) CreateTaskStagesList(ctx context.Context,
	taskStagesList []*MsTaskStages) error {
	_, err := t.db.Insert(&taskStagesList)
	return err
}

func NewTaskXormDao(engine *xorm.Engine) TaskDao {
	engine.Sync(
		new(Task),
		new(TaskMember),
	)
	return &taskXormDao{
		db: engine,
	}
}
