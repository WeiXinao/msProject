package dao

import (
	"context"

	"github.com/WeiXinao/msProject/pkg/ormx"
	"xorm.io/xorm"
)

type taskXormDao struct {
	db *xorm.Engine
}

// FindTaskByAssignTo implements TaskDao.
func (t *taskXormDao) FindTaskByAssignTo(ctx context.Context, memberId int64, done int,
	page int64, pageSize int64) ([]*Task, int64, error) {
	tasks := make([]*Task, 0)
	condition := "assign_to = ? and done = ? and deleted = 0"
	offset := (page - 1) * pageSize
	err := t.db.Context(ctx).
		Where(condition, memberId, done).
		Limit(int(pageSize), int(offset)).Find(&tasks)
	if err != nil {
		return nil, 0, err
	}
	total, err := t.db.Context(ctx).Where(condition, memberId, done).Count(&Task{})
	return tasks, total, err
}

// FindTaskByCreateBy implements TaskDao.
func (t *taskXormDao) FindTaskByCreateBy(ctx context.Context, memberId int64, done int,
	page int64, pageSize int64) ([]*Task, int64, error) {
	tasks := make([]*Task, 0)
	condition := "create_by = ? and done = ? and deleted = 0"
	offset := (page - 1) * pageSize
	err := t.db.Context(ctx).
		Where(condition, memberId, done).
		Limit(int(pageSize), int(offset)).Find(&tasks)
	if err != nil {
		return nil, 0, err
	}
	total, err := t.db.Context(ctx).Where(condition, memberId, done).Count(&Task{})
	return tasks, total, err
}

// FindTaskByMemberCode implements TaskDao.
func (t *taskXormDao) FindTaskByMemberCode(ctx context.Context, memberId int64, done int,
	page int64, pageSize int64) ([]*Task, int64, error) {
	tasks := make([]*Task, 0)
	offset := (page - 1) * pageSize
	err := t.db.Context(ctx).Table("ms_task").
		Join("inner", "ms_task_member", "ms_task.id = ms_task_member.task_code").
		Where("ms_task_member.member_code = ? AND ms_task.done = ? AND deleted = 0", memberId, done).
		Limit(int(pageSize), int(offset)).
		Find(&tasks)
	if err != nil {
		return nil, 0, err
	}
	total, err := t.db.Context(ctx).Table("ms_task").
		Join("inner", "ms_task_member", "ms_task.id = ms_task_member.task_code").
		Where("ms_task_member.member_code = ? AND ms_task.done = ? AND deleted = 0", memberId, done).
		Count(&Task{})
	return tasks, total, err
}

// Move implements TaskDao.
func (t *taskXormDao) Move(ctx context.Context, toStageCode int, task Task, nextTask Task) error {
	old := t.db
	defer func() {
		t.db = old
	}()
	return ormx.NewTxSession(t.db.Context(ctx)).Tx(func(session any) error {
		sess, ok := session.(*xorm.Session)
		if !ok {
			return ErrTypeConvert
		}
		t.db = sess.Engine()

		// 想将 task 后面的向上移动
		moveUpSql := "UPDATE ms_task SET sort = sort - 1 WHERE stage_code = ? AND sort > ?"
		_, err := t.db.Exec(moveUpSql, task.StageCode, task.Sort)
		if err != nil {
			return err
		}
		// 在将从 nextTask 开始，下面的，往下移动
		task.StageCode = toStageCode
		task.Sort = max(nextTask.Sort, 1)
		moveDownSql := "UPDATE ms_task SET sort = sort + 1 WHERE stage_code = ? AND sort >= ?"
		_, err = t.db.Exec(moveDownSql, task.StageCode, task.Sort)
		if err != nil {
			return err
		}
		return t.UpdateTask(ctx, task)
	})
}

// UpdateTask implements TaskDao.
func (t *taskXormDao) UpdateTask(ctx context.Context, ts Task) error {
	_, err := t.db.Where("id = ?", ts.Id).Update(&ts)
	return err
}

// FindTaskById implements TaskDao.
func (t *taskXormDao) FindTaskById(ctx context.Context, id int64) (Task, error) {
	task := Task{}
	_, err := t.db.Where("id = ?", id).Get(&task)
	return task, err
}

// CreateTaskAndMember implements TaskDao.
func (t *taskXormDao) CreateTaskAndMember(ctx context.Context, task *Task, taskMember *TaskMember) error {
	oldDB := t.db
	defer func() {
		t.db = oldDB
	}()
	return ormx.NewTxSession(t.db.Context(ctx)).Tx(func(session any) error {
		sess, ok := session.(*xorm.Session)
		if !ok {
			return ErrTypeConvert
		}
		t.db = sess.Engine()

		err := t.CreateTask(ctx, task)
		if err != nil {
			return err
		}
		taskMember.TaskCode = task.Id
		err = t.CreateTaskMember(ctx, taskMember)
		return err
	})
}

// CreateTask implements TaskDao.
func (t *taskXormDao) CreateTask(ctx context.Context, task *Task) error {
	_, err := t.db.InsertOne(task)
	return err
}

// CreateTaskMember implements TaskDao.
func (t *taskXormDao) CreateTaskMember(ctx context.Context,
	taskMember *TaskMember) error {
	_, err := t.db.InsertOne(taskMember)
	return err
}

// FindTaskSort implements TaskDao.
func (t *taskXormDao) FindTaskSort(ctx context.Context, projectCode int64, stageCode int64) (int64, error) {
	maxSort := int64(0)
	_, err := t.db.Table(new(Task)).
		Where("project_code = ? AND stage_code = ?", projectCode, stageCode).
		Select("MAX(sort) as maxIdNum").
		Get(&maxSort)
	return maxSort, err
}

// FindTaskMaxIdNum implements TaskDao.
func (t *taskXormDao) FindTaskMaxIdNum(ctx context.Context, projectCode int64) (int64, error) {
	maxIdNum := int64(0)
	_, err := t.db.Table(new(Task)).
		Where("project_code = ?", projectCode).
		Select("MAX(id_num) as maxIdNum").Get(&maxIdNum)
	return maxIdNum, err
}

// FindById implements TaskDao.
func (t *taskXormDao) FindById(ctx context.Context, id int64) (*MsTaskStages, bool, error) {
	ts := &MsTaskStages{}
	has, err := t.db.Where("id = ?", id).Get(ts)
	return ts, has, err
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
	total, err := t.db.Where("project_code = ?", projectCode).Count(new(MsTaskStages))
	if err != nil {
		return nil, 0, err
	}

	err = t.db.Where("project_code = ?", projectCode).
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
		new(MsTaskStages),
	)
	return &taskXormDao{
		db: engine,
	}
}
