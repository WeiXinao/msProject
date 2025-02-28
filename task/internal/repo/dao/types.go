package dao

import (
	"context"
	"errors"
)

var (
	ErrTypeConvert = errors.New("类型转换错误")
)

type TaskDao interface {
	CreateTaskStagesList(ctx context.Context, taskStagesList []*MsTaskStages) error
	FindStagesByProjectIdPagination(ctx context.Context, projectCode int64,
		page int64, pageSize int64) ([]*MsTaskStages, int64, error)
	FindTaskByStageCode(ctx context.Context, stageCode int) ([]*Task, error)
    FindTaskMemberByTaskId(ctx context.Context, taskCode int64, memberId int64) ([]*TaskMember, bool, error)
	FindById(ctx context.Context, id int64) (*MsTaskStages, bool, error)
	FindTaskMaxIdNum(ctx context.Context, projectCode int64) (int64, error)
	FindTaskSort(ctx context.Context, projectCode int64, stageCode int64) (int64, error)
    CreateTaskAndMember(ctx context.Context, task *Task, taskMember *TaskMember) error
    CreateTask(ctx context.Context, task *Task) error
    CreateTaskMember(ctx context.Context, task *TaskMember) error
	FindTaskById(ctx context.Context, id int64) (Task, error)
	UpdateTask(ctx context.Context, ts Task) error
	Move(ctx context.Context, toStageCode int, task Task, nextTask Task) error
	FindTaskByAssignTo(ctx context.Context, memberId int64, done int,
        page int64, pageSize int64) ([]*Task, int64, error)
	FindTaskByMemberCode(ctx context.Context, memberId int64, done int, 
        page int64, pageSize int64) ([]*Task, int64, error)
	FindTaskByCreateBy(ctx context.Context, memberId int64, done int,
        page int64, pageSize int64) ([]*Task, int64, error)
	FindTaskMemberByTaskIdPagination(ctx context.Context, taskId int64, 
		page int64, pageSize int64) ([]*TaskMember, int64, error)
	FindTaskByIds(ctx context.Context, taskIds []int64) ([]*Task, error)

	SaveTaskWorkTime(ctx context.Context, taskWorkTime TaskWorkTime) error
	FindWorkTimeListByTaskId(ctx context.Context, taskId int64) ([]*TaskWorkTime, int64, error)

	SaveProjectLog(ctx context.Context, projectLog ProjectLog) error
	FindLogByTaskCode(ctx context.Context, taskCode int64, comment int) ([]*ProjectLog,
	int64, error)
	FindLogByTaskCodePagination(ctx context.Context, taskCode int64, comment int, 
		page int64, pageSize int64) ([]*ProjectLog, int64, error)

	SaveComment(ctx context.Context, comment ProjectLog) error
	FindLogByMemberCode(ctx context.Context, memberId int64, page int64, 
		pageSize int64) ([]*ProjectLog, int64, error)
}

type MsTaskStages struct {
	Id          int    `xorm:"'id' int autoincr pk notnull"`
	Name        string `xorm:"'name' varchar(255) null default(null) comment('类型名称')"`
	ProjectCode int64  `xorm:"'project_code' bigint null comment('项目id')"`
	Sort        int    `xorm:"'sort' int null default(0) comment('排序')"`
	Description string `xorm:"'description' text null default('') comment('备注')"`
	CreateTime  int64  `xorm:"'create_time' bigint null comment('创建时间')"`
	Deleted     int    `xorm:"'deleted' tinyint(1) null default(0) comment('删除标记')"`
}

func (*MsTaskStages) TableName() string {
	return "ms_task_stages"
}

type Task struct {
    Id int64 `xorm:"'id' bigint unsigned autoincr pk notnull"`
    ProjectCode int64 `xorm:"'project_code' bigint pk notnull comment('项目编号')"`
    Name string `xorm:"'name' varchar(255) null default(null)"`
    Pri int `xorm:"'pri' tinyint unsigned null default(0) comment('紧急程度')"`
    ExecuteStatus int `xorm:"'execute_status' tinyint null default(null) comment('执行状态')"`
    Description string `xorm:"'description' text null comment('详情')"`
    CreateBy int64 `xorm:"'create_by' bigint null default(null) comment('创建人')"`
    DoneBy int64 `xorm:"'done_by' bigint null default(null) comment('完成人')"`
    DoneTime int64 `xorm:"'done_time' bigint null default(null) comment('完成时间')"`
    CreateTime int64 `xorm:"'create_time' bigint null default(null) comment('创建日期')"`
    AssignTo int64 `xorm:"'assign_to' bigint null default(null) comment('指派给谁')"`
    Deleted int `xorm:"'deleted' tinyint(1) null default(0) comment('回收站')"`
    StageCode int `xorm:"'stage_code' int null default(null) comment('任务列表')"`
    TaskTag string `xorm:"'task_tag' varchar(255) null default(null) comment('任务标签')"`
    Done int `xorm:"'done' tinyint null default(0) comment('是否完成')"`
    BeginTime int64 `xorm:"'begin_time' bigint null default(null) comment('开始时间')"`
    EndTime int64 `xorm:"'end_time' bigint null default(null) comment('截止时间')"`
    RemindTime int64 `xorm:"'remind_time' bigint null default(null) comment('提醒时间')"`
    Pcode int64 `xorm:"'pcode' bigint null default(null) comment('父任务id')"`
    Sort int `xorm:"'sort' int null default(0) comment('排序')"`
    Like int `xorm:"'like' int null default(0) comment('点赞数')"`
    Star int `xorm:"'star' int null default(0) comment('收藏数')"`
    DeletedTime int64 `xorm:"'deleted_time' bigint null default(null) comment('删除时间')"`
    Private int `xorm:"'private' tinyint(1) null default(0) comment('是否隐私模式')"`
    IdNum int `xorm:"'id_num' int null default(1) comment('任务id编号')"`
    Path string `xorm:"'path' text null comment('上级任务路径')"`
    Schedule int `xorm:"'schedule' int null default(0) comment('进度百分比')"`
    VersionCode int64 `xorm:"'version_code' bigint null default(0) comment('版本id')"`
    FeaturesCode int64 `xorm:"'features_code' bigint null default(0) comment('版本库id')"`
    WorkTime int `xorm:"'work_time' int null default(0) comment('预估工时')"`
    Status int `xorm:"'status' tinyint null default(0) comment('执行状态。0：未开始，1：已完成，2：进行中，3：挂起，4：测试中')"`
}

func (*Task) TableName() string {
    return "ms_task"
}

type TaskMember struct {
    Id int64 `xorm:"'id' bigint autoincr pk notnull"`
    TaskCode int64 `xorm:"'task_code' bigint null default(0) comment('任务ID')"`
    IsExecutor int `xorm:"'is_executor' tinyint(1) null default(0) comment('执行者')"`
    MemberCode int64 `xorm:"'member_code' bigint null default(null) comment('成员id')"`
    JoinTime int64 `xorm:"'join_time' bigint null default(null)"`
    IsOwner int `xorm:"'is_owner' tinyint(1) null default(0) comment('是否创建人')"`
}

func (*TaskMember) TableName() string {
    return "ms_task_member"
}
type TaskWorkTime struct {
    Id int64 `xorm:"'id' bigint autoincr pk notnull"`
    TaskCode int64 `xorm:"'task_code' bigint null default(0) comment('任务ID')"`
    MemberCode int64 `xorm:"'member_code' bigint null default(null) comment('成员id')"`
    CreateTime int64 `xorm:"'create_time' bigint null default(null)"`
    Content string `xorm:"'content' varchar(500) null default(null) comment('描述')"`
    BeginTime int64 `xorm:"'begin_time' bigint null default(null) comment('开始时间')"`
    Num int `xorm:"'num' int null default(0) comment('工时')"`
}

func (*TaskWorkTime) TableName() string {
    return "ms_task_work_time"
}

type ProjectLog struct {
    Id int64 `xorm:"'id' bigint autoincr pk notnull"`
    MemberCode int64 `xorm:"'member_code' bigint null default(0) comment('操作人id')"`
    Content string `xorm:"'content' text null comment('操作内容')"`
    Remark string `xorm:"'remark' text null`
    Type string `xorm:"'type' varchar(255) null default('create') comment('操作类型')"`
    CreateTime int64 `xorm:"'create_time' bigint null default(null) comment('添加时间')"`
    SourceCode int64 `xorm:"'source_code' bigint null default(0) comment('任务id')"`
    ActionType string `xorm:"'action_type' varchar(30) null default(null) comment('场景类型')"`
    ToMemberCode int64 `xorm:"'to_member_code' bigint null default(0)"`
    IsComment int `xorm:"'is_comment' tinyint(1) null default(0) comment('是否评论，0：否')"`
    ProjectCode int64 `xorm:"'project_code' bigint null default(null)"`
    Icon string `xorm:"'icon' varchar(20) null default(null)"`
    IsRobot int `xorm:"'is_robot' tinyint(1) null default(0) comment('是否机器人')"`
}

func (*ProjectLog) TableName() string {
    return "ms_project_log"
}