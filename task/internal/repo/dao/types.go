package dao

import "context"

type TaskDao interface {
	CreateTaskStagesList(ctx context.Context, taskStagesList []*MsTaskStages) error
	FindStagesByProjectIdPagination(ctx context.Context, projectCode int64,
		page int64, pageSize int64) ([]*MsTaskStages, int64, error)
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