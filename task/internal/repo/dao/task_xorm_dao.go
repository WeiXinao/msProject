package dao

import (
	"context"

	"github.com/WeiXinao/msProject/pkg/ormx"
	"xorm.io/xorm"
)

type taskXormDao struct {
	db *xorm.Engine
}


// CreateTaskAndMember implements TaskDao.
func (t *taskXormDao) CreateTaskAndMember(ctx context.Context, task *Task, taskMember *TaskMember) error {
	oldDB := t.db
	defer func ()  {
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
